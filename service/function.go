package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func (_ *Utility) HashSHA256(input string) string { // 创建hash256
	hash := hmac.New(sha256.New, []byte(GetEvn("HASH_KEY")))
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func (_ *Utility) VerifySHA256(input string, Hash string) (bool, error) { // 验证token
	if Hash == utilityFunction.HashSHA256(input) {
		return true, nil
	}
	return false, nil
}

// 中间件

func (_ *Utility) UserPasswdVerify(userName string, passwd string) (gin.H, bool) { // 用户密码认证函数
	DBData, err := DataBase.ReadOneDB("user", bson.D{bson.E{Key: "name", Value: userName}}, gin.H{})
	if err != nil {
		return nil, false
	}
	err = bcrypt.CompareHashAndPassword([]byte(DBData["passwd"].(string)), []byte(passwd)) // 密码验证
	// fmt.Println(DBData)
	if err != nil {
		return nil, false
	} else {
		token, err := utilityFunction.JWTCreate(DBData["_id"].(primitive.ObjectID).Hex())
		if err != nil {
			return nil, false
		}
		refreshToken, err := utilityFunction.JWTRefreshCreate(DBData["_id"].(primitive.ObjectID).Hex())
		if err != nil {
			return nil, false
		}
		return gin.H{"token": token, "refresh_token": refreshToken}, true
	}
}

func (_ *Utility) ReturnHeader() gin.HandlerFunc { // 通过cookie认证
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Allow", "GET,POST,PUT,DELETE")
	}
}

func (_ *Utility) CheckLoginMiddleware() gin.HandlerFunc { // 通过cookie认证
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if ok, data, err := utilityFunction.JWTOptVerify(token); ok {
			if err != nil {
				ctx.JSON(http.StatusOK, GeneralJSONHeader{
					Code: ServerError,
					Msg:  "server error",
					Path: ctx.Request.URL.Path,
					Data: nil,
				})
			}
			if ok, _ := utilityFunction.VerifySHA256(data["user_id"].(string), data["verify"].(string)); ok {
				ctx.Set("user_id", data["user_id"]) // 将操作用户id添加到传入变量
				ctx.Next()
			}
			ctx.Abort()
		} else {
			ctx.JSON(http.StatusOK, GeneralJSONHeader{
				Code: ServerError,
				Msg:  "JWT verify error",
				Path: ctx.Request.URL.Path,
				Data: nil,
			})
			ctx.Abort()
		}
	}
}

func (_ *Utility) NoFoundPage(ctx *gin.Context) { // 404
	ctx.JSON(http.StatusNotFound, GeneralJSONHeader{
		Code: http.StatusNotFound,
		Msg:  "页面走丢了~",
		Path: ctx.Request.URL.Path,
		Data: nil,
	})
}

func (_ *Utility) Log(data ...any) { // Debug输出
	if gin.Mode() == gin.DebugMode {
		fmt.Println(data...)
	}
}

func (_ *Utility) JWTCreate(uid string) (string, error) { // 创建JWT认证码
	claims := jwt.MapClaims{
		"user_id": uid,                                  // 用户id
		"verify":  utilityFunction.HashSHA256(uid),      // 校验用户id
		"exp":     time.Now().Add(time.Hour * 2).Unix(), // 生效时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(GetEvn("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (_ *Utility) JWTRefreshCreate(uid string) (string, error) { // 创建JWT刷新认证码
	claims := jwt.MapClaims{
		"token": uid,
		"exp":   time.Now().Add(time.Hour * 168).Unix(), // 一周生效时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(GetEvn("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (_ *Utility) JWTOptVerify(tokenString string) (bool, gin.H, error) { // 认证JWT操作认证码
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // 验证JWT有效性
			return nil, errors.New("signature is invalid")
		} else {
			return []byte(GetEvn("JWT_KEY")), nil
		}
	})
	if err != nil {
		// 验证签名有效性
		return false, nil, errors.New("JWT Verify Error")
	}
	if token.Valid { // 验证JWT有效性
		// 有效
		// 校验是否为用户token
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// 获取需要的字段值
			if userId, exists := claims["user_id"].(string); exists {
				if verify, ok := claims["verify"].(string); ok {
					return true, gin.H{"user_id": userId, "verify": verify}, nil
				} else {
					return false, nil, nil
				}
			} else {
				return false, nil, nil
			}
			// 可以继续获取其他字段...
		} else {
			return false, nil, nil
		}
	} else {
		// 无效
		return false, nil, errors.New("JWT Verify Error")
	}
}

func (_ *Utility) JWTRefreshVerify(tokenString string) (bool, gin.H, error) { //认证JWT刷新码
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // 验证JWT有效性
			return nil, errors.New("signature is invalid")
		} else {
			return []byte(GetEvn("JWT_KEY")), nil
		}
	})
	if err != nil {
		// 验证签名有效性
		if validationErr, ok := err.(*jwt.ValidationError); ok { // 判断Token是否过期
			if validationErr.Errors == jwt.ValidationErrorExpired {
				// JWT 过期,生成新的用户 JWT
				return true, gin.H{"user_id": token.Claims.(jwt.MapClaims)["user_id"]}, nil
			}
			return false, nil, errors.New("未验证")
		}
	}
	return false, nil, errors.New("未过期")
}

func GetEvn(key string) string {
	return os.Getenv(key)
}

func ServerBegin() { // 服务器启动前配置函数
	// 初始化环境变量
	err := godotenv.Load("Service.env")
	if err != nil {
		utilityFunction.Log(err)
	}

	DataBase = DBService{ // 初始化数据库函数
		Client: InitDB(),
	}
}

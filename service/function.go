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
	"net/http"
	"os"
	"time"
)

func (_ *Utility) HashSHA256(input string) string { // 创建hash256
	hash := hmac.New(sha256.New, []byte(GetEvn("HASH_KEY")))
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

// 中间件

func (_ *Utility) UserPasswdVerify(userName string, passwd string) bool { // 用户密码认证函数
	DBData, err := DataBase.ReadOneDB("user", bson.D{bson.E{Key: "name", Value: userName}}, gin.H{})
	if err != nil {
		return false
	}
	if DBData["name"] != userName {
		return false
	}

	if DBData["passwd"] != passwd {
		return false
	}
	return true
}

func (_ *Utility) ReturnHeader() gin.HandlerFunc { // 通过cookie认证
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Allow", "GET,POST,PUT,DELETE")
	}
}

func (_ *Utility) CheckLoginMiddleware() gin.HandlerFunc { // 通过cookie认证
	return func(ctx *gin.Context) {
		// cookie, err := ctx.Cookie("")
	}
}

func (_ *Utility) verifyHeaderLoginCode() gin.HandlerFunc { // 通过请求头认证
	return func(ctx *gin.Context) {
		// fmt.Println(ctx.Request.Header.Get("token")) // 获取请求头认证信息
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
		"user_id": uid,                                   // 用户id
		"verify":  utilityFunction.HashSHA256(uid),       // 校验用户id
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 生效时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString([]byte(GetEvn("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (_ *Utility) JWTVerify(tokenString string) (bool, error) { // 认证JWT认证码
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetEvn("JWT_KEY")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	if claims["verify"] != utilityFunction.HashSHA256(claims["user_id"].(string)) {
		return false, errors.New("JWT Verify Error")
	}
	return true, nil
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

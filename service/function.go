package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"net/http"
)

// 中间件

func AuthMiddleware(requiredHeaders map[string]string) gin.HandlerFunc { // 校验请求的参数
	return func(context *gin.Context) {
		for key, value := range requiredHeaders {
			if context.GetHeader(key) != value {
				context.JSON(http.StatusUnauthorized,
					GeneralJSONHeader{
						Code: ErrorPermissions,
						Msg:  "Permissions ERROR",
						Path: context.Request.URL.Path,
						Data: nil,
					})
				context.Abort()
				return
			}
		}
		context.Next()
	}
}

func verifyCode(long int64) string { // 随机生成验证码
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, long)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (_ *Utility) UserPasswdVerify(data userLoginData) bool { // 用户密码认证函数
	// userHAS, err := bcrypt.GenerateFromPassword([]byte(data.Username), bcrypt.DefaultCost) // 生成用户名的HAS值
	//if err != nil {
	//	return false
	//}

	DbData, have := dBService.ReadOneDB(DataBase, "user", bson.D{{"name", data.Username}})
	if have {
		utility.Log("is:", DbData)
		if DbData["name"] != data.Username {
			return false
		}
	} else {
		return false
	}
	utility.Log(len(data.Username))
	return true
}

func (_ *Utility) CheckLoginMiddleware() gin.HandlerFunc { // 通过cookie认证
	return func(ctx *gin.Context) {

	}
}

func (_ *Utility) verifyHeaderLoginCode() gin.HandlerFunc { // 通过请求头认证
	return func(ctx *gin.Context) {

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

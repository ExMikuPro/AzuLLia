package service

import (
	"github.com/gin-gonic/gin"
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

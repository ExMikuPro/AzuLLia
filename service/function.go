package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"os"
)

// 中间件

func (_ *Utility) UserPasswdVerify(data userLoginData) bool { // 用户密码认证函数
	// userHAS, err := bcrypt.GenerateFromPassword([]byte(data.Username), bcrypt.DefaultCost) // 生成用户名的HAS值
	//if err != nil {
	//	return false
	//}

	DbData, have := dBService.ReadOneDB(DataBase, "user", bson.D{bson.E{Key: "name", Value: data.Username}})
	if have {
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
		// cookie, err := ctx.Cookie("")
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

func EvnLoad() {
	err := godotenv.Load("Service.env")
	if err != nil {
		utility.Log(err)
	}
}

func GetEvn(key string) string {
	return os.Getenv(key)
}

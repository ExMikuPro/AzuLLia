package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routerDataGroup(context *gin.Context) {

}

func routerApiGroup(context *gin.Context) {

}

func routerMain(router *gin.Engine) {
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, GeneralJSONHeader{
			Code: WorkProgress,
			Msg:  "页面正在施工喵～",
			Path: context.Request.URL.Path,
			Data: nil,
		})
	})
}

func Router(router *gin.Engine) {
	router.NoRoute(NoFoundPage)
	router.Group("/data/", routerDataGroup)
	router.Group("/api/", routerApiGroup)
	routerMain(router) // 主页路由处理
}

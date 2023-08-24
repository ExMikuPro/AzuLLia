package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routerMain(router *gin.Engine) {
	// 功能型接口路由组
	data := router.Group("/data/")
	api := router.Group("/api/")

	api.GET("/paperList", paperListPage) // 全部文章列表

	data.GET("/paperContext/:cid", paperContextPage) // 文章内容

	api.GET("/tagList", tagListPage) // 文章标签列表
	data.GET("/tag/:tid", tagPage)   // 文章标签列表

	// 通用型路由组
	router.GET("/", func(context *gin.Context) { // 主页
		context.JSON(http.StatusOK, GeneralJSONHeader{
			Code: WorkProgress,
			Msg:  "页面正在施工喵～",
			Path: context.Request.URL.Path,
			Data: nil,
		})
	})

	router.GET("/about", aboutPage) // 关于界面

	router.GET("/test", paperListPage) // 测试界面
}

func Router(router *gin.Engine) {
	router.NoRoute(NoFoundPage)
	routerMain(router) // 主页路由处理
}

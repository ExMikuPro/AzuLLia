package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routerMain(router *gin.Engine) {
	// 功能型接口路由组
	data := router.Group(pageAddr["data"])
	api := router.Group(pageAddr["api"])

	api.GET(pageAddr["paperList"], paperListPage) // 全部文章列表
	api.GET(pageAddr["tagList"], tagListPage)     // 文章标签列表

	data.GET(pageAddr["paperContext"], paperContextPage) // 文章内容
	data.GET(pageAddr["tag"], tagPage)                   // 文章标签内容

	// 通用型路由组
	router.GET("/", func(context *gin.Context) { // 主页
		context.JSON(http.StatusOK, GeneralJSONHeader{
			Code: WorkProgress,
			Msg:  "页面正在施工喵～",
			Path: context.Request.URL.Path,
			Data: nil,
		})
	})
	router.GET(pageAddr["about"], aboutPage) // 关于界面

	router.GET(pageAddr["verify"]) // api认证申请接口

	router.POST("/upload", UploadFile) // 配置上传界面位置

	router.GET("/test", testPage) // 测试界面
}

func Router(router *gin.Engine) {
	router.NoRoute(NoFoundPage)
	routerMain(router) // 主页路由处理
}

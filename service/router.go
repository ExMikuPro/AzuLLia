package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func routerMain(router *gin.Engine) {
	// 功能型接口路由组
	data := router.Group(pageAddr["data"])
	api := router.Group(pageAddr["api"])

	api.GET(pageAddr["paperList"], basicPages.PostListPage) // 全部文章列表
	api.GET(pageAddr["tagList"], basicPages.tageListPage)   // 文章标签列表

	data.GET(pageAddr["paperContext"], basicPages.PostContextPage) // 文章内容
	data.GET(pageAddr["tag"], basicPages.tagePage)                 // 文章标签内容

	add := data.Group("/add/") // 添加相关函数
	add.POST("/tage")          // 添加标签接口 todo 添加身份认证

	// 通用型路由组
	router.GET("/", func(context *gin.Context) { // 主页
		context.JSON(http.StatusOK, GeneralJSONHeader{
			Code: WorkProgress,
			Msg:  "页面正在施工喵～",
			Path: context.Request.URL.Path,
			Data: nil,
		})
	})
	router.GET(pageAddr["about"], basicPages.aboutPage) // 关于界面

	router.GET(pageAddr["verify"]) // api认证申请接口

	router.POST("/upload", UploadFile) // 配置上传界面位置

	router.GET("/test", basicPages.aboutPage) // 测试界面
}

func Router(router *gin.Engine) {
	router.NoRoute(basicPages.NoFoundPage)
	routerMain(router) // 主页路由处理
}

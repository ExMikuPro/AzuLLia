package service

import (
	"github.com/gin-gonic/gin"
)

func routerMain(router *gin.Engine) {
	// 功能型接口路由组
	dataGroup := router.Group(pageAddr["data"])
	apiGroup := router.Group(pageAddr["api"])
	addGroup := dataGroup.Group("/add/")       // 添加相关地址
	updateGroup := dataGroup.Group("/update/") // 更新文档相关地址

	apiGroup.GET(pageAddr["paperList"], get.GetPostList) // 全部文章列表
	apiGroup.GET(pageAddr["tagList"], get.GetTagList)    // 文章标签列表
	apiGroup.GET("/typeList", get.GetTypeList)           // 文章标签列表
	apiGroup.POST("/userLogin", user.UserLogin)          // 用户登陆接口

	dataGroup.GET(pageAddr["paperContext"], get.GetPostContext) // 文章内容
	dataGroup.GET(pageAddr["tag"], get.GetTag)                  // 文章标签内容

	addGroup.POST("/tag", add.AddTag)     // 添加标签接口 todo 添加身份认证
	addGroup.POST("/type", add.AddType)   // 添加分类接口
	addGroup.POST("/group", add.AddGroup) // 添加用户组接口
	addGroup.POST("/user", add.AddUser)   // 添加用户接口

	updateGroup.POST("/tag", update.UpdateTag)     // 更新标签接口
	updateGroup.POST("/type", update.UpdateType)   // 更新类别接口
	updateGroup.POST("/group", update.UpdateGroup) // 更新用户组接口
	updateGroup.POST("/user", update.UpdateUser)   // 更新用户接口

	// 通用型路由组
	router.GET("/", get.MainPage)               // 主页页面
	router.GET(pageAddr["about"], get.GetAbout) // 关于界面

	router.GET(pageAddr["verify"]) // api认证申请接口

	router.POST("/upload", UploadFile) // 配置上传界面位置

	router.GET("/test", get.GetAbout) // 测试界面
}

func Router(router *gin.Engine) {
	router.NoRoute(utility.NoFoundPage)
	routerMain(router) // 主页路由处理
}

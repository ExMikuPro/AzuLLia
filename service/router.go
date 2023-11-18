package service

import (
	"github.com/gin-gonic/gin"
)

func routerMain(router *gin.Engine) {
	// 功能型接口路由组
	listGroup := router.Group("/list/")         // 列表
	contentGroup := router.Group("/content/")   // 内容
	adminGroup := router.Group("/admin/")       // 管理
	addGroup := adminGroup.Group("/add/")       // 添加
	updateGroup := adminGroup.Group("/update/") // 更新

	// 添加中间键
	adminGroup.Use(utility.CheckLoginMiddleware())  // 登陆认证
	adminGroup.Use(utility.verifyHeaderLoginCode()) // 请求头认证

	listGroup.GET("/article", get.ArticleList)   // 文章列表
	listGroup.GET("/tag", get.TagList)           // 标签列表
	listGroup.GET("/category", get.CategoryList) // 分类列表

	contentGroup.GET("/article/:id", get.GetArticle)   // 文章内容
	contentGroup.GET("/tag/:id", get.GetTag)           // 标签内容
	contentGroup.GET("/category/:id", get.GetCategory) // 标签内容

	addGroup.POST("/tag", add.AddTag)           // 添加标签
	addGroup.POST("/category", add.AddCategory) // 添加分类
	addGroup.POST("/group", add.AddGroup)       // 添加用户组
	addGroup.POST("/user", add.AddUser)         // 添加用户
	addGroup.POST("/article", add.AddArticle)   // 添加文章

	updateGroup.POST("/tag", update.UpdateTag)          // 更新标签
	updateGroup.POST("/category", update.UpdateArticle) // 更新类别
	updateGroup.POST("/group", update.UpdateGroup)      // 更新用户组
	updateGroup.POST("/user", update.UpdateUser)        // 更新用户

	// 通用型路由组
	router.GET("/", get.MainPage) // 主页

	router.GET("/about", get.GetAbout) // 关于

	router.POST("/upload", UploadFile) // 上传

	router.POST("/userLogin", user.UserLogin) // 登陆

	router.GET("/test", test) // 测试

}

func Router(router *gin.Engine) {
	router.NoRoute(utility.NoFoundPage) // 404错误
	routerMain(router)                  // 主路由
}

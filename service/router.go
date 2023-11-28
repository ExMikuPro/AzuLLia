package service

import (
	"github.com/gin-gonic/gin"
)

func routerMain(router *gin.Engine) {
	// 功能型接口路由组
	listGroup := router.Group("/list/")         // 列表
	getGroup := router.Group("/get/")           // 内容
	adminGroup := router.Group("/admin/")       // 管理
	addGroup := adminGroup.Group("/add/")       // 添加
	updateGroup := adminGroup.Group("/update/") // 更新
	deleteGroup := adminGroup.Group("/delete/") // 删除

	// 添加中间键
	adminGroup.Use(utilityFunction.CheckLoginMiddleware())  // 登陆认证
	adminGroup.Use(utilityFunction.verifyHeaderLoginCode()) // 请求头认证

	listGroup.GET("/article", getFunction.ArticleList)   // 文章列表
	listGroup.GET("/tag", getFunction.TagList)           // 标签列表
	listGroup.GET("/category", getFunction.CategoryList) // 分类列表

	getGroup.GET("/article/:id", getFunction.GetArticle)   // 文章内容
	getGroup.GET("/tag/:id", getFunction.GetTag)           // 标签内容
	getGroup.GET("/category/:id", getFunction.GetCategory) // 标签内容

	addGroup.POST("/tag", addFunction.AddTag)           // 添加标签
	addGroup.POST("/category", addFunction.AddCategory) // 添加分类
	addGroup.POST("/group", addFunction.AddGroup)       // 添加用户组
	addGroup.POST("/user", addFunction.AddUser)         // 添加用户
	addGroup.POST("/article", addFunction.AddArticle)   // 添加文章

	updateGroup.POST("/tag", updateFunction.UpdateTag)          // 更新标签
	updateGroup.POST("/category", updateFunction.UpdateArticle) // 更新类别
	updateGroup.POST("/group", updateFunction.UpdateGroup)      // 更新用户组
	updateGroup.POST("/user", updateFunction.UpdateUser)        // 更新用户

	deleteGroup.POST("/tag", deleteFunction.DeleteTag)           // 删除标签
	deleteGroup.POST("/category", deleteFunction.DeleteCategory) // 删除类别
	deleteGroup.POST("/group", deleteFunction.DeleteGroup)       // 删除用户组
	deleteGroup.POST("/user", deleteFunction.DeleteUser)         // 删除用户
	deleteGroup.POST("/article", deleteFunction.DeleteArticle)   // 删除文章
	// 通用型路由组
	router.GET("/", getFunction.MainPage) // 主页

	router.GET("/about", getFunction.GetAbout) // 关于

	router.POST("/upload", UploadFile) // 上传

	router.POST("/userLogin", userFunction.UserLogin) // 登陆

	router.GET("/test", test) // 测试

}

func Router(router *gin.Engine) {
	router.NoRoute(utilityFunction.NoFoundPage) // 404错误
	routerMain(router)                          // 主路由
}

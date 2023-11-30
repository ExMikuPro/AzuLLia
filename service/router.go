package service

import (
	"github.com/gin-gonic/gin"
)

func routerMain(router *gin.Engine) {
	// 功能型接口路由组
	adminGroup := router.Group("/admin/") // 管理
	apiGroup := router.Group("/api/")     // 外部调用接口
	v1Group := apiGroup.Group("/v1/")     // 版本号

	// 添加中间键
	router.Use(utilityFunction.ReturnHeader())              // 响应头中间件
	adminGroup.Use(utilityFunction.CheckLoginMiddleware())  // 登陆认证
	adminGroup.Use(utilityFunction.verifyHeaderLoginCode()) // 请求头认证

	v1Group.GET("/article", getFunction.ArticleList)   // 文章列表
	v1Group.GET("/tag", getFunction.TagList)           // 标签列表
	v1Group.GET("/category", getFunction.CategoryList) // 分类列表

	v1Group.GET("/article/:id", getFunction.GetArticle)   // 文章内容
	v1Group.GET("/tag/:id", getFunction.GetTag)           // 标签内容
	v1Group.GET("/category/:id", getFunction.GetCategory) // 标签内容

	v1Group.POST("/tag", addFunction.AddTag)           // 添加标签
	v1Group.POST("/category", addFunction.AddCategory) // 添加分类
	v1Group.POST("/group", addFunction.AddGroup)       // 添加用户组
	v1Group.POST("/user", addFunction.AddUser)         // 添加用户
	v1Group.POST("/article", addFunction.AddArticle)   // 添加文章

	v1Group.PUT("/tag", updateFunction.UpdateTag)         // 更新标签
	v1Group.PUT("/article", updateFunction.UpdateArticle) // 更新文章
	v1Group.PUT("/group", updateFunction.UpdateGroup)     // 更新用户组
	v1Group.PUT("/user", updateFunction.UpdateUser)       // 更新用户
	v1Group.PUT("/category")                              // 更新分类

	v1Group.DELETE("/tag", deleteFunction.DeleteTag)           // 删除标签
	v1Group.DELETE("/category", deleteFunction.DeleteCategory) // 删除类别
	v1Group.DELETE("/group", deleteFunction.DeleteGroup)       // 删除用户组
	v1Group.DELETE("/user", deleteFunction.DeleteUser)         // 删除用户
	v1Group.DELETE("/article", deleteFunction.DeleteArticle)   // 删除文章

	v1Group.POST("/userLogin", userFunction.UserLogin) // 用户登陆

	// 通用型路由组
	router.GET("/", getFunction.MainPage) // 主页

	router.GET("/about", getFunction.GetAbout) // 关于

	router.POST("/upload", UploadFile) // 上传

	router.GET("/test", test) // 测试

}

func Router(router *gin.Engine) {
	router.NoRoute(utilityFunction.NoFoundPage) // 404错误
	routerMain(router)                          // 主路由
}

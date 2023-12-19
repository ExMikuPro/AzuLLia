package main

import (
	_ "AzuLLia/docs"
	"AzuLLia/service"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title AzuLLia
// @version 1.0.1206
// @description 一只轻量级无头CMS
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	service.ServerBegin()                   // 初始化运行环境
	gin.SetMode(service.GetEvn("GIN_MODE")) //  设置运行模式
	router := gin.Default()                 // 注册默认静态路由
	router.MaxMultipartMemory = 8 << 20     // 设置默认上传大小限制为8MB
	service.Router(router)                  // 注册路由服务器
	service.StaticFile(router)              // 注册静态资源服务器
	if gin.Mode() == "debug" {              // 开发文档
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	err := router.Run(service.GetEvn("WEBSITE_PORT")) // 启动服务器
	if err != nil {
		fmt.Println("Server Start ERROR:", err)
	}
}

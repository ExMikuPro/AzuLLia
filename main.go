package main

import (
	_ "AzuLLia/docs"
	"AzuLLia/service"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"io"
	"os"
)

// @title AzuLLia
// @version 1.0.1206
// @description 一只轻量级无头CMS
// @host 127.0.0.1:8080
// @BasePath /
func main() {
	if os.Getenv("RunMode") != "" {
		gin.SetMode(gin.ReleaseMode)   // 设置模式为发行模式
		gin.DefaultWriter = io.Discard // 关闭控制台日志输出
	}

	router := gin.Default() // 注册默认静态路由
	// router.Use(service.AuthMiddleware(service.RequiredHeaders))
	router.MaxMultipartMemory = 8 << 20 // 设置默认上传大小限制为8MB
	service.Router(router)              // 注册路由服务器
	service.StaticFile(router)          // 注册静态资源服务器
	if os.Getenv("RunMode") == "" {     // 如歌哦巍峨
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	if os.Getenv("ServerPort") != "" {
		err := router.Run(os.Getenv("ServerPort")) // 启动服务器
		if err != nil {
			fmt.Println("Server Start ERROR:", err)
		}
	} else {
		err := router.Run(":8080") // 启动测试服务器
		if err != nil {
			fmt.Println("Server Start ERROR:", err)
		}
	}

}

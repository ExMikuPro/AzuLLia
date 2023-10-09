package main

import (
	"GoK-on/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // 注册默认静态路由
	// router.Use(service.AuthMiddleware(service.RequiredHeaders))
	service.Router(router)     // 注册路由服务器
	service.StaticFile(router) // 注册静态资源服务器
	err := router.Run(":82")   // 启动服务器
	if err != nil {
		fmt.Println("Server Start ERROR:", err)
		return
	}
}

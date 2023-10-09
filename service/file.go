package service

import "github.com/gin-gonic/gin"

// 静态资源处理函数

func StaticFile(router *gin.Engine) {
	// 注册静态文件存放位置
	router.Static("/file", "./staticFile")
	// 加载标签图标
	router.StaticFile("/favicon.ico", "./staticFile/favicon.ico")
}

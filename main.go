package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "页面正在施工喵～",
			"data": "[]",
		})
	})
}

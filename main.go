package main

import (
	"GoK-on/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.Use(service.AuthMiddleware(service.RequiredHeaders))
	service.Router(router)
	err := router.Run(":82")
	if err != nil {
		fmt.Println("Server Start ERROR:", err)
		return
	}
}

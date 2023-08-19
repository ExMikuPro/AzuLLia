package main

import (
	"GoK-on/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	service.Router(router)
	err := router.Run(":82")
	if err != nil {
		fmt.Println("ERROR:server", err)
		return
	}
}

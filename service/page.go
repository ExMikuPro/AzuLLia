package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoFoundPage(c *gin.Context) { // 设置404页面
	c.JSON(http.StatusNotFound, GeneralJSONHeader{
		Code: http.StatusNotFound,
		Msg:  "页面走丢了~",
		Path: c.Request.URL.Path,
		Data: nil,
	})
}

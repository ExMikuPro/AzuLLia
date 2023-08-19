package service

import "github.com/gin-gonic/gin"

const WorkProgress = 520

type GeneralJSONHeader struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Path string `json:"path"`
	Data gin.H  `json:"data"`
}

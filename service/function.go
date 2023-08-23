package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件

func AuthMiddleware(requiredHeaders map[string]string) gin.HandlerFunc { // 校验请求的参数
	return func(context *gin.Context) {
		for key, value := range requiredHeaders {
			if context.GetHeader(key) != value {
				context.JSON(http.StatusUnauthorized,
					GeneralJSONHeader{
						Code: ErrorPermissions,
						Msg:  "Permissions ERROR",
						Path: context.Request.URL.Path,
						Data: nil,
					})
				context.Abort()
				return
			}
		}
		context.Next()
	}
}

// 数据库处理函数

var DataBase = InitDB() // 初始化数据库

func aboutPage(context *gin.Context) { // 关于界面
	context.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: context.Request.URL.Path,
		Data: gin.H{
			"version": Version,
		},
	})
}

func paperListPage(context *gin.Context) { // 获取文章列表
	data := ReadAllDB(DataBase, "paperList")
	context.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: context.Request.URL.Path,
		Data: gin.H{
			"length": len(data), // 列表长度
			"list":   data,      // 列表内容
		},
	})
}

func paperContextPage(context *gin.Context) {
	from := context.Query("from")
	context.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: context.Request.URL.Path,
		Data: gin.H{
			"isExist": true,          // todo 判断文章是否存在
			"from":    from,          // 查找对应的文章
			"context": "<p>文章主题</p>", // todo 文章主体内容
		},
	})
}

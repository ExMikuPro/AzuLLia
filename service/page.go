package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	// todo 添加api认证
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

func paperContextPage(context *gin.Context) { // 获取文章内容
	from := context.Query("from") // 文章对应的内容 todo 使用数据库原生id
	// cid := context.Query("cid") // 文章cid

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

func tagListPage(context *gin.Context) { // 标签列表
	data := ReadAllDB(DataBase, "tagList")
	context.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: context.Request.URL.Path,
		Data: gin.H{
			"length": len(data), // 列表长度
			"list":   data,      // 列表数据
		},
	})
}

func tagPage(context *gin.Context) { // 查询单个标签内容
	filter := bson.D{
		{"tid", context.Param("tid")}, // tid为String
	}
	data := ReadOneDB(DataBase, "tagList", filter)
	context.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: context.Request.URL.Path,
		Data: data,
	})
}

func testPage(context *gin.Context) { // 设置临时测试界面
	context.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: context.Request.URL.Path,
		Data: gin.H{
			"message": "如你所见这是个测试界面awa",
			"code":    verifyCode(15),
		},
	})
}

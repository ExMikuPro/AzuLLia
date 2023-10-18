package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func (_ *BasicPage) aboutPage(ctx *gin.Context) { // 框架版本界面
	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: ctx.Request.URL.Path,
		Data: gin.H{
			"version": Version,
		},
	})
}

func (_ *BasicPage) NoFoundPage(ctx *gin.Context) { // 404
	ctx.JSON(http.StatusNotFound, GeneralJSONHeader{
		Code: http.StatusNotFound,
		Msg:  "页面走丢了~",
		Path: ctx.Request.URL.Path,
		Data: nil,
	})
}

func (_ *BasicPage) PostListPage(ctx *gin.Context) { // 文章列表界面
	data := ReadAllDB(DataBase, "paperList")
	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: ctx.Request.URL.Path,
		Data: gin.H{
			"length": len(data), // 列表长度
			"list":   data,      // 列表内容
		},
	})
}

func (_ *BasicPage) PostContextPage(ctx *gin.Context) { // 文章内容页面
	from := ctx.Query("from") // 文章对应的内容 todo 使用数据库原生id
	// cid := context.Query("cid") // 文章cid
	// todo 建立文章数据表

	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: ctx.Request.URL.Path,
		Data: gin.H{
			"isExist": true,          // todo 判断文章是否存在
			"from":    from,          // 查找对应的文章
			"context": "<p>文章主题</p>", // todo 文章主体内容
		},
	})
}

func (_ *BasicPage) tageListPage(ctx *gin.Context) { // 标签列表界面
	data := ReadAllDB(DataBase, "tagList")
	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: ctx.Request.URL.Path,
		Data: gin.H{
			"length": len(data), // 列表长度
			"list":   data,      // 列表数据
		},
	})
}

func (_ *BasicPage) tagePage(ctx *gin.Context) {
	filter := bson.D{
		{"tid", ctx.Param("tid")}, // tid为String
	}
	data, isExist := ReadOneDB(DataBase, "tage", filter)
	if isExist { // 判断标签是否存在
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: data,
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: NoDocuments,
			Msg:  "DB No Documents",
			Path: ctx.Request.URL.Path,
			Data: data,
		})
	}
}

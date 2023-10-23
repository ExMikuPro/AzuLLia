package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
	"time"
)

// 静态资源处理函数

func StaticFile(router *gin.Engine) {
	// 注册静态文件存放位置
	router.Static("/src", "./staticFile")
	// 加载标签图标
	router.StaticFile("/favicon.ico", "./staticFile/favicon.ico")
}

func UploadFile(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  err.Error(),
			Path: context.Request.URL.Path,
			Data: nil,
		})
		return
	}

	ext := filepath.Ext(file.Filename) // 获取文件拓展名

	var isFileType = false
	for _, allowedExt := range FileType { // 判断文件类型
		if ext == allowedExt {
			isFileType = true
		}
	}

	if isFileType { // 如果文件类型允许上传
		newFileName := uuid.New().String() + ext
		err = context.SaveUploadedFile(file, pageAddr["uploadFile"]+newFileName) // 将文件名修改为UUID并保存
		if err != nil {
			context.JSON(http.StatusOK, GeneralJSONHeader{
				Code: FileSaveError,
				Msg:  err.Error(),
				Path: context.Request.URL.Path,
				Data: nil,
			})
			return
		}
		context.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: context.Request.URL.Path,
			Data: gin.H{
				"OldFileName": file.Filename,
				"FileName":    newFileName,
				"Date":        time.Now().Unix(),
				"Size":        file.Size,
			},
		})
	} else { // 如果文件不可上传
		context.JSON(http.StatusOK, GeneralJSONHeader{
			Code: UnKnownFileType,
			Msg:  "Unknown file type",
			Path: context.Request.URL.Path,
			Data: nil,
		})
		return
	}

}

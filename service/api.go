package service

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"
)

func (_ *Get) GetAbout(ctx *gin.Context) { // 框架版本界面
	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: ctx.Request.URL.Path,
		Data: gin.H{
			"version": Version,
		},
	})
}

func (_ *Get) GetPostList(ctx *gin.Context) { // 文章列表界面
	data := dBService.ReadAllDB(DataBase, "paperList")
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

func (_ *Get) GetPostContext(ctx *gin.Context) { // 文章内容页面
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

func (_ *Get) GetTagList(ctx *gin.Context) { // 标签列表界面
	data := dBService.ReadAllDB(DataBase, "tag")
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

func (_ *Get) GetTypeList(ctx *gin.Context) { // 分类列表界面
	data := dBService.ReadAllDB(DataBase, "type")
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

func (_ *Get) GetTag(ctx *gin.Context) {
	tid, _ := primitive.ObjectIDFromHex(ctx.Param("id"))
	filter := bson.D{
		{"_id", tid},
	}
	data, isExist := dBService.ReadOneDB(DataBase, "tag", filter)
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

func (_ *Get) MainPage(ctx *gin.Context) { // 主页页面
	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: WorkProgress,
		Msg:  "页面正在施工喵～",
		Path: ctx.Request.URL.Path,
		Data: nil,
	})
}

func (_ *Add) AddTag(ctx *gin.Context) { // 添加新标签
	order, _ := strconv.Atoi(ctx.PostForm("order"))
	tageData := tageTable{
		ID:          primitive.NewObjectID(),     // 分配新_id
		Name:        ctx.PostForm("name"),        // 名称
		Slug:        ctx.PostForm("slug"),        // 项目缩略名
		Type:        ctx.PostForm("type"),        // 项目类型
		Description: ctx.PostForm("description"), // 选项描述
		Count:       0,                           //
		Order:       order,                       // 项目排序
	}

	tid, isOK := dBService.WriteOneDB(DataBase, "tag", tageData)
	if isOK {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: gin.H{
				"tagInfo": tageData,
				"tid":     tid,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: DBWriteError,
			Msg:  "DB Write Error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
	}
}

func (_ *Add) AddType(ctx *gin.Context) {
	order, _ := strconv.Atoi(ctx.PostForm("order"))
	typeData := typeTable{
		ID:          primitive.NewObjectID(),     // 分配新_id
		Name:        ctx.PostForm("name"),        // 名称
		Slug:        ctx.PostForm("slug"),        // 项目缩略名
		Description: ctx.PostForm("description"), // 选项描述
		Count:       0,                           //
		Order:       order,                       // 项目排序
	}
	tid, isOK := dBService.WriteOneDB(DataBase, "type", typeData)
	if isOK {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: gin.H{
				"typeInfo": typeData,
				"tid":      tid,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: DBWriteError,
			Msg:  "DB Write Error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
	}
}

func (_ *Add) AddGroup(ctx *gin.Context) {
	userGroup := userGroupTable{
		ID:   primitive.NewObjectID(),
		Name: ctx.PostForm("name"),
	}
	tid, isOK := dBService.WriteOneDB(DataBase, "group", userGroup)
	if isOK {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: gin.H{
				"groupInfo": userGroup,
				"tid":       tid,
			},
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: DBWriteError,
			Msg:  "DB Write Error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
	}
}

func (_ *Add) AddUser(ctx *gin.Context) { // 添加用户
	userData := userTable{
		ID:         primitive.NewObjectID(),
		Name:       ctx.PostForm("name"),
		Passwd:     ctx.PostForm("passwd"),
		Mail:       ctx.PostForm("mail"),
		Url:        ctx.PostForm("url"),
		ScreenName: ctx.PostForm("screenName"),
		Created:    time.Now().Unix(),
		UpData:     time.Now().Unix(),
		Group:      ctx.PostForm("group"),
	}
	uid, isOK := dBService.WriteOneDB(DataBase, "user", userData)
	if isOK {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: gin.H{
				"userInfo": userData,
				"uid":      uid,
			},
		})
	}
}

func (_ *User) UserLogin(ctx *gin.Context) {
	var userLogin userLoginData
	err := ctx.ShouldBindJSON(&userLogin)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: UserInformationVerifyError,
			Msg:  "user information verify error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	if utility.UserPasswdVerify(userLogin) { // 判断是否符合数据库内数据
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: gin.H{
				"status": "success",
			},
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: UserInformationVerifyError,
			Msg:  "user information verify error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
	}
}

func (_ *Update) UpdateUser(ctx *gin.Context) { // todo 前端注意请求值
	uid, _ := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	passwd := md5.Sum([]byte(ctx.PostForm("passwd")))
	passwdMD5 := hex.EncodeToString(passwd[:])
	mail := ctx.PostForm("mail")
	url := ctx.PostForm("url")
	screenName := ctx.PostForm("screenName")
	group := ctx.PostForm("group")
	filter := bson.D{{"_id", uid}}
	update := bson.D{
		{"$set", bson.D{
			{"name", name},
			{"passwd", passwdMD5},
			{"mail", mail},
			{"url", url},
			{"screenName", screenName},
			{"upData", time.Now().Unix()},
			{"group", group},
		}},
	}

	data, isOK := dBService.UpdateOneDB(DataBase, "user", filter, update)
	if isOK {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: data,
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: DBUpdateError,
			Msg:  "DB Update Error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
	}
}

func (_ *Update) UpdateTag(ctx *gin.Context) {
	tid, _ := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	slug := ctx.PostForm("slug")
	types := ctx.PostForm("type")
	description := ctx.PostForm("description")
	order, _ := strconv.Atoi(ctx.PostForm("order"))

	filter := bson.D{{"_id", tid}}
	update := bson.D{
		{"$set", bson.D{
			{"name", name},
			{"slug", slug},
			{"type", types},
			{"description", description},
			{"order", order},
		}},
	}

	data, isOK := dBService.UpdateOneDB(DataBase, "tag", filter, update)
	if isOK {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: data,
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: DBUpdateError,
			Msg:  "DB Update Error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
	}
}

func (_ *Update) UpdateType(ctx *gin.Context) {
	tid, _ := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	slug := ctx.PostForm("slug")
	description := ctx.PostForm("description")
	order, _ := strconv.Atoi(ctx.PostForm("order"))

	filter := bson.D{{"_id", tid}}
	update := bson.D{
		{"$set", bson.D{
			{"name", name},
			{"slug", slug},
			{"description", description},
			{"order", order},
		}},
	}

	data, isOK := dBService.UpdateOneDB(DataBase, "type", filter, update)
	if isOK {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: data,
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: DBUpdateError,
			Msg:  "DB Update Error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
	}
}

func (_ *Update) UpdateGroup(ctx *gin.Context) {
	tid, _ := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	name := ctx.PostForm("name")

	filter := bson.D{{"_id", tid}}
	update := bson.D{
		{"$set", bson.D{
			{"name", name},
		}},
	}

	data, isOK := dBService.UpdateOneDB(DataBase, "group", filter, update)
	if isOK {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: data,
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: DBUpdateError,
			Msg:  "DB Update Error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
	}
}

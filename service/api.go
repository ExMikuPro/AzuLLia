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

// GetAbout @Title 系统
// @Tags 系统
// @Summary	获取框架版本信息
// @Produce	json
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/about [GET]
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

// ArticleList @Title 文章
// @Tags 文章
// @Summary	获取列表
// @Produce	json
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/list/article [GET]
func (_ *Get) ArticleList(ctx *gin.Context) { // 文章列表界面
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

// GetArticle @Title 文章
// @Tags 文章
// @Summary	获取文章内容
// @Produce	json
// @Param id path string true "文章ID"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/content/article/{id} [GET]
func (_ *Get) GetArticle(ctx *gin.Context) { // 文章内容页面
	cid, _ := primitive.ObjectIDFromHex(ctx.Param("id")) // 文章cid

	filter := bson.D{
		{Key: "_id", Value: cid},
	}

	data, isExist := dBService.ReadOneDB(DataBase, "article", filter)

	if isExist {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: data,
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: gin.H{},
		})
	}
}

// TagList @Title 标签
// @Tags 标签
// @Summary	获取标签列表
// @Produce	json
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/list/tag [GET]
func (_ *Get) TagList(ctx *gin.Context) { // 标签列表界面
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

// CategoryList @Title 分类
// @Tags 分类
// @Summary	获取分类列表
// @Produce	json
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/list/category [GET]
func (_ *Get) CategoryList(ctx *gin.Context) { // 分类列表界面
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

// GetType @Title 分类
// @Tags 分类
// @Summary	获取分类
// @Produce	json
// @Param id path string true "分类ID"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/content/category/{id} [GET]
func (_ *Get) GetType(ctx *gin.Context) {
	cid, _ := primitive.ObjectIDFromHex(ctx.Param("id"))

	filter := bson.D{
		{Key: "_id", Value: cid},
	}
	data, isExist := dBService.ReadOneDB(DataBase, "type", filter)
	if isExist {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: data,
		})
	} else {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: gin.H{},
		})
	}
}

// GetTag @Title 标签
// @Tags 标签
// @Summary	获取标签
// @Produce	json
// @Param id path string true "标签ID"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/content/tag/{id} [GET]
func (_ *Get) GetTag(ctx *gin.Context) {
	tid, _ := primitive.ObjectIDFromHex(ctx.Param("id"))
	filter := bson.D{
		{Key: "_id", Value: tid},
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

// AddTag @Title 标签
// @Tags 标签
// @Summary	添加标签
// @Produce	json
// @Param name formData string true "名称"
// @Param slug formData string true "缩略名"
// @Param type formData string true "类型"
// @Param description formData string true "描述"
// @Param order formData int true "排序"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/admin/add/tag [POST]
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

// AddArticle @Title 文章
// @Tags 文章
// @Summary	添加文章
// @Produce	json
// @Consumes json
// @Param title body string true "标题"
// @Param slug body string true "标题缩略名"
// @Param text body string true "内容文字"
// @Param order body int true "排序"
// @Param authorId body string true "内容所属用户"
// @Param template body string true "模版文件"
// @Param type body string true "内容类别"
// @Param status body string true "内容状态"
// @Param password body string true "密码"
// @Param allowComment body bool true "是否允许评论"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/admin/add/article [POST]
func (_ *Add) AddArticle(ctx *gin.Context) {
	var postData contextTable
	err := ctx.ShouldBindJSON(&postData)
	postData.ID = primitive.NewObjectID()
	postData.Created = time.Now().Unix()
	postData.Modified = time.Now().Unix()
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerDeCodeError,
			Msg:  "Server Error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: ctx.Request.URL.Path,
		Data: gin.H{
			"postID": postData.ID,
		},
	})

}

// AddCategory @Title 分类
// @Tags 分类
// @Summary	添加分类
// @Produce	json
// @Param name formData string true "名称"
// @Param slug formData string true "缩略名"
// @Param description formData string true "描述"
// @Param order formData int true "排序"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/admin/add/category [POST]
func (_ *Add) AddCategory(ctx *gin.Context) {
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

// AddGroup @Title 用户
// @Tags 用户
// @Summary	添加用户组
// @Produce	json
// @Param name formData string true "名称"
// @Param order formData int true "排序"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/admin/add/group [POST]
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

// AddUser @Title 用户
// @Tags 用户
// @Summary	添加用户
// @Produce	json
// @Param name formData string true "名称"
// @Param passwd formData string true "密码"
// @Param mail formData string true "邮箱"
// @Param url formData string true "网站地址"
// @Param screenName formData string true "用户昵称"
// @Param group formData string true "用户组"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/admin/add/user [POST]
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

func (_ *Update) UpdateUser(ctx *gin.Context) {
	uid, _ := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	passwd := md5.Sum([]byte(ctx.PostForm("passwd")))
	passwdMD5 := hex.EncodeToString(passwd[:])
	mail := ctx.PostForm("mail")
	url := ctx.PostForm("url")
	screenName := ctx.PostForm("screenName")
	group := ctx.PostForm("group")
	filter := bson.D{{Key: "_id", Value: uid}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "passwd", Value: passwdMD5},
			{Key: "mail", Value: mail},
			{Key: "url", Value: url},
			{Key: "screenName", Value: screenName},
			{Key: "upData", Value: time.Now().Unix()},
			{Key: "group", Value: group},
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

	filter := bson.D{{Key: "_id", Value: tid}}
	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "slug", Value: slug},
			{Key: "type", Value: types},
			{Key: "description", Value: description},
			{Key: "order", Value: order},
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

func (_ *Update) UpdateArticle(ctx *gin.Context) {
	tid, _ := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	slug := ctx.PostForm("slug")
	description := ctx.PostForm("description")
	order, _ := strconv.Atoi(ctx.PostForm("order"))

	filter := bson.D{{Key: "_id", Value: tid}}
	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "slug", Value: slug},
			{Key: "description", Value: description},
			{Key: "order", Value: order},
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

	filter := bson.D{{Key: "_id", Value: tid}}
	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
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

func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"context": "<p>hello</p>",
	})
}

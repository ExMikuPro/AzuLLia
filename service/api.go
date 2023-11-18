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
	data, err := DataBase.ReadAllDB("paperList", []gin.H{})
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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

	data, err := DataBase.ReadOneDB("article", filter, gin.H{})

	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: ctx.Request.URL.Path,
		Data: data,
	})
}

// TagList @Title 标签
// @Tags 标签
// @Summary	获取标签列表
// @Produce	json
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/list/tag [GET]
func (_ *Get) TagList(ctx *gin.Context) { // 标签列表界面
	data, err := DataBase.ReadAllDB("tag", []gin.H{})
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
	data, err := DataBase.ReadAllDB("type", []gin.H{})
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"length": len(data), // 列表长度
			"list":   data,      // 列表数据
		},
	})
}

// GetCategory @Title 分类
// @Tags 分类
// @Summary	获取分类
// @Produce	json
// @Param id path string true "分类ID"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/content/category/{id} [GET]
func (_ *Get) GetCategory(ctx *gin.Context) {
	cid, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	filter := bson.D{
		{Key: "_id", Value: cid},
	}
	data, err := DataBase.ReadOneDB("type", filter, gin.H{})
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: ctx.Request.URL.Path,
		Data: data,
	})
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
	data, err := DataBase.ReadOneDB("tag", filter, gin.H{})
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, GeneralJSONHeader{
		Code: SuccessCode,
		Msg:  "success",
		Path: ctx.Request.URL.Path,
		Data: data,
	})
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

	err := DataBase.WriteOneDB("tag", tageData)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"tagInfo": tageData,
		},
	})
}

// AddArticle @Title 文章
// @Tags 文章
// @Summary	添加文章
// @Param request_body body articleTable true "JSON数据"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router /admin/add/article [POST]
func (_ *Add) AddArticle(ctx *gin.Context) {
	var articleData articleTable
	err := ctx.ShouldBindJSON(&articleData)
	articleData.ID = primitive.NewObjectID()
	articleData.Created = time.Now().Unix()
	articleData.Modified = time.Now().Unix()
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}

	err = DataBase.WriteOneDB("article", articleData)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"postID": articleData.ID,
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
	err := DataBase.WriteOneDB("type", typeData)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"typeInfo": typeData,
		},
	})
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
	err := DataBase.WriteOneDB("group", userGroup)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"groupInfo": userGroup,
		},
	})
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
	err := DataBase.WriteOneDB("user", userData)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"userInfo": userData,
		},
	})

}

// UserLogin @Title 用户
// @Tags 用户
// @Summary	用户登陆
// @Produce	json
// @Param request_body body userLoginData true "JSON数据"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/userLogin [POST]
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

	err := DataBase.UpdateOneDB("user", filter, update)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"id": uid,
		},
	})
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

	err := DataBase.UpdateOneDB("tag", filter, update)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"id": tid,
		},
	})
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

	err := DataBase.UpdateOneDB("type", filter, update)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"id": tid,
		},
	})
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

	err := DataBase.UpdateOneDB("group", filter, update)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
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
			"id": tid,
		},
	})
}

func test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"context": "<p>hello</p>",
	})
}

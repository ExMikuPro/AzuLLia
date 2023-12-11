package service

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

// GetAbout @Title 系统
// @Tags 系统
// @Summary	获取框架版本信息
// @Description 获取框架版本信息
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
// @Description 获取全部文章列表
// @Produce	json
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/article [GET]
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
// @Description 获取指定id文章列表
// @Produce	json
// @Param id path string true "文章ID"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/article/{id} [GET]
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
// @Description 获取全部标签列表
// @Produce	json
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/tag [GET]
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
// @Description 获取全部分类列表
// @Produce	json
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/category [GET]
func (_ *Get) CategoryList(ctx *gin.Context) { // 分类列表界面
	data, err := DataBase.ReadAllDB("category", []gin.H{})
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
// @Description 获取指定id分类列表
// @Produce	json
// @Param id path string true "分类ID"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/category/{id} [GET]
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
	data, err := DataBase.ReadOneDB("category", filter, gin.H{})
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
// @Description 获取全部标签列表
// @Produce	json
// @Param id path string true "标签ID"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/tag/{id} [GET]
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
// @Description 添加新标签(需要相关权限token)
// @Produce	json
// @Param name formData string true "名称"
// @Param slug formData string true "缩略名"
// @Param type formData string true "类型"
// @Param description formData string true "描述"
// @Param order formData int true "排序"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/tag [POST]
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
// @Produce	json
// @Description 添加新文章(需要相关权限token)
// @Param request_body body articleTable true "JSON数据"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router /api/v1/article [POST]
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
// @Description 添加新分类(需要相关权限token)
// @Param name formData string true "名称"
// @Param slug formData string true "缩略名"
// @Param description formData string true "描述"
// @Param order formData int true "排序"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/category [POST]
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
	err := DataBase.WriteOneDB("category", typeData)
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

// AddGroup @Title 用户组
// @Tags 用户组
// @Summary	添加用户组
// @Produce	json
// @Description 添加新用户组(需要相关权限token)
// @Param name formData string true "名称"
// @Param order formData int true "排序"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/group [POST]
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

// GetUserList @Title 用户
// @Tags 用户
// @Summary	用户列表
// @Produce	json
// @Description 获取全部用户列表(需要相关权限token)
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/user [GET]
func (_ *Get) GetUserList(ctx *gin.Context) { // 获取用户列表
	data, err := DataBase.ReadAllDB("user", []gin.H{})
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
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

// AddUser @Title 用户
// @Tags 用户
// @Summary	添加用户
// @Produce	json
// @Description 添加新用户(需要相关权限token)
// @Param name formData string true "名称"
// @Param passwd formData string true "密码"
// @Param mail formData string true "邮箱"
// @Param url formData string true "网站地址"
// @Param screenName formData string true "用户昵称"
// @Param group formData string true "用户组"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/user [POST]
func (_ *Add) AddUser(ctx *gin.Context) { // 添加用户
	passwd, err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("passwd")), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	userData := userTable{
		ID:         primitive.NewObjectID(),
		Name:       ctx.PostForm("name"),
		Passwd:     string(passwd),
		Mail:       ctx.PostForm("mail"),
		Url:        ctx.PostForm("url"),
		ScreenName: ctx.PostForm("screenName"),
		Created:    time.Now().Unix(),
		UpData:     time.Now().Unix(),
		Group:      ctx.PostForm("group"),
	}
	err = DataBase.WriteOneDB("user", userData)
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

// SignIn @Title Token
// @Tags Token
// @Summary	用户登陆
// @Produce	json
// @Description 使用账户名和密码进行登录并返回token和refresh_token，其中token每两小过期，需要refresh_token获取新的token，refresh_token每一周过期
// @Param name formData string true "用户名"
// @Param passwd formData string true "密码"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/auth/signin [POST]
func (_ *User) SignIn(ctx *gin.Context) {
	userName := ctx.PostForm("name")
	passwd := ctx.PostForm("passwd")

	if token, ok := utilityFunction.UserPasswdVerify(userName, passwd); ok { // 判断是否符合数据库内数据
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: gin.H{
				"token": token,
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

// UpdateUser @Title 用户
// @Tags 用户
// @Summary	更新用户
// @Produce	json
// @Description 更新指定id用户(需要相关权限token)
// @Param name formData string true "用户名"
// @Param passwd formData string true "密码"
// @Param mail formData string true "邮箱"
// @Param url formData string true "网址"
// @Param screenName formData string true "昵称"
// @Param group formData string true "用户组"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/user [PUT]
func (_ *Update) UpdateUser(ctx *gin.Context) {
	uid, err := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	name := ctx.PostForm("name")
	passwd, err := bcrypt.GenerateFromPassword([]byte(ctx.PostForm("passwd")), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	mail := ctx.PostForm("mail")
	url := ctx.PostForm("url")
	screenName := ctx.PostForm("screenName")
	group := ctx.PostForm("group")
	filter := bson.D{{Key: "_id", Value: uid}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "passwd", Value: string(passwd)},
			{Key: "mail", Value: mail},
			{Key: "url", Value: url},
			{Key: "screenName", Value: screenName},
			{Key: "upData", Value: time.Now().Unix()},
			{Key: "group", Value: group},
		}},
	}

	err = DataBase.UpdateOneDB("user", filter, update)
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

// UpdateTag @Title 标签
// @Tags 标签
// @Summary 更新标签
// @Produce	json
// @Description 更新指定id标签(需要相关权限token)
// @Param name formData string true "名称"
// @Param slug formData string true "缩略名"
// @Param type formData string true "分类"
// @Param description formData string true "描述"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/tag [PUT]
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

// UpdateCategory @Title 分类
// @Tags 分类
// @Summary 更新分类
// @Produce	json
// @Description 更新指定id分类(需要相关权限token)
// @Param name formData string true "名称"
// @Param slug formData string true "缩略名"
// @Param description formData string true "选项描述"
// @Param count formData string true "项目个数"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/category [PUT]
func (_ *Update) UpdateCategory(ctx *gin.Context) {
	tid, _ := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	name := ctx.PostForm("name")
	slug := ctx.PostForm("slug")
	description := ctx.PostForm("description")
	count := ctx.PostForm("count")
	order, _ := strconv.Atoi(ctx.PostForm("order"))

	filter := bson.D{{Key: "_id", Value: tid}}
	update := bson.D{
		bson.E{Key: "$set", Value: bson.D{
			{Key: "name", Value: name},
			{Key: "slug", Value: slug},
			{Key: "type", Value: description},
			{Key: "description", Value: description},
			{Key: "count", Value: count},
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

// UpdateArticle @Title 文章
// @Tags 文章
// @Summary 更新文章
// @Produce	json
// @Description 更新指定id文章(需要相关权限token)
// @Param name formData string true "名称"
// @Param slug formData string true "缩略名"
// @Param type formData string true "分类"
// @Param description formData string true "描述"
// @Param order formData string true "排序"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/article [PUT]
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

	err := DataBase.UpdateOneDB("category", filter, update)
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

// UpdateGroup @Title 用户组
// @Tags 用户组
// @Summary 更新用户组
// @Produce	json
// @Description 更新指定id用户组(需要相关权限token)
// @Param id formData string true "ID"
// @Param name formData string true "名称"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/group [PUT]
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

// DeleteTag @Title 标签
// @Tags 标签
// @Summary	删除标签
// @Produce	json
// @Description 删除指定id标签(需要相关权限token)
// @Param id formData string true "标签id"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/tag [DELETE]
func (_ *Delete) DeleteTag(ctx *gin.Context) { // 删除标签函数
	tid, err := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	filter := bson.D{{Key: "_id", Value: tid}}
	err = DataBase.DeleteOneDB("tag", filter)
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

// DeleteCategory @Title 分类
// @Tags 分类
// @Summary	删除分类
// @Produce	json
// @Description 删除指定id分类(需要相关权限token)
// @Param id formData string true "分类id"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/category [DELETE]
func (_ *Delete) DeleteCategory(ctx *gin.Context) { // 删除分类函数
	cid, err := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	filter := bson.D{{Key: "_id", Value: cid}}
	err = DataBase.DeleteOneDB("category", filter)
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
			"id": cid,
		},
	})
}

// DeleteGroup @Title 用户组
// @Tags 用户组
// @Summary	删除用户组
// @Produce	json
// @Description 删除指定id分类(需要相关权限token)
// @Param id formData string true "用户组id"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/group [DELETE]
func (_ *Delete) DeleteGroup(ctx *gin.Context) { // 删除用户组函数
	gid, err := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	filter := bson.D{{Key: "_id", Value: gid}}
	err = DataBase.DeleteOneDB("group", filter)
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
			"id": gid,
		},
	})
}

// DeleteUser @Title 用户
// @Tags 用户
// @Summary	删除用户
// @Produce	json
// @Description 删除指定id用户(需要相关权限token)
// @Param id formData string true "用户id"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/user [DELETE]
func (_ *Delete) DeleteUser(ctx *gin.Context) { // 删除用户函数
	uid, err := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	filter := bson.D{{Key: "_id", Value: uid}}
	err = DataBase.DeleteOneDB("user", filter)
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

// DeleteArticle @Title 文章
// @Tags 文章
// @Summary	删除文章
// @Produce	json
// @Description 删除指定id文章(需要相关权限token)
// @Param id formData string true "文章id"
// @Param token header string  false "token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/article [DELETE]
func (_ *Delete) DeleteArticle(ctx *gin.Context) { // 删除文章函数
	aid, err := primitive.ObjectIDFromHex(ctx.PostForm("id"))
	if err != nil {
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: ServerError,
			Msg:  "server error",
			Path: ctx.Request.URL.Path,
			Data: nil,
		})
		return
	}
	filter := bson.D{{Key: "_id", Value: aid}}
	err = DataBase.DeleteOneDB("article", filter)
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
			"id": aid,
		},
	})
}

// Refresh @Title Token
// @Tags Token
// @Summary	身份认证获取
// @Description 当refresh_token失效时请求，需要传入过期的refresh_token，如通过校验则返回新的refresh_token
// @Produce	json
// @Param refresh_token formData string true "Refresh_Token"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/api/v1/auth/refresh [POST]
func (_ *User) Refresh(ctx *gin.Context) {
	// 接入传输token
	refreshToken := ctx.PostForm("refresh_token")
	// 判断刷新token是否过期
	if ok, err := utilityFunction.JWTRefreshVerify(refreshToken); !ok { // 验证是否被篡改
		if err != nil {
			ctx.JSON(http.StatusOK, GeneralJSONHeader{
				Code: ServerError,
				Msg:  "server error",
				Path: ctx.Request.URL.Path,
				Data: nil,
			})
		}
	} else {
		// 尝试生成新的refreshToken
		refreshToken, err := utilityFunction.JWTRefreshCreate()
		if err != nil {
			ctx.JSON(http.StatusOK, GeneralJSONHeader{
				Code: ServerError,
				Msg:  "server error",
				Path: ctx.Request.URL.Path,
				Data: nil,
			})
		}
		ctx.JSON(http.StatusOK, GeneralJSONHeader{
			Code: SuccessCode,
			Msg:  "success",
			Path: ctx.Request.URL.Path,
			Data: gin.H{
				"refresh_token": refreshToken,
			},
		})
	}
	// 判断登录token是否过期
	// 刷新token过期重新请求登录token

}

// DeleteArticle @Title 测试
// @Tags 测试
// @Summary	测试函数
// @Produce	json
// @Description 这是一个测试函数，用于在开发模式中测试各种功能(发行版模式下会失效)
// @Param id formData string false "文章id"
// @Param to header string  false "token"
// @Header 200 {string} Token "访问令牌"
// @Success 200 {object} GeneralJSONHeader "OK"
// @Router		/test/123 [GET]
func test(ctx *gin.Context) {
	token, _ := utilityFunction.JWTCreate("123")
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

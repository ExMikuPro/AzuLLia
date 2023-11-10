package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const Version = "1.0.0"              // 系统版本号
const SuccessCode = 0                // 正常
const UserInformationVerifyError = 4 // 用户信息认证错误
const UnKnownFileType = 14           // 未知文件类型
const FileSaveError = 15             // 文件保存错误
const ServerDeCodeError = 400        // 数据解析错误
const NoDocuments = 20               // 无文档记录
const DBWriteError = 21              // 数据库写入错误
const DBUpdateError = 22             // 数据库更新错误
const WorkProgress = 520             // 正在构建的界面

const DBAddress = "mongodb://localhost:27017" // 数据库地址
const DBDataBase = "test"                     // 数据库名称

var FileType = []string{".jpg", ".jpeg", ".png"} // 允许上传文件类型

type GeneralJSONHeader struct { // 全局通用标头
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Path string                 `json:"path"`
	Data map[string]interface{} `json:"data"`
}

type Add struct { // 添加信息函数
}

type AdminPage struct { // 后台管理界面
}

type AdminFunction struct { // 数据管理处理函数
}

type Get struct { // 基础页面函数
}

type DBService struct { // 数据库操作函数
}

type Utility struct { // 基本处理函数
}

type User struct { // 用户相关函数
}

type Update struct {
}

type articleTable struct { // 文章内容表
	ID           primitive.ObjectID `bson:"_id" json:"id"`                    // 主键，自增
	Title        string             `bson:"title" json:"title"`               // 内容标题
	Slug         string             `bson:"slug" json:"slug"`                 // 内容缩略名
	Created      int64              `bson:"created" json:"created"`           // 内容生成时的时间戳
	Modified     int64              `bson:"modified" json:"modified"`         // 内容修改时的时间戳
	Text         string             `bson:"text" json:"text"`                 // 内容文字
	Order        int8               `bson:"order" json:"order"`               // 排序
	AuthorId     string             `bson:"authorId" json:"authorId"`         // 内容所属用户
	Template     string             `bson:"template" json:"template"`         // 内容所使用的模版
	Type         string             `bson:"type" json:"type"`                 // 内容类别
	Status       string             `bson:"status" json:"status"`             // 内容状态
	Password     string             `bson:"password" json:"password"`         // 受保护内容的密码
	AllowComment bool               `bson:"allowComment" json:"allowComment"` // 是否允许评论
}

type tageTable struct { // 标签关系表
	ID          primitive.ObjectID `bson:"_id" json:"id"`                  // 主键，自增
	Name        string             `bson:"name" json:"name"`               // 名称
	Slug        string             `bson:"slug" json:"slug"`               // 项目缩略名
	Type        string             `bson:"type" json:"type"`               // 项目类型
	Description string             `bson:"description" json:"description"` // 选项描述
	Count       int16              `bson:"count" json:"count"`             // 项目个数
	Order       int                `bson:"order" json:"order"`             // 项目排序
}

type typeTable struct { // 分类表
	ID          primitive.ObjectID `bson:"_id" json:"id"`                  // 主键，自增
	Name        string             `bson:"name" json:"name"`               // 名称
	Slug        string             `bson:"slug" json:"slug"`               // 项目缩略名
	Description string             `bson:"description" json:"description"` // 选项描述
	Count       int16              `bson:"count" json:"count"`             // 项目个数
	Order       int                `bson:"order" json:"order"`             // 项目排序
}

type userGroupTable struct { // 用户组表
	ID   primitive.ObjectID `bson:"_id" json:"id"`    // 主键
	Name string             `bson:"name" json:"name"` // 组名
}

type userTable struct { // 用户表
	ID         primitive.ObjectID `bson:"_id" json:"id"`                // 主键，自增
	Name       string             `bson:"name" json:"name"`             // 用户名
	Passwd     string             `bson:"passwd" json:"passwd"`         // 用户密码
	Mail       string             `bson:"mail" json:"mail"`             // 用户邮箱
	Url        string             `bson:"url" json:"url"`               // 用户主页地址
	ScreenName string             `bson:"screenName" json:"screenName"` // 用户昵称
	Created    int64              `bson:"created" json:"created"`       // 用户注册时间
	UpData     int64              `bson:"upData" json:"upData"`         // 用户修改时间
	Group      string             `bson:"group" json:"group"`           // 所属用户组
}

type userLoginData struct { // 网页传回参数
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}

var get = &Get{}                  // 基础页面主结构体
var add = &Add{}                  // 添加函数结构体
var dBService = &DBService{}      // 数据库相关操作函数
var utility = &Utility{}          // 基本操作函数
var user = &User{}                // 用户操作函数
var update = &Update{}            // 更新操作函数
var DataBase = dBService.InitDB() // 初始化数据库

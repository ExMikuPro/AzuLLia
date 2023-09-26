package service

import "github.com/gin-gonic/gin"

const Version = "1.0.0"      // 系统版本号
const SuccessCode = 0        //正常
const ErrorVerify = 1        // 错误的验证密钥
const ErrorSession = 2       // 错误的连接
const ErrorPermissions = 500 // 权限错误
const WorkProgress = 520     // 正在构建的界面

const seed = 15 //  生成随机认证码

const DBAddress = "mongodb://localhost:27017" // 数据库地址
const DBDataBase = "test"                     // 数据库名称

var pageAddr = map[string]string{ // 路由地址
	// 组地址
	"data": "/data/", // 内容获取函数组
	"api":  "/api/",  // 列表获取函数组
	// 页面地址
	// data组
	"paperContext": "/paperContext/:cid", // 文章内容
	"tag":          "/tag/:tid",
	// api组
	"tagList":   "/tagList",   // 标签列表
	"paperList": "/paperList", // 文章列表
	// 系统页面地址
	"about":  "/about",  // 关于框架界面
	"verify": "/verify", // 获取认证界面
}

type GeneralJSONHeader struct { // 全局通用标头
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Path string `json:"path"`
	Data gin.H  `json:"data"`
}

var RequiredHeaders = map[string]string{ // 验证服务
	// 生成随机数
	"Verify": verifyCode(seed),
}

var DataBase = InitDB() // 初始化数据库

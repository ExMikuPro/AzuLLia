package service

import "github.com/gin-gonic/gin"

const Version = "1.0.0"      // 系统版本号
const SuccessCode = 0        //正常
const ErrorVerify = 1        // 错误的验证密钥
const ErrorSession = 2       // 错误的连接
const ErrorPermissions = 500 // 权限错误
const WorkProgress = 520     // 正在构建的界面

const DBAddress = "mongodb://localhost:27017" // 数据库地址
const DBDataBase = "test"                     // 数据库名称

type GeneralJSONHeader struct { // 全局通用标头
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Path string `json:"path"`
	Data gin.H  `json:"data"`
}

var RequiredHeaders = map[string]string{ // 验证服务
	//todo 编写认证码生成器
	"Verify": "dfVndFuvKbJnSdOKe_9517",
}

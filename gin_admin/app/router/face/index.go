package router_face

import (
	"github.com/gin-gonic/gin"
	// V1 "erp_api/gin_admin/app/router/v1"
)

// router 全局上下文
type Ctx = *gin.Context

// // 全局app对象
// var App ginApp.GinApp

// 全局路由对象
type BaseRouter struct {
	AdminBaseApiPath string
	Status           ResponseStatus
	R                *gin.Engine
}

// response status  响应结构体
type ResponseStatus struct {
	// 成功
	OK int
	// 服务器错误
	ServerErr int
	// 请求query错误
	QueryErr int
	// 上传参数错误
	BodyErr int
	// token错误
	TokenErr int
	// 越权操作
	PowerErr int
}

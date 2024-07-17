package router

import (
	"github.com/gin-gonic/gin"
	"erp_api/gin_admin/app/global"
	face "erp_api/gin_admin/app/router/face"
	V1 "erp_api/gin_admin/app/router/v1"
)

type GroupRouter struct {
	base face.BaseRouter
	V1   V1.V1
}

var BaseRouterApp = new(GroupRouter)

var Status face.ResponseStatus = face.ResponseStatus{
	OK:        200,
	ServerErr: 500,
	QueryErr:  400,
	BodyErr:   400,
	TokenErr:  401,
	PowerErr:  403,
}

// 用于向 Gin 引擎添加路由
func Inject(e *gin.Engine) {
	BaseRouterApp.base.R = e
	BaseRouterApp.base.AdminBaseApiPath = global.BASEROUTERPATH
	BaseRouterApp.base.Status = Status
	// 注入V1路由
	BaseRouterApp.V1.V1RouterInject(e,&BaseRouterApp.base)
}

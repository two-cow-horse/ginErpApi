package main

import (
	"erp_api/gin_admin/app"
)

func main() {
	var App  ginApp.GinApp
	// 启动服务
	App.Service()
}

// app.go
package ginApp

import (
	// "erp_api/app/router/index"
	"fmt"
	"github.com/gin-gonic/gin"
	"erp_api/gin_admin/app/global"
	"erp_api/gin_admin/app/initialize"
	middleware "erp_api/gin_admin/app/middleware"
	model "erp_api/gin_admin/app/model"
	router "erp_api/gin_admin/app/router"
)

type GinApp struct {
	// 路由
	Engine *gin.Engine
	// 后台基础路径
}

var App GinApp

func (s *GinApp) Service() {
	/* 创建路由 */
	s.Engine = gin.Default()

	/* 中间件 */
	s.Engine.Use(middleware.MiddleWare())   // 其他中间件
	s.Engine.Use(middleware.JWTAdminAuth()) // 验证jwt

	/* 初始化Viper */
	global.VP = initialize.Viper()

	/* 注入路由 */
	router.Inject(s.Engine)

	/* 注入对应实例 */
	initialize.Inject()

	/* 连接数据库 */
	global.DB = initialize.Gorm()

	/* 生成表结构 */
	model.ModelsApp.AutoMigrate()

	/* 监听端口 */
	err := s.Engine.Run(":8000")
	if err != nil {
		fmt.Println("server start faild, err:", err)
		return
	}
}

package v1_router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1Ctl "erp_api/gin_admin/app/controller/admin/v1"
	face "erp_api/gin_admin/app/router/face"
)

type V1 struct{}

var (
	V1PATH            = "/v1"
	AdminInfo         = "/info"          // 用户信息
	AdminLogin        = "/login"         // 登录
	AdminLogout       = "/logout"        // 登出
	Menu              = "/menu"          // 当前用户路由权限
	Auth              = "/auth"          // 当前按钮权限
	AdminUser         = "/adminUser"     // 用户 Restful
	AdminUserFindByID = "/adminUser/:id" // 用户 ID查询信息
	Role              = "/role"          // 角色 Restful
	RoleFindByID      = "/role/:id"      // 角色 ID查询信息
)

func (s *V1) V1RouterInject(e *gin.Engine, app *face.BaseRouter) {
	// 注入路由
	v1 := e.Group(fmt.Sprintf("%s%s", app.AdminBaseApiPath, V1PATH))
	
	v1.POST(AdminLogin, v1Ctl.GroupApp.User.Login)
	v1.GET(AdminInfo, v1Ctl.GroupApp.User.Info)
	v1.GET(Menu, v1Ctl.GroupApp.Menu.SelfMenuList)
	v1.GET(Auth, v1Ctl.GroupApp.Menu.SelfAuthList)
	v1.POST(AdminUser, v1Ctl.GroupApp.User.Create)
	v1.PUT(AdminUser, v1Ctl.GroupApp.User.Update)
	v1.GET(AdminUser, v1Ctl.GroupApp.User.List)
	v1.GET(AdminUserFindByID, v1Ctl.GroupApp.User.Datelis)
	v1.GET(Role, v1Ctl.GroupApp.Role.List)
	v1.GET(RoleFindByID, v1Ctl.GroupApp.Role.Detail)
	v1.POST(Role, v1Ctl.GroupApp.Role.Create)
	v1.PUT(Role, v1Ctl.GroupApp.Role.Update)
}

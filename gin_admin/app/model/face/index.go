package model_face

import (
	"erp_api/gin_admin/app/global"
	admin "erp_api/gin_admin/app/model/admin"
	"log"
)

type Models struct {
	AdminUser admin.User
	AdminRole admin.Role
	Authority admin.Authority
	UserAuthority admin.RoleAuthority
}

// 同步表结构
func (m *Models) AutoMigrate() {
	InitModels := []interface{}{
		&admin.User{},
		&admin.Role{},
		&admin.RoleAuthority{},
		&admin.Authority{},
	}
	for _, model := range InitModels {

		if err := global.DB.AutoMigrate(model); err != nil {
			log.Println("init table error", err)
		}
	}
	m.AdminUser.InitUser()
	m.AdminRole.InitRole()
	m.Authority.InitAuthority()
}


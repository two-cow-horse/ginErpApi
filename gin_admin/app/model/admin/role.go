package admin_model

import (
	"log"
	"erp_api/gin_admin/app/global"
)

// admin 角色表
type Role struct {
	global.GVA_MODEL
	Name      string `json:"name" gorm:"type:varchar(200);default:'';comment:'角色名称'"`
	Status uint   `json:"status" gorm:"type:int(1);default:0;comment:'是否禁用 0正常1禁用'"`
}

// 表名称
func (Role) TableName() string {
	return "erp_sys_role"
}

// 默认角色
var defaultRole = []Role{
	{
		Name:      "_root_",
		Status: 0,
	},
}

// 初始化
func (s *Role) InitRole() {
	var count int64 = 0
	global.DB.Model(&Role{}).Where("name = ?", "_root_").Count(&count)
	if count == 0 {
		defaultRole[0].ID = 1
		global.DB.Create(&defaultRole)
		log.Println("添加默认角色-----")
	}
}

// 默认将所有的role查询到redis中
func (s *Role) insertKeep() {
	var count int64 = 0
	global.DB.Model(&Role{}).Where("name = ?", "_root_").Count(&count)
	if count == 0 {
		defaultRole[0].ID = 1
		global.DB.Create(&defaultRole)
		log.Println("添加默认角色-----")
	}
}

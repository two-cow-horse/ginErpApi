package admin_model

import (
	"log"
	"erp_api/gin_admin/app/global"
)

// admin 权限表
type Authority struct {
	ID        uint   `gorm:"primarykey" json:"id"` // 主键ID
	Auth_type int8   `json:"auth_type" gorm:"type:int(1);default:1;comment:'权限类型/1-路由/2-按钮api权限'"`
	Name      string `json:"name" gorm:"type:varchar(200);default:'';comment:'按钮/路由权限名称'"`
	Url       string `json:"url" gorm:"type:varchar(200);default:'';comment:'按钮/路由权限地址'"`
	PID       uint   `json:"p_id" gorm:"type:int(10);default:0;comment:'父ID'"`
	Icon      string `json:"icon" gorm:"type:varchar(200);default:'';comment:'icon图标'"`
	ShowState int8   `json:"show_state" gorm:"type:int(1);default:1;comment:'是否显示/1-显示/0-不显示'"`
	Updated   int64  `gorm:"autoUpdateTime:milli" json:"update_at"` // 使用时间戳毫秒数填充更新时间
	Created   int64  `gorm:"autoCreateTime" json:"create_at"`       // 使用时间戳秒数填充创建时间
}

// 表名称
func (Authority) TableName() string {
	return "erp_sys_authority"
}

// 默认权限
var defaultAuthority = []Authority{
	{
		ID:        1,
		Auth_type: 1,
		Name:      "首页",
		Icon:      "icon",
		ShowState: 1,
		Url:       "/",
		PID:       0,
	},
	{
		ID:        2,
		Auth_type: 2,
		Name:      "测试按钮",
		Icon:      "icon",
		ShowState: 1,
		Url:       "auth",
	},
	{
		ID:        11,
		Auth_type: 1,
		Name:      "数据看板",
		Icon:      "icon",
		ShowState: 1,
		Url:       "/showboard",
		PID:       1,
	},
}

// 初始化用户
func (s *Authority) InitAuthority() {
	var count int64 = 0
	global.DB.Model(&Authority{}).Count(&count)
	// 数据库的长度与默认长度不一直时，默认数据优先覆盖
	if count != int64(len(defaultAuthority)) {
		global.DB.Where("1 = 1").Delete(&Authority{}) // 移除所有数据
		global.DB.Create(&defaultAuthority)           // 数据覆盖
		log.Println("添加默认权限-----")
	}
}

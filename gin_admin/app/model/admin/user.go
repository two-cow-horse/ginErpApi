package admin_model

import (
	"log"
	"erp_api/gin_admin/app/global"
)

// admin 用户表
type User struct {
	global.GVA_MODEL
	Username string `json:"username" form:"username" gorm:"type:varchar(200);default:'';comment:'用户名称'"`
	Account  string `json:"account" form:"account" gorm:"type:varchar(200);default:'';comment:'账号'"`
	Password string `json:"password" form:"password"  gorm:"type:varchar(300);default:'';comment:'密码'"`
	RoleID   uint   `json:"role_id" form:"role_id" gorm:"type:int(10);default:0;comment:'权限ID'"`
	Status   int8   `json:"status" form:"status" gorm:"type:int(1);default:0;comment:'禁用状态/1-禁用/0-启用'"`
	Role     Role   `json:"role" form:"role" gorm:"foreignkey:RoleID;references:ID"` // 使用 RoleID 作为外键
}

// 表名称
func (User) TableName() string {
	return "erp_sys_user"
}

// 默认用户
var defaultUser = []User{
	{
		Username: "admin",
		Account:  "admin",
		RoleID:   1,
		Status:   0,
	},
}

// 初始化
func (s *User) InitUser() {
	var count int64 = 0
	global.DB.Model(&User{}).Count(&count)
	if count == 0 {
		p, _ := global.UTLIS.HashPassword("admin123")
		defaultUser[0].Password = p
		global.DB.CreateInBatches(defaultUser,len(defaultUser))
		log.Println("添加默认用户-----")
	}
}

package admin_model
import (
	"erp_api/gin_admin/app/global"
)
// 用户权限
type RoleAuthority struct {
    global.GVA_MODEL
	Idx uint `json:"idx" gorm:"column:idx;comment:'角色ID/role对应的角色ID'"`
	AuthorityID uint `json:"authority_id" gorm:"column:authority_id;comment:'权限ID/authority_id对应权限表ID'"`
	Authority Authority  `gorm:"foreignkey:AuthorityID;references:ID"` // 使用 AuthorityId 作为引用
	Role     Role   `gorm:"foreignkey:Idx;references:ID"` // 使用 Idx 作为引用
}

// 表名称
func (RoleAuthority) TableName() string {
    return "erp_sys_role_authority"
}

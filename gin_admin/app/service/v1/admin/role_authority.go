package v1_service_sys

import (
	"erp_api/gin_admin/app/global"
	admin "erp_api/gin_admin/app/model/admin"
)

type RoleAuthority struct {
	admin.RoleAuthority
}

// FindAll 获取对应用户的权限
func (s *RoleAuthority) FindShlfAll(id int, authType int) []admin.Authority {
	var data []admin.RoleAuthority
	var authority []admin.Authority
	
	tabName := s.TableName()
	jionStr := "left join erp_sys_authority t1 on " + tabName + ".authority_id = t1.id"
	global.DB.Table(tabName).Joins(jionStr).Where(tabName + ".idx = ? and t1.auth_type = ? ", id, authType).First(&data)

	for _, v := range data {
		authority = append(authority, v.Authority)
	}
	return authority
}

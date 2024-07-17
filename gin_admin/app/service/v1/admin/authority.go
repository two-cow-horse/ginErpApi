package v1_service_sys

import (
	"erp_api/gin_admin/app/global"
	admin "erp_api/gin_admin/app/model/admin"
)

type Authority struct {
	admin.Authority
}

// 获取所有权限
func (s *Authority) FindTypeAll(authType int) []admin.Authority {
	var data []admin.Authority
	global.DB.Where("auth_type = ?",authType).Find(&data)
	return data
}
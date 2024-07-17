package v1_admin_sys_ctl

import (
	v1AdminService "erp_api/gin_admin/app/service/v1/admin"
)

type GroupCtl struct {
	User User
	Menu Menu
}

var (
	UserService = new(v1AdminService.User)
	UserAuthorityService = new(v1AdminService.RoleAuthority)
	AuthorityService = new(v1AdminService.Authority)
	RoleService = new(v1AdminService.Role)
)

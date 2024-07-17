package v1_admin_sys_ctl

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
	"erp_api/gin_admin/app/global"
	admin "erp_api/gin_admin/app/model/admin"
	. "erp_api/gin_admin/app/service/jwt"
)

type Menu struct{}

type MenuTree struct {
	admin.Authority
	Children []MenuTree `json:"children"`
}

func toTree(nodes []admin.Authority) []MenuTree {
	var data []MenuTree

	// 构建MenuTree
	for _, node := range nodes {
		c := MenuTree{
			Authority: node,
			Children:  make([]MenuTree, 0),
		}
		data = append(data, c)
	}

	// 构建map
	treeMap := make(map[uint]*MenuTree)
	for i := range data {
		treeMap[data[i].ID] = &data[i]
	}

	// 构建树
	for i := range data {
		if data[i].PID == 0 { // 假设 PID 为 0 表示顶级节点
			continue // 或者你可以直接将顶级节点添加到另一个切片中
		}
		parent, ok := treeMap[data[i].PID]
		if ok {
			parent.Children = append(parent.Children, data[i]) // 这里使用 MenuTree 而不是 Authority
		}
	}

	// 提取顶级节点（如果 PID 为 0 表示顶级节点）
	var topLevelMenus []MenuTree
	for i := range data {
		if data[i].PID == 0 {
			topLevelMenus = append(topLevelMenus, data[i])
		}
	}

	return topLevelMenus
}

// 获取当前用户的权限
func (m *Menu) SelfMenuList(c *gin.Context) {
	j := NewJWT()
	claims, err := j.ParseToken(strings.TrimSpace(c.Request.Header.Get("Authorization")))
	if err != nil {
		log.Println(err)
		global.UTLIS.Response(c, global.BodyErr, "token error", nil)
	}
	// 如果当前用户为超级管理员，获取所有权限
	if claims.RoleID == 1 {
		res := AuthorityService.FindTypeAll(1)
		global.UTLIS.Response(c, global.OK, "all", toTree(res))
		return
	}
	// 非超级管理员，获取当前用户的权限
	res := UserAuthorityService.FindShlfAll(int(claims.RoleID), 1)
	global.UTLIS.Response(c, global.OK, "shlf", toTree(res))
}

// 获取当前用户的按钮权限
func (m *Menu) SelfAuthList(c *gin.Context) {
	j := NewJWT()
	var list []string
	claims, err := j.ParseToken(strings.TrimSpace(c.Request.Header.Get("Authorization")))
	if err != nil {
		log.Println(err)
		global.UTLIS.Response(c, global.BodyErr, "token error", nil)
	}
	// 如果当前用户为超级管理员，获取所有权限
	if claims.RoleID == 1 {
		res := AuthorityService.FindTypeAll(2)
		for i := 0; i < len(res); i++ {
			list = append(list, res[i].Url)
		}
		global.UTLIS.Response(c, global.OK, "all", list)
		return
	}
	// 非超级管理员，获取当前用户的权限
	res := UserAuthorityService.FindShlfAll(int(claims.RoleID), 2)
	for i := 0; i < len(res); i++ {
		list = append(list, res[i].Url)
	}
	global.UTLIS.Response(c, global.OK, "shlf", list)
}

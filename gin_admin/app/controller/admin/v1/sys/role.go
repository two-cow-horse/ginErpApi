package v1_admin_sys_ctl

import (
	"errors"
	"log"
	"strconv"
	"erp_api/gin_admin/app/global"
	v1AdminService "erp_api/gin_admin/app/service/v1/admin"

	"github.com/gin-gonic/gin"
)

type Role struct{}

// 列表
func (s *Role) List(c *gin.Context) {
	r := v1AdminService.Role{}
	data, err := r.List()
	if err != nil {
		log.Println("get role list err", err)
	}
	global.UTLIS.Response(c, global.OK, "OK", data)
}

// 详情
func (s *Role) Detail(c *gin.Context) {
	r := v1AdminService.Role{}
	idStr := c.Param("id")
	// 将字符串转换为uint类型
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		global.UTLIS.Response(c, global.BodyErr, err.Error(), nil)
		return
	}
	r.ID = uint(id)
	if r.ID == 1 {
		global.UTLIS.Response(c, global.ServerErr, "role id 1 not find", nil)
		return
	}
	data, err := r.FindByID()
	if err != nil {
		global.UTLIS.Response(c, global.ServerErr, err.Error(), nil)
		return
	}
	global.UTLIS.Response(c, global.OK, "OK", data)
}

type CreateRole struct {
	AuthIDs []uint `json:"auth_ids"`
	v1AdminService.Role
}

func (s *Role) CreateOrUpdate(c *gin.Context, up bool) error {
	createRole := CreateRole{}

	//  将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&createRole); err != nil {
		return err
	}
	var byteLen []byte = []byte(createRole.Name)
	log.Println("get role ----->", len(byteLen), "----", len(createRole.Name))
	if up {
		var json map[string]interface{}
		// role :=
		err := createRole.Role.Update(createRole.AuthIDs,json)
		if err != nil {
			return err
		}
	} else {
		if len(byteLen) > 200 || len(createRole.Name) < 2 {
			return errors.New("name  min lenth is 2 and max byte lenth is 200")
		}
		err := createRole.Role.Create(createRole.AuthIDs)
		if err != nil {
			return err
		}
	}

	return nil
}

// 创建
func (s *Role) Create(c *gin.Context) {
	if err := s.CreateOrUpdate(c, false); err != nil {
		global.UTLIS.Response(c, global.BodyErr, err.Error(), nil)
	} else {
		global.UTLIS.Response(c, global.OK, "OK", nil)
	}
}

// 更新
func (s *Role) Update(c *gin.Context) {
	if err := s.CreateOrUpdate(c, true); err != nil {
		global.UTLIS.Response(c, global.BodyErr, err.Error(), nil)
	} else {
		global.UTLIS.Response(c, global.OK, "OK", nil)
	}
}

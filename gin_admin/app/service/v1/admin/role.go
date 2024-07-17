package v1_service_sys

import (
	"erp_api/gin_admin/app/global"
	admin "erp_api/gin_admin/app/model/admin"

	"gorm.io/gorm"
)

type Role struct {
	admin.Role
}

// 通过ID查找数据
func (s *Role) FindByID() (Role, error) {
	var role Role
	err := global.DB.Where("id = ?", s.ID).First(&role).Error
	return role, err
}

// 增加role
func (s *Role) Create(authorityIDs []uint) error {
	// 创建事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 创建role
		if err := tx.Model(&Role{}).Create(&s).Error; err != nil {
			return err
		}

		userAuthList := make([]admin.RoleAuthority, 0, len(authorityIDs))
		// 根据ids数 创建 UserAuthority
		for _, v := range authorityIDs {
			userAuthList = append(userAuthList, admin.RoleAuthority{Idx: s.ID, AuthorityID: v})
		}

		// 批量插入
		err := tx.Model(&admin.RoleAuthority{}).CreateInBatches(userAuthList, len(userAuthList)).Error
		if err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

// 更新role
func (s *Role) Update(authorityIDs []uint,data map[string]interface{}) error {
	// 创建事务
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		// 更新role
		if err := tx.Model(&s).Updates(data).Error; err != nil {
			return err
		}
		// 插入前将对应的Role id 记录清除
		err := tx.Model(&admin.RoleAuthority{}).Where("idx = ?", s.ID).Delete(&admin.RoleAuthority{}).Error
		if err != nil {
			return err
		}

		userAuthList := make([]admin.RoleAuthority, 0, len(authorityIDs))
		// 根据ids数 创建 UserAuthority
		for _, v := range authorityIDs {
			userAuthList = append(userAuthList, admin.RoleAuthority{Idx: s.ID, AuthorityID: v})
		}

		// 批量插入
		if err := tx.Model(&admin.RoleAuthority{}).CreateInBatches(userAuthList, len(userAuthList)).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

// 列表查询
func (s *Role) List() ([]Role, error) {
	var roles []Role
	db := global.DB.Model(&Role{})
	db.Where("name != ?", "_root_")
	if s.Name != "" && s.Name != "_root_" {
		db.Where("name = ?", s.Name)
	}
	if s.Status != 3 {
		db.Where("status = ?", s.Status)
	}
	err := db.Find(&roles).Error
	return roles, err
}

package v1_service_sys

import (
	"fmt"
	"github.com/go-ozzo/ozzo-validation/v3"
	"erp_api/gin_admin/app/global"
	admin "erp_api/gin_admin/app/model/admin"
)

type User struct {
	admin.User
}

// 用户登录数据结构体
type UserLogin struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// 用户登录状态返回
type LoginResponse struct {
	Status int  `json:"status"`
	User   User `json:"user,omitempty"`
}

func (s *User) UserList() []User {
	var list []User
	global.DB.First(&list)
	return list
}

// 查找对应ID的记录
func (s *User) FindById() (User,error) {
	var user User
	err := global.DB.First(&user, s.ID).Error
	return user,err
}

// 更新对应ID的记录
func (s *User) Update() error {
	err := global.DB.Model(&s).Updates(s).Error
	return err
}

// 登录
func (s *UserLogin) Login() LoginResponse {
	var user User

	global.DB.Where("account = ?", s.Account).First(&user)

	if user.ID == 0 {
		return LoginResponse{
			Status: 2,
		}
	}
	if !global.UTLIS.ComparePasswords(user.Password, s.Password) {
		return LoginResponse{
			Status: 4,
		}
	}
	if user.Status == 1 {
		return LoginResponse{
			Status: 4,
		}
	}

	return LoginResponse{
		Status: 0,
		User:   user,
	}
}

// 用户信息
func (s *User) UserInfo() (User, error) {
	var user User
	err := global.DB.Preload("Role").First(&user, s.ID).Error
	return user, err
}

// 验证增加用户数据是否通过
func (s User) ValidateAddUser() error {
	return validation.ValidateStruct(&s,
		// 用户名称 5-50
		validation.Field(&s.Username, validation.Required, validation.Length(5, 50)),
		// 密码 5-50
		validation.Field(&s.Password, validation.Required, validation.Length(5, 50)),
		// 账户名 5-50
		validation.Field(&s.Account, validation.Required, validation.Length(5, 50)),
		//  角色ID 1-99999
		validation.Field((&s.RoleID), validation.Required, validation.Min(uint(1))),
	)
}

// 验证增加用户数据是否通过
func (s User) ValidateUpdateUser() error {
	return validation.ValidateStruct(&s,
		// 用户名称 5-50
		validation.Field(&s.Username, validation.Required, validation.Length(5, 50)),
	)
}

// 通过账号查找数据
func (s *User) FindByAccount(account string) User {
	var user User
	global.DB.Where("account = ?", account).First(&user)
	return user
}

// 添加用户
func (s *User) Add() error {
	return global.DB.Create(s).Error
}

type UserList struct {
	List  []User `json:"list"`
	Total int64  `json:"total"`
	Error error  `json:"msg"`
}

type UserLIstQuery struct {
	User
	global.Page
}

// 列表查询
func (where *User) List(page, limit int, filters map[string]string) *UserList {
	var list []User
	var count int64 = 0
	var userList = new(UserList)
	db := global.DB.Model(&User{})
	if where.ID != 0 {
		db = db.Where("id = ?", where.ID)
	}
	if where.Status != -1 {
		db = db.Where("status = ?", where.Status)
	}
	// 动态添加模糊查询条件
	for field, value := range filters {
		if value != "" {
			switch field {
			case "username_like_all":
				db = db.Where("username LIKE ?", "%"+value+"%") // username 全模糊查询
			case "account_like_all":
				db = db.Where("account LIKE ?", "%"+value+"%") // account 全模糊查询
			default:
				db = db.Where(fmt.Sprintf("%s LIKE ?", field), "%"+value+"%") // 其他字段包含模糊查询
			}
		}
	}
	db = db.Offset(page).Limit(limit)
	err := db.Find(&list).Error
	if err != nil {
		userList.Error = err
		return userList
	}
	err = db.Count(&count).Error
	if err != nil {
		userList.Error = err
		return userList
	}
	userList.List = list
	userList.Total = count
	return userList
}

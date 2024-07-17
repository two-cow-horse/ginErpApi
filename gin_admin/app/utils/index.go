package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Utils struct{}

// HashPassword 使用 Bcrypt 算法生成密码哈希值
func (u *Utils) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords 比较输入的密码与哈希值是否匹配
func (u *Utils) ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Response 统一响应结构体
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// 结构体统一响应
func (r *Utils) Response(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(code, Response{code, data, msg})
}

// 查找一个字符数组中是否包含该数组
func (r *Utils) Include(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

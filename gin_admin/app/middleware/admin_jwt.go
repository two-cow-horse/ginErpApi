package middleware

import (
	"net/http"
	"strings"
	"erp_api/gin_admin/app/global"
	. "erp_api/gin_admin/app/global"
	. "erp_api/gin_admin/app/router/v1"
	. "erp_api/gin_admin/app/service/jwt"

	"github.com/gin-gonic/gin"
)

// admin 不需要校验的路由
var routerArr = []string{
	0: BASEROUTERPATH + V1PATH + AdminLogin,
	1: BASEROUTERPATH + V1PATH + AdminLogout,
}

// amdin 鉴权
func JWTAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if global.UTLIS.Include(routerArr, path) {
			c.Next()
		} else {
			token := c.Request.Header.Get("Authorization")
			if token == "" {
				global.UTLIS.Response(c, global.TokenErr, "token error", nil)
				c.Abort()
				return
			}
			j := NewJWT()
			claims, err := j.ParseToken(strings.TrimSpace(token))
			if err != nil || claims == nil {
				global.UTLIS.Response(c, http.StatusUnauthorized, "token error", nil)
				c.Abort()
				return
			}
			c.Set("claims", claims)
		}
		c.Next()
	}
}

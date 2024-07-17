package v1_admin_sys_ctl

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
	"time"
	"erp_api/gin_admin/app/global"
	. "erp_api/gin_admin/app/service/jwt"
	v1AdminService "erp_api/gin_admin/app/service/v1/admin"
)

type User struct{}

// 用户登录
func (s *User) Login(c *gin.Context) {
	// 声明接收的变量
	var json v1AdminService.UserLogin

	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		global.UTLIS.Response(c, global.BodyErr, "Login json error", nil)
		return
	}

	res := json.Login()
	user := res.User

	// 没有该账号
	if res.Status == 2 {
		global.UTLIS.Response(c, global.QueryErr, " not find account ", nil)
		return
	}

	//  密码错误
	if res.Status == 4 {
		global.UTLIS.Response(c, global.QueryErr, " your input password error", nil)
		return
	}

	// 账号已经被禁用
	if res.Status == 3 {
		global.UTLIS.Response(c, global.QueryErr, "your account has been disabled", nil)
		return
	}

	// 生成token
	jwtService := NewJWT()
	claims := CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		RoleID:   user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "gin_admin",
		},
	}
	token, err := jwtService.CreateToken(claims)
	if err != nil {
		global.UTLIS.Response(c, global.ServerErr, "token assign error", nil)
		return
	}
	// val ,err := global.REDIS.Get(context.Background(),"v").Result()
	// log.Println("登录成功",val,)
	// 登录完成
	global.UTLIS.Response(c, global.OK, "", gin.H{"token": token})
}

// 用户信息
func (s *User) Info(c *gin.Context) {
	log.Println("用户信息")
	j := NewJWT()
	claims, err := j.ParseToken(strings.TrimSpace(c.Request.Header.Get("Authorization")))
	if err != nil || claims == nil {
		global.UTLIS.Response(c, global.QueryErr, "token error", gin.H{})
		return
	}
	u := new(v1AdminService.User)
	log.Println("用户信息", claims)
	u.ID = claims.ID
	res, err := u.UserInfo()
	if err != nil {
		global.UTLIS.Response(c, global.QueryErr, "current   id is not find  user", gin.H{})
		return
	}
	res.Password = "password"
	res.Role.Name = ""
	res.Role.ID = 0xff00
	global.UTLIS.Response(c, global.OK, "", res)
}

// 增加用户
func (s *User) Create(c *gin.Context) {
	// 声明接收的变量
	var user v1AdminService.User

	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		global.UTLIS.Response(c, global.BodyErr, "create json error", nil)
		return
	}
	// 该账号是否存在
	if u := user.FindByAccount(user.Account); u.ID != 0 {
		log.Println("账号已存在", u)
		global.UTLIS.Response(c, global.BodyErr, "your account has been used", nil)
		return
	}
	// 用户不能主动添加admin权限
	if user.RoleID == 1 {
		global.UTLIS.Response(c, global.BodyErr, "role has not find", nil)
		return

	}
	role := v1AdminService.Role{}
	role.ID = user.RoleID
	// 该权限是否存在
	if r, err := role.FindByID(); err != nil || r.ID == 1 {
		global.UTLIS.Response(c, global.BodyErr, "role has not find", nil)
		return
	}
	// 验证数据
	if err := user.ValidateAddUser(); err != nil {
		global.UTLIS.Response(c, global.BodyErr, err.Error(), nil)
		return
	}
	pass,err := global.UTLIS.HashPassword(user.Password)
	if err != nil {
	    global.UTLIS.Response(c, global.BodyErr, err.Error(), nil)
		return
	}
	user.Password = pass
	// 添加
	if err := user.Add(); err != nil {
		global.UTLIS.Response(c, global.ServerErr, err.Error(), nil)
	} else {
		global.UTLIS.Response(c, global.BodyErr, "your '"+user.Account+"' account successfully added", nil)
	}
}

// 编辑用户
func (s *User) Update(c *gin.Context) {
	// 声明接收的变量
	var user v1AdminService.User
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&user); err != nil {
		global.UTLIS.Response(c, global.BodyErr, "body error", nil)
		return
	}
	// 用户不能编辑到admin权限
	if user.RoleID == 1 {
		global.UTLIS.Response(c, global.BodyErr, "role has not find ", nil)
		return

	}
	// 该角色是否存在
	role := v1AdminService.Role{}
	role.ID = user.RoleID
	// 该权限是否存在
	if r, err := role.FindByID(); err != nil || r.ID == 1 {
		global.UTLIS.Response(c, global.BodyErr, "role has not find", nil)
		return
	}
	// 该用户是否存在
	_, err := user.UserInfo()
	if err != nil {
		global.UTLIS.Response(c, global.QueryErr, "current   id is not find  user", gin.H{})
		return
	}
	// 验证数据
	if err := user.ValidateUpdateUser(); err != nil {
		global.UTLIS.Response(c, global.BodyErr, err.Error(), nil)
		return
	}
	updateUser := v1AdminService.User{}
	updateUser.RoleID = user.RoleID
	updateUser.Username = user.Username
	updateUser.ID = user.ID
	// 添加
	if err := updateUser.Update(); err != nil {
		global.UTLIS.Response(c, global.ServerErr, err.Error(), nil)
	} else {
		global.UTLIS.Response(c, global.BodyErr, " info change successfully", nil)
	}
}

// 用户列表
func (s *User) List(c *gin.Context) {
	// 声明接收的变量
	var user v1AdminService.User
	var pageData global.Page

	fuliter := make(map[string]string)
	// 将request的query中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindQuery(&user); err != nil {
		global.UTLIS.Response(c, global.BodyErr, "query json error", nil)
		return
	}
	if err := c.ShouldBindQuery(&pageData); err != nil {
		global.UTLIS.Response(c, global.BodyErr, "query json error", nil)
		return
	}
	// Username
	if user.Username != "" {
		fuliter["username_like_all"] = user.Username
	}
	// Account
	if user.Account != "" {
		fuliter["account_like_all"] = user.Account
	}
	// page data
	limit, page := pageData.GetPage()
	list := user.List(page, limit, fuliter)
	if list.Error != nil {
		global.UTLIS.Response(c, global.ServerErr, list.Error.Error(), nil)
		return
	}
	global.UTLIS.Response(c, global.OK, "OK", gin.H{
		"list":  list.List,
		"total": list.Total,
		"page":  pageData.PageNum,
	})
}

// 用户详情
func (s *User) Datelis(c *gin.Context) {
	user := v1AdminService.User{}
	idStr := c.Param("id")
	// 将字符串转换为uint类型
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		global.UTLIS.Response(c, global.BodyErr, "id error", nil)
		return
	}
	user.ID = uint(id)
	data, err := user.FindById()
	if err != nil {
		global.UTLIS.Response(c, global.BodyErr, err.Error(), nil)
		return
	}
	global.UTLIS.Response(c, global.BodyErr, "OK", data)
}

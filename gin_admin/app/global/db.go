package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"sync"
	"erp_api/gin_admin/app/config"
	"erp_api/gin_admin/app/utils"
)

var (
	// 当前orm实例
	DB    *gorm.DB
	REDIS redis.UniversalClient
	lock  sync.RWMutex
	// 全局Viper
	VP             *viper.Viper
	MODEL          *GVA_MODEL
	UTLIS          *utils.Utils
	CONFIG         config.Config
	BASEROUTERPATH = "/admin/api"
	// 统一Http响应码
	OK        = 0   // 成功
	ServerErr = 500 // 服务器错误
	QueryErr  = 400 // 查询错误
	ParamErr  = 406 // 参数错误
	BodyErr   = 400 // 请求体错误
	TokenErr  = 403 // token错误
	PowerErr  = 404 // 权限错误
)

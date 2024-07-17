package initialize

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"erp_api/gin_admin/app/config"
	"erp_api/gin_admin/app/global"
)

// Viper //
// 优先级: 命令行 > 环境变量 > 默认值
// Author [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {
	env := flag.String("env", "dev", "环境配置: dev 或 test")
	flag.Parse()

	var configFileName string
	switch *env {
	case "dev":
		configFileName = config.ConfigDefaultFile
		log.Printf("当前环境为: %s (默认环境)", *env)
	case "test":
		configFileName = config.ConfigTestFile
		log.Printf("当前环境为: %s (测试环境)", *env)
	default:
		log.Fatalf("未知的环境配置: %s", *env)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件错误: %v", err)
	}
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		log.Fatalf("配置文件解析错误: %v", err)
	}
	v.WatchConfig()
	return v
}

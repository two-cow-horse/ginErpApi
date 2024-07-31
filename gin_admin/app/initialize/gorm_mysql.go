package initialize

import (
	"erp_api/gin_admin/app/config"
	"erp_api/gin_admin/app/global"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mitchellh/mapstructure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// GormMysql 初始化Mysql数据库
func GormMysql() *gorm.DB {
	// 获取配置文件中的mysql配置
	var m config.Mysql

	// 从配置提供者中获取 mysql 配置的 map
	mysqlConfigMap := global.VP.Get("mysql").(map[string]interface{})

	// 从配置文件中获取mysql配置
	if err := mapstructure.Decode(mysqlConfigMap, &m); err != nil {
		log.Fatalf("配置解析错误: %v", err)
	}
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		// m.Dbname +
		fmt.Println("mysql connect error", m.Dsn(), err)
		return nil
	} else {
		selectBAse := db.Exec(fmt.Sprintf("USE %s", m.Dbname))
		log.Printf("cannot err: %v", db.Error)
		if selectBAse.Error != nil {
			// 如果数据库不存在，则创建它
			if orm := db.Exec(fmt.Sprintf("CREATE DATABASE %s CHARACTER SET utf8mb4", m.Dbname)); orm != nil {
				log.Println("cannot create database: ",orm, orm.Error)
			}
			// 选择新创建的数据库
			if orm := db.Exec(fmt.Sprintf("USE %s", m.Dbname)); orm.Error != nil {
				log.Println("cannot use database: ",orm, orm.Error)
			}
		} 
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// 启迪时同步表结构
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate()
}

package initialize

import (
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	dbType := "mysql"
	switch dbType {
	case "mysql":
		return GormMysql()
	}
	return nil
}

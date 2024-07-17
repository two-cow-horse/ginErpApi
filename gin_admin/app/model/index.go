package model

import (
	"log"
	modelFace "erp_api/gin_admin/app/model/face"
)

var ModelsApp = new(modelFace.Models)

func init() {
	// 生成表结构
	log.Println("init model")
}

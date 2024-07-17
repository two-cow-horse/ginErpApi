package service

import (
	v1 "erp_api/gin_admin/app/service/v1"
)
type service struct {
	V1 v1.ServiceGroupApp
}
var ServiceApp = new(service)

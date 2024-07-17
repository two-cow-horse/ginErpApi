package v1_controller
import (
    sys "erp_api/gin_admin/app/controller/admin/v1/sys"
)
type GroupCtl struct {
    User sys.User
    Menu sys.Menu
    Role sys.Role
}
var GroupApp = new(GroupCtl)
package system

import "github.com/cnpythongo/goal/pkg/basic"

// 角色菜单
type RoleMenu struct {
	basic.BaseModel
	RoleID int64 `json:"role_id" gorm:"column:role_id;not null;comment:角色ID"`
	MenuID int64 `json:"menu_id" gorm:"column:menu_id;not null;comment:菜单ID"`
}

func (r *RoleMenu) TableName() string {
	return "system_role_menu"
}

func NewRoleMenu() *RoleMenu {
	return &RoleMenu{}
}

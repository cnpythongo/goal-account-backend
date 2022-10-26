package system

import (
	"github.com/cnpythongo/goal/pkg/basic"
)

// 角色
type Role struct {
	basic.BaseModel
	Name   string `json:"name" gorm:"column:name;type:varchar(128);unique;not null;comment:角色名称"`
	Status string `json:"status" gorm:"column:status;type:enum('active', 'freeze', 'delete');comment:状态"`

	RoleUsers []RoleUser `json:"-"`
	RoleMenus []RoleMenu `json:"-"`
}

// 角色用户
type RoleUser struct {
	basic.BaseModel
	RoleID int64 `json:"role_id" gorm:"column:role_id;not null;comment:角色ID"`
	UserID int64 `json:"user_id" gorm:"column:user_id;not null;comment:用户ID"`
}

func (r *Role) TableName() string {
	return "system_role"
}

func (ru *RoleUser) TableName() string {
	return "system_role_user"
}

func NewRole() *Role {
	return &Role{}
}

func NewRoleUser() *RoleUser {
	return &RoleUser{}
}

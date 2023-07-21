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

func (r *Role) TableName() string {
	return "system_role"
}

func NewRole() *Role {
	return &Role{}
}

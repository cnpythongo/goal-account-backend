package system

import (
	"github.com/cnpythongo/goal/pkg/basic"
)

// 角色用户
type RoleUser struct {
	basic.BaseModel
	RoleID int64 `json:"role_id" gorm:"column:role_id;not null;comment:角色ID"`
	UserID int64 `json:"user_id" gorm:"column:user_id;not null;comment:用户ID"`
}

func (ru *RoleUser) TableName() string {
	return "system_role_user"
}

func NewRoleUser() *RoleUser {
	return &RoleUser{}
}

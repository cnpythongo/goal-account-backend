package system

import "github.com/cnpythongo/goal/pkg/basic"

// 页面操作
type MenuAction struct {
	basic.BaseModel
	MenuID int64  `json:"menu_id" gorm:"column:menu_id;not null;comment:菜单ID"`
	Name   string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:操作名称"`
	Code   string `json:"code" gorm:"column:code;type:varchar(256);unique;not null;comment:操作代码"`
}

func (m *MenuAction) TableName() string {
	return "system_menu_action"
}

func NewMenuAction() *MenuAction {
	return &MenuAction{}
}

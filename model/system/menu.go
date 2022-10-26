package system

import "github.com/cnpythongo/goal/pkg/basic"

// 菜单路由
type Menu struct {
	basic.BaseModel
	ParentID int64   `json:"parent_id" gorm:"cloumn:parent_id;default:0;comment:父菜单"`
	Name     string `json:"name" gorm:"column:name;type:varchar(128);unique;not null;comment:名称"`
	URL      string `json:"url" gorm:"column:url;type:varchar(512);default:'';comment:页面URL"`
	Desc     string `json:"desc" gorm:"column:desc;type:varchar(512);defautl:'';comment:描述"`
	Status   string `json:"status" gorm:"column:status;type:enum('active', 'freeze', 'delete');comment:状态"` // freeze-停用

	MenuActions []MenuAction
}

// 页面操作
type MenuAction struct {
	basic.BaseModel
	MenuID int64   `json:"menu_id" gorm:"column:menu_id;not null;comment:菜单ID"`
	Name   string `json:"name" gorm:"column:name;type:varchar(128);not null;comment:操作名称"`
	Code   string `json:"code" gorm:"column:code;type:varchar(256);unique;not null;comment:操作代码"`
}

func (m *Menu) TableName() string {
	return "system_menu"
}

func (m *MenuAction) TableName() string {
	return "system_menu_action"
}

func NewMenu() *Menu {
	return &Menu{}
}

func NewMenuAction() *MenuAction {
	return &MenuAction{}
}

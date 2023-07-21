package system

import "github.com/cnpythongo/goal/pkg/basic"

// 菜单路由
type Menu struct {
	basic.BaseModel
	ParentID int64  `json:"parent_id" gorm:"cloumn:parent_id;default:0;comment:父菜单"`
	Name     string `json:"name" gorm:"column:name;type:varchar(128);unique;not null;comment:名称"`
	URL      string `json:"url" gorm:"column:url;type:varchar(512);default:'';comment:页面URL"`
	Desc     string `json:"desc" gorm:"column:desc;type:varchar(512);defautl:'';comment:描述"`
	Status   string `json:"status" gorm:"column:status;type:enum('active', 'freeze', 'delete');comment:状态"` // freeze-停用

	MenuActions []MenuAction
}

func (m *Menu) TableName() string {
	return "system_menu"
}

func NewMenu() *Menu {
	return &Menu{}
}

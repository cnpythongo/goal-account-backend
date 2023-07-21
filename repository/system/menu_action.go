package system

import (
	"github.com/cnpythongo/goal/model/system"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IMenuActionRepository interface {
	FindMenuActionsByMenuId(menuId int64) ([]*system.MenuAction, error)
	CreateMenuAction(menu *system.Menu) (*system.MenuAction, error)
	UpdateMenuAction(menu *system.Menu) (*system.MenuAction, error)
	DeleteMenuActions(ids []int64) error
	GetMenuActionById(id int64) (*system.MenuAction, error)
	GetMenuActionByName(name string) (*system.MenuAction, error)
}

type MenuActionRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (m *MenuActionRepository) FindMenuActionsByMenuId(menuId int64) ([]*system.MenuAction, error) {
	panic("implement me")
}

func (m *MenuActionRepository) CreateMenuAction(menu *system.Menu) (*system.MenuAction, error) {
	panic("implement me")
}

func (m *MenuActionRepository) UpdateMenuAction(menu *system.Menu) (*system.MenuAction, error) {
	panic("implement me")
}

func (m *MenuActionRepository) DeleteMenuActions(ids []int64) error {
	panic("implement me")
}

func (m *MenuActionRepository) GetMenuActionById(id int64) (*system.MenuAction, error) {
	panic("implement me")
}

func (m *MenuActionRepository) GetMenuActionByName(name string) (*system.MenuAction, error) {
	panic("implement me")
}

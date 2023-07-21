package system

import (
	"github.com/cnpythongo/goal/model/system"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IMenuRepository interface {
	FindMenus(page, size int, conditions interface{}) ([]*system.Menu, int64, error)
	CreateMenu(menu *system.Menu) (*system.Menu, error)
	UpdateMenu(menu *system.Menu) (*system.Menu, error)
	DeleteMenus(ids []int64) error
	GetMenuById(id int64) (*system.Menu, error)
	GetMenuByName(name string) (*system.Menu, error)
}

type MenuRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (m *MenuRepository) FindMenus(page, size int, conditions interface{}) ([]*system.Menu, int64, error) {
	panic("implement me")
}

func (m *MenuRepository) CreateMenu(menu *system.Menu) (*system.Menu, error) {
	panic("implement me")
}

func (m *MenuRepository) UpdateMenu(menu *system.Menu) (*system.Menu, error) {
	panic("implement me")
}

func (m *MenuRepository) DeleteMenus(ids []int64) error {
	panic("implement me")
}

func (m *MenuRepository) GetMenuById(id int64) (*system.Menu, error) {
	panic("implement me")
}

func (m *MenuRepository) GetMenuByName(name string) (*system.Menu, error) {
	panic("implement me")
}

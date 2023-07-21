package system

import (
	"github.com/cnpythongo/goal/model/system"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IRoleMenuRepository interface {
	FindRoleMenus(page, size int, conditions interface{}) ([]*system.RoleMenu, int, error)
	FindRoleMenusByRoleId(roleId int64) ([]*system.RoleMenu, int, error)
	FindRoleMenusByMenuId(menuId int64) ([]*system.RoleMenu, int, error)

	GetRoleMenuById(id int64) (*system.RoleMenu, error)
	DeleteRoleMenus(ids []int64) error
}

type RoleMenuRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (r *RoleMenuRepository) FindRoleMenus(page, size int, conditions interface{}) ([]*system.RoleMenu, int, error) {
	panic("implement me")
}

func (r *RoleMenuRepository) FindRoleMenusByRoleId(roleId int64) ([]*system.RoleMenu, int, error) {
	panic("implement me")
}

func (r *RoleMenuRepository) FindRoleMenusByMenuId(menuId int64) ([]*system.RoleMenu, int, error) {
	panic("implement me")
}

func (r *RoleMenuRepository) GetRoleMenuById(id int64) (*system.RoleMenu, error) {
	panic("implement me")
}

func (r *RoleMenuRepository) DeleteRoleMenus(ids []int64) error {
	panic("implement me")
}

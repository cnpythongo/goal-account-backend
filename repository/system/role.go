package system

import (
	"github.com/cnpythongo/goal/model/system"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IRoleRepository interface {
	FindRoles(page, size int, conditions interface{}) ([]*system.Role, int, error)

	GetRoleById(id int64) (*system.Role, error)
	GetRoleByName(name string) (*system.Role, error)

	CreateRole(role *system.Role) (*system.Role, error)
	UpdateRole(role *system.Role) (*system.Role, error)
	DeleteRoles(ids []int64) error
}

type RoleRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (r *RoleRepository) FindRoles(page, size int, conditions interface{}) ([]*system.Role, int, error) {
	panic("implement me")
}

func (r *RoleRepository) GetRoleById(id int64) (*system.Role, error) {
	panic("implement me")
}

func (r *RoleRepository) GetRoleByName(name string) (*system.Role, error) {
	panic("implement me")
}

func (r *RoleRepository) CreateRole(role *system.Role) (*system.Role, error) {
	panic("implement me")
}

func (r *RoleRepository) UpdateRole(role *system.Role) (*system.Role, error) {
	panic("implement me")
}

func (r *RoleRepository) DeleteRoles(ids []int64) error {
	panic("implement me")
}

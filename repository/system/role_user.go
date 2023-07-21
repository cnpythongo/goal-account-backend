package system

import (
	"github.com/cnpythongo/goal/model/system"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IRoleUserRepository interface {
	FindRoleUsers(page, size int, conditions interface{}) ([]*system.RoleUser, int, error)
	FindRoleUsersByRoleId(page, size int, roleId int64) ([]*system.RoleUser, int, error)
	GetRoleUserById(id int64) (*system.RoleUser, error)
	GetRoleUserByUserId(userId int64) (*system.RoleUser, error)
	DeleteRoleUsers(ids []int64) error
}

type RoleUserRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (r *RoleRepository) FindRoleUsers(page, size int, conditions interface{}) ([]*system.RoleUser, int, error) {
	panic("implement me")
}

func (r *RoleRepository) FindRoleUsersByRoleId(page, size int, roleId int64) ([]*system.RoleUser, int, error) {
	panic("implement me")
}

func (r *RoleRepository) GetRoleUserById(id int64) (*system.RoleUser, error) {
	panic("implement me")
}

func (r *RoleRepository) GetRoleUserByUserId(userId int64) (*system.RoleUser, error) {
	panic("implement me")
}

func (r *RoleRepository) DeleteRoleUsers(ids []int64) error {
	panic("implement me")
}

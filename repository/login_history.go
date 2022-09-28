package repository

import (
	"github.com/cnpythongo/goal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ILoginHistoryRepository interface {
	GetLoginHistoryObject(id int) (*model.LoginHistory, error)
	GetUserLoginHistoryQueryset(userId, page, size int) ([]*model.LoginHistory, error)
	GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*model.LoginHistory, error)
}

type LoginHistoryRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (l *LoginHistoryRepository) GetLoginHistoryObject(id int) (*model.LoginHistory, error) {
	panic("implement me")
}

func (l *LoginHistoryRepository) GetUserLoginHistoryQueryset(userId, page, size int) ([]*model.LoginHistory, error) {
	panic("implement me")
}

func (l *LoginHistoryRepository) GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*model.LoginHistory, error) {
	panic("implement me")
}

package account

import (
	"github.com/cnpythongo/goal/model/account"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ILoginHistoryRepository interface {
	GetLoginHistoryObject(id int) (*account.LoginHistory, error)
	GetUserLoginHistoryQueryset(userId, page, size int) ([]*account.LoginHistory, error)
	GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*account.LoginHistory, error)
}

type LoginHistoryRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (l *LoginHistoryRepository) GetLoginHistoryObject(id int) (*account.LoginHistory, error) {
	panic("implement me")
}

func (l *LoginHistoryRepository) GetUserLoginHistoryQueryset(userId, page, size int) ([]*account.LoginHistory, error) {
	panic("implement me")
}

func (l *LoginHistoryRepository) GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*account.LoginHistory, error) {
	panic("implement me")
}

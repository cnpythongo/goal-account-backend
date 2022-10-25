package account

import (
	account2 "github.com/cnpythongo/goal/model/account"
	"github.com/cnpythongo/goal/repository/account"
)

type ILoginHistoryService interface {
	GetLoginHistoryObject(id int) (*account2.LoginHistory, error)
	GetUserLoginHistoryQueryset(userId, page, size int) ([]*account2.LoginHistory, error)
	GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*account2.LoginHistory, error)
}

type LoginHistoryService struct {
	LoginHistoryRepo account.ILoginHistoryRepository `inject:"LoginHistoryRepo"`
}

func (l *LoginHistoryService) GetLoginHistoryObject(id int) (*account2.LoginHistory, error) {
	panic("implement me")
}

func (l *LoginHistoryService) GetUserLoginHistoryQueryset(userId, page, size int) ([]*account2.LoginHistory, error) {
	panic("implement me")
}

func (l *LoginHistoryService) GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*account2.LoginHistory, error) {
	panic("implement me")
}

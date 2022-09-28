package service

import (
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/repository"
)

type ILoginHistoryService interface {
	GetLoginHistoryObject(id int) (*model.LoginHistory, error)
	GetUserLoginHistoryQueryset(userId, page, size int) ([]*model.LoginHistory, error)
	GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*model.LoginHistory, error)
}

type LoginHistoryService struct {
	LoginHistoryRepo repository.ILoginHistoryRepository `inject:"LoginHistoryRepo"`
}

func (l *LoginHistoryService) GetLoginHistoryObject(id int) (*model.LoginHistory, error) {
	panic("implement me")
}

func (l *LoginHistoryService) GetUserLoginHistoryQueryset(userId, page, size int) ([]*model.LoginHistory, error) {
	panic("implement me")
}

func (l *LoginHistoryService) GetLoginHistoryQueryset(page, size int, condition interface{}) ([]*model.LoginHistory, error) {
	panic("implement me")
}

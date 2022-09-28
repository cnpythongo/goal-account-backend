package service

import (
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/repository"
)

type IUserProfileService interface {
	// 根据用户ID获取用户资料
	GetUserProfileObjectByUserId(userId int) (*model.UserProfile, error)
}

type UserProfileService struct {
	UserProfileRepo repository.IUserProfileRepository `inject:"UserProfileRepo"`
}

func (u *UserProfileService) GetUserProfileObjectByUserId(userId int) (*model.UserProfile, error) {
	panic("implement me")
}

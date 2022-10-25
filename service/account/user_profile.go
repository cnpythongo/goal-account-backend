package account

import (
	account2 "github.com/cnpythongo/goal/model/account"
	"github.com/cnpythongo/goal/repository/account"
)

type IUserProfileService interface {
	// 根据用户ID获取用户资料
	GetUserProfileObjectByUserId(userId int) (*account2.UserProfile, error)
}

type UserProfileService struct {
	UserProfileRepo account.IUserProfileRepository `inject:"UserProfileRepo"`
}

func (u *UserProfileService) GetUserProfileObjectByUserId(userId int) (*account2.UserProfile, error) {
	panic("implement me")
}

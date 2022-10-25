package account

import (
	account2 "github.com/cnpythongo/goal/model/account"
	"github.com/cnpythongo/goal/service/account"
)

type IUserProfileController interface {
	// 根据用户ID获取用户资料
	GetUserProfileObjectByUserId(userId int) (*account2.UserProfile, error)
}

type UserProfileController struct {
	UserProfileSvc account.IUserProfileService `inject:"UserProfileSvc"`
}

func (u *UserProfileController) GetUserProfileObjectByUserId(userId int) (*account2.UserProfile, error) {
	panic("implement me")
}

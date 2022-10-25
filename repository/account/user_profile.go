package account

import (
	"github.com/cnpythongo/goal/model/account"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserProfileRepository interface {
	// 根据用户ID获取用户资料
	GetUserProfileObjectByUserId(userId int) (*account.UserProfile, error)
}

type UserProfileRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (u *UserProfileRepository) GetUserProfileObjectByUserId(userId int) (*account.UserProfile, error) {
	panic("implement me")
}

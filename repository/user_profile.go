package repository

import (
	"github.com/cnpythongo/goal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserProfileRepository interface {
	// 根据用户ID获取用户资料
	GetUserProfileObjectByUserId(userId int) (*model.UserProfile, error)
}

type UserProfileRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (u *UserProfileRepository) GetUserProfileObjectByUserId(userId int) (*model.UserProfile, error) {
	panic("implement me")
}

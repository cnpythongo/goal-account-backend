package repository

import (
	"github.com/cnpythongo/goal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserRepository interface {
	// 创建用户
	CreateUser(user *model.User) (*model.User, error)
	// 根据ID获取用户
	GetUserById(id int) (*model.User, error)
	// 根据UUID获取用户
	GetUserByUuid(uuid string) (*model.User, error)
	// 获取用户查询集
	GetUserQueryset(page, size int, conditions interface{}) ([]*model.User, int, error)
	// 根据条件获取单一用户
	GetUserByCondition(condition interface{}) (*model.User, error)
	// 根据username获取用户
	GetUserByUsername(username string) (*model.User, error)
	// 根据email获取用户
	GetUserByEmail(email string) (*model.User, error)
}

type UserRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (u *UserRepository) GetUserByCondition(condition interface{}) (*model.User, error) {
	result := model.NewUser()
	err := u.DB.Where(condition).First(result).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			u.Logger.Errorf("admin.user.UserRepository.GetUserByCondition Error ==> %v", err)
			u.Logger.Infof("condition ==> %v", condition)
		}
		return nil, err
	}
	return result, nil
}

func (u *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	condition := map[string]interface{}{"username": username}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

func (u *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	condition := map[string]interface{}{"email": email}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

func (u *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	err := u.DB.Create(user).Error
	if err != nil {
		u.Logger.Errorf("admin.user.UserRepository.CreateUser Error ==> %v", err)
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetUserByUuid(uuid string) (*model.User, error) {
	condition := map[string]interface{}{"uuid": uuid}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

func (u *UserRepository) GetUserQueryset(page, size int, conditions interface{}) ([]*model.User, int, error) {
	qs := u.DB.Model(model.NewUser())
	if conditions != nil {
		qs = qs.Where(conditions)
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		u.Logger.Errorf("admin.user.UserRepository.GetUserQueryset Count Error ==> ", err)
		return nil, 0, err
	}
	result := model.NewUsers()
	err = qs.Find(&result).Error
	if err != nil {
		u.Logger.Errorf("admin.user.UserRepository.GetUserQueryset Query Error ==> ", err)
		return nil, 0, err
	}
	return result, int(total), nil
}

func (u *UserRepository) GetUserById(userID int) (*model.User, error) {
	condition := map[string]interface{}{"id": userID}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

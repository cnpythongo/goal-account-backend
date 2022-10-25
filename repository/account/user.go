package account

import (
	"github.com/cnpythongo/goal/model/account"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IUserRepository interface {
	// 创建用户
	CreateUser(user *account.User) (*account.User, error)
	// 根据ID获取用户
	GetUserById(id int) (*account.User, error)
	// 根据UUID获取用户
	GetUserByUuid(uuid string) (*account.User, error)
	// 获取用户查询集
	GetUserQueryset(page, size int, conditions interface{}) ([]*account.User, int, error)
	// 根据条件获取单一用户
	GetUserByCondition(condition interface{}) (*account.User, error)
	// 根据username获取用户
	GetUserByUsername(username string) (*account.User, error)
	// 根据email获取用户
	GetUserByEmail(email string) (*account.User, error)
	// 更新用户信息、状态等
	UpdateUser(user *account.User) (*account.User, error)
	// 批量更新用户信息、状态等
	UpdateUsers(uuids []string, values interface{}) error
}

type UserRepository struct {
	DB     *gorm.DB       `inject:""`
	Logger *logrus.Logger `inject:""`
}

func (u *UserRepository) GetUserByCondition(condition interface{}) (*account.User, error) {
	result := account.NewUser()
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

func (u *UserRepository) GetUserByUsername(username string) (*account.User, error) {
	condition := map[string]interface{}{"username": username}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

func (u *UserRepository) GetUserByEmail(email string) (*account.User, error) {
	condition := map[string]interface{}{"email": email}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

func (u *UserRepository) CreateUser(user *account.User) (*account.User, error) {
	err := u.DB.Create(user).Error
	if err != nil {
		u.Logger.Errorf("admin.user.UserRepository.CreateUser Error ==> %v", err)
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetUserByUuid(uuid string) (*account.User, error) {
	condition := map[string]interface{}{"uuid": uuid}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

func (u *UserRepository) GetUserQueryset(page, size int, conditions interface{}) ([]*account.User, int, error) {
	qs := u.DB.Model(account.NewUser())
	if conditions != nil {
		qs = qs.Where(conditions)
	}
	var total int64
	err := qs.Count(&total).Error
	if err != nil {
		u.Logger.Errorf("admin.user.UserRepository.GetUserQueryset Count Error ==> ", err)
		return nil, 0, err
	}
	if page > 0 && size > 0 {
		offset := (page - 1) * size
		qs = qs.Limit(size).Offset(offset)
	}
	result := account.NewUsers()
	err = qs.Find(&result).Error
	if err != nil {
		u.Logger.Errorf("admin.user.UserRepository.GetUserQueryset Query Error ==> ", err)
		return nil, 0, err
	}
	return result, int(total), nil
}

func (u *UserRepository) GetUserById(userID int) (*account.User, error) {
	condition := map[string]interface{}{"id": userID}
	result, err := u.GetUserByCondition(condition)
	return result, err
}

func (u *UserRepository) UpdateUser(user *account.User) (*account.User, error) {
	err := u.DB.Save(&user).Error
	return user, err
}

func (u *UserRepository) UpdateUsers(uuids []string, values interface{}) error {
	err := u.DB.Model(account.NewUser()).Where("uuid in ?", uuids).Updates(values).Error
	return err
}

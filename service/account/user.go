package account

import (
	"github.com/cnpythongo/goal/model/account"
	accountRepo "github.com/cnpythongo/goal/repository/account"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"time"

	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/router/middleware"
)

// 查询用户结构体
type ReqGetUserListPayload struct {
	Page             int    `json:"page" form:"page" binding:"required"`
	Size             int    `json:"size" form:"size" binding:"required"`
	LastLoginAtStart string `json:"last_login_at_start" form:"last_login_at_start"`
	LastLoginAtEnd   string `json:"last_login_at_end" form:"last_login_at_end"`
}

type ReqAuthLoginPayload struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RespAuthUser struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type ReqUpdateOneUser struct {
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	Gender      int    `json:"gender"`
	Signature   string `json:"signature"`
	Status      string `json:"status"`
	LastLoginAt int64  `json:"-"`
	UpdatedAt   int64  `json:"-"`
	DeletedAt   int64  `json:"-"`
}

type reqBatchUserUuids struct {
	UUIDs []string `json:"uuids" binding:"required"`
}

type ReqUpdateUserAttrs struct {
	reqBatchUserUuids
	Status string `json:"status" binding:"required"`
}

type ReqDeleteUsers reqBatchUserUuids

type IUserService interface {
	// 登录
	Login(payload *ReqAuthLoginPayload) (*RespAuthUser, int, error)
	// 创建用户
	CreateUser(payload *account.User) (*account.User, int, error)
	// 根据ID获取用户
	GetUserById(id int) (*account.User, int, error)
	// 根据UUID获取用户
	GetUserByUuid(uuid string) (*account.User, int, error)
	// 获取用户查询集
	FindUsers(payload *ReqGetUserListPayload, conditions interface{}) (result []*account.User, total, code int, err error)
	// 根据条件获取单一用户
	GetUserByCondition(condition interface{}) (*account.User, error)
	// 根据username获取用户
	GetUserByUsername(username string) (*account.User, int, error)
	// 根据email获取用户
	GetUserByEmail(email string) (*account.User, int, error)
	// 更新用户信息
	UpdateOneUser(uid string, payload *ReqUpdateOneUser) (*account.User, int, error)
	// 批量更新用户属性，如：状态
	UpdateUsers(payload *ReqUpdateUserAttrs) (int, error)
	// 更新用户信息, 支持批量删
	DeleteUsers(uuids []string) (int, error)
}

type UserService struct {
	UserRepo accountRepo.IUserRepository `inject:"UserRepo"`
}

func (u *UserService) Login(payload *ReqAuthLoginPayload) (*RespAuthUser, int, error) {
	user, code, err := u.GetUserByUsername(payload.Username)
	if err != nil {
		return nil, code, err
	}
	if utils.VerifyPassword(payload.Password, user.Password, user.Salt) {
		return nil, response.AccountPasswordError, err
	}
	if user.Status == model.UserStatusFreeze {
		return nil, response.AccountUserFreezeError, nil
	} else if user.Status == model.UserStatusDelete {
		return nil, response.AccountUserNotExistError, nil
	}

	token, err := middleware.GenerateToken(user.Username, user.Password)
	if err != nil {
		return nil, response.JWTTokenGenError, err
	}
	au := &RespAuthUser{
		UUID:     user.UUID,
		Username: user.Username,
		Token:    token,
	}
	go func() {
		_, _, _ = u.UpdateOneUser(user.UUID, &ReqUpdateOneUser{LastLoginAt: time.Now().Unix()})
	}()
	return au, response.SuccessCode, nil
}

func (u *UserService) GetUserByCondition(condition interface{}) (*account.User, error) {
	result, err := u.UserRepo.GetUserByCondition(condition)
	return result, err
}

func (u *UserService) GetUserByUsername(username string) (*account.User, int, error) {
	result, err := u.UserRepo.GetUserByUsername(username)
	if err != nil {
		code := response.AccountQueryUserError
		if err == gorm.ErrRecordNotFound {
			code = response.AccountUserNotExistError
		}
		return nil, code, err
	}
	return result, response.SuccessCode, nil
}

func (u *UserService) GetUserByEmail(email string) (*account.User, int, error) {
	result, err := u.UserRepo.GetUserByEmail(email)
	if err != nil {
		code := response.AccountQueryUserError
		if err == gorm.ErrRecordNotFound {
			code = response.AccountUserNotExistError
		}
		return nil, code, err
	}
	return result, response.SuccessCode, nil
}

func (u *UserService) CreateUser(payload *account.User) (*account.User, int, error) {
	eu, _, _ := u.GetUserByUsername(payload.Username)
	if eu != nil {
		return nil, response.AccountUserExistError, nil
	}
	ue, _, _ := u.GetUserByEmail(payload.Email)
	if ue != nil {
		return nil, response.AccountEmailExistsError, nil
	}
	result, err := u.UserRepo.CreateUser(payload)
	if err != nil {
		return nil, response.AccountCreateError, err
	}
	return result, response.SuccessCode, nil
}

func (u *UserService) GetUserById(id int) (*account.User, int, error) {
	result, err := u.UserRepo.GetUserById(id)
	if err != nil {
		code := response.AccountQueryUserError
		if err == gorm.ErrRecordNotFound {
			code = response.AccountUserNotExistError
		}
		return nil, code, err
	}
	return result, response.SuccessCode, err
}

func (u *UserService) GetUserByUuid(uuid string) (*account.User, int, error) {
	result, err := u.UserRepo.GetUserByUuid(uuid)
	if err != nil {
		code := response.AccountQueryUserError
		if err == gorm.ErrRecordNotFound {
			code = response.AccountUserNotExistError
		}
		return nil, code, err
	}
	return result, response.SuccessCode, nil
}

func (u *UserService) FindUsers(payload *ReqGetUserListPayload, conditions interface{}) (result []*account.User, total, code int, err error) {
	page := payload.Page
	size := payload.Size
	result, total, err = u.UserRepo.FindUsers(page, size, conditions)
	if err != nil {
		return nil, total, response.AccountQueryUserListError, err
	}
	return result, total, response.SuccessCode, nil
}

func (u *UserService) UpdateOneUser(uid string, payload *ReqUpdateOneUser) (*account.User, int, error) {
	user, code, err := u.GetUserByUuid(uid)
	if err != nil {
		return nil, code, err
	}

	err = copier.Copy(&user, &payload)
	if err != nil {
		return nil, response.PayloadCopyError, err
	}

	user.UpdatedAt = time.Now().Unix()
	if user.Status == model.UserStatusDelete && (payload.Status == model.UserStatusActive || payload.Status == model.UserStatusFreeze) {
		user.DeletedAt = 0
	}
	user, err = u.UserRepo.UpdateUser(user)
	if err != nil {
		return nil, response.AccountUserUpdateError, err
	}
	return user, response.SuccessCode, nil
}

func (u *UserService) UpdateUsers(payload *ReqUpdateUserAttrs) (int, error) {
	values := map[string]interface{}{
		"status": payload.Status,
	}
	if payload.Status == model.UserStatusActive || payload.Status == model.UserStatusFreeze {
		values["deleted_at"] = 0
	}
	err := u.UserRepo.UpdateUsers(payload.UUIDs, values)
	if err != nil {
		return response.AccountUserUpdateError, err
	}
	return response.SuccessCode, nil
}

func (u *UserService) DeleteUsers(uuids []string) (int, error) {
	now := time.Now().Unix()
	values := map[string]interface{}{
		"status":     model.UserStatusDelete,
		"updated_at": now,
		"deleted_at": now,
	}
	err := u.UserRepo.UpdateUsers(uuids, values)
	if err != nil {
		return response.AccountUserDeleteError, err
	}
	return response.SuccessCode, nil
}

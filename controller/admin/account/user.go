package account

import (
	account2 "github.com/cnpythongo/goal/model/account"
	"github.com/cnpythongo/goal/service/account"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strconv"

	"github.com/cnpythongo/goal-tools/utils"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/response"
)

type IUserController interface {
	// 登录
	Login(c *gin.Context)
	// 创建用户
	CreateUser(c *gin.Context)
	// 根据ID获取用户
	GetUserById(c *gin.Context)
	// 根据UUID获取用户
	GetUserByUuid(c *gin.Context)
	// 获取用户查询集
	GetUserList(c *gin.Context)
	// 更新用户信息、状态等
	UpdateOneUser(c *gin.Context)
	// 更新用户信息、状态等
	UpdateUsers(c *gin.Context)
	// 删除用户, 逻辑删, 支持批量
	DeleteUsers(c *gin.Context)
}

type UserController struct {
	Logger  *logrus.Logger       `inject:""`
	UserSvc account.IUserService `inject:"UserSvc"`
}

func (u *UserController) Login(c *gin.Context) {
	payload := &account.ReqAuthLoginPayload{}
	err := c.ShouldBindJSON(payload)
	if err != nil {
		response.FailJsonResp(c, response.PayloadError, nil)
		return
	}
	user, code, err := u.UserSvc.Login(payload)
	if code != response.SuccessCode {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, user, nil)
}

func (u *UserController) CreateUser(c *gin.Context) {
	payload := account2.NewUser()
	err := c.ShouldBindJSON(payload)
	if err != nil {
		response.FailJsonResp(c, response.PayloadError, nil)
		return
	}
	user, code, err := u.UserSvc.CreateUser(payload)
	if code != response.SuccessCode {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, user, nil)
}

func (u *UserController) GetUserById(c *gin.Context) {
	pk := c.Param("id")
	id, e := strconv.Atoi(pk)
	if e != nil {
		response.FailJsonResp(c, response.AccountUserIdError, nil)
		return
	}
	result, code, err := u.UserSvc.GetUserById(id)
	if err != nil {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, result, nil)
}

func (u *UserController) GetUserByUuid(c *gin.Context) {
	uid := c.Param("uid")
	result, code, err := u.UserSvc.GetUserByUuid(uid)
	if err != nil {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, result, nil)
}

// 获取用户列表
func (u *UserController) GetUserList(c *gin.Context) {
	payload := &account.ReqGetUserListPayload{}
	err := c.ShouldBindQuery(payload)
	if err != nil {
		response.FailJsonResp(c, response.AccountQueryUserParamError, nil)
		return
	}
	// conditions := map[string]interface{}{}
	users, total, code, err := u.UserSvc.GetUserQueryset(payload, nil)
	if err != nil {
		response.FailJsonResp(c, code, err)
		return
	}
	result := response.Pagination(payload.Page, payload.Size, total)
	result["rows"] = users
	response.SuccessJsonResp(c, result, nil)
}

func (u *UserController) UpdateOneUser(c *gin.Context) {
	uid := c.Param("uid")
	payload := &account.ReqUpdateOneUser{}
	err := c.ShouldBindJSON(payload)
	if err != nil {
		response.FailJsonResp(c, response.PayloadError, err)
		return
	}
	user, code, err := u.UserSvc.UpdateOneUser(uid, payload)
	if err != nil {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, user, nil)
}

func (u *UserController) UpdateUsers(c *gin.Context) {
	payload := &account.ReqUpdateUserAttrs{}
	err := c.ShouldBindJSON(payload)
	if err != nil {
		response.FailJsonResp(c, response.PayloadError, err)
		return
	}
	status := payload.Status
	if status != "" && utils.StrInArrayIndex(status, model.UserStatusAll) == -1 {
		response.FailJsonResp(c, response.PayloadError, nil)
		return
	}

	code, err := u.UserSvc.UpdateUsers(payload)
	if err != nil {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, nil, nil)
}

func (u *UserController) DeleteUsers(c *gin.Context) {
	payload := &account.ReqDeleteUsers{}
	err := c.ShouldBindJSON(payload)
	if err != nil {
		response.FailJsonResp(c, response.PayloadError, err)
		return
	}
	code, err := u.UserSvc.DeleteUsers(payload.UUIDs)
	if err != nil {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, nil, nil)
}

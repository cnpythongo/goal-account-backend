package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/service"
)

type IUserController interface {
	// 根据UUID获取用户
	GetUserByUuid(c *gin.Context)
}

type UserController struct {
	UserSvc service.IUserService `inject:"UserSvc"`
}

func (u *UserController) GetUserByUuid(c *gin.Context) {
	uid := c.Param("uid")
	result, err := u.UserSvc.GetUserByUuid(uid)
	if err != nil {
		code := response.AccountQueryUserError
		if err == gorm.ErrRecordNotFound {
			code = response.AccountUserNotExistError
		}
		response.FailJsonResp(c, code, nil)
		return
	}
	response.SuccessJsonResp(c, result, nil)
}

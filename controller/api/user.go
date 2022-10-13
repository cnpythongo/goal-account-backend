package api

import (
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/cnpythongo/goal/service"
	"github.com/gin-gonic/gin"
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
	result, code, err := u.UserSvc.GetUserByUuid(uid)
	if err != nil {
		response.FailJsonResp(c, code, err)
		return
	}
	response.SuccessJsonResp(c, result, nil)
}

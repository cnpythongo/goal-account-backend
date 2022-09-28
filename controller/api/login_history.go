package api

import (
	"github.com/cnpythongo/goal/service"
	"github.com/gin-gonic/gin"
)

type ILoginHistoryController interface {
	GetUserLoginHistory(c *gin.Context)
}

type LoginHistoryController struct {
	LoginHistorySvc service.ILoginHistoryService `inject:"LoginHistorySvc"`
}

func (l *LoginHistoryController) GetUserLoginHistory(c *gin.Context) {
	panic("implement me")
}

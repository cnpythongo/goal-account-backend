package admin

import (
	"github.com/cnpythongo/goal/service"
	"github.com/gin-gonic/gin"
)

type ILoginHistoryController interface {
	GetLoginHistory(c *gin.Context)
	GetLoginHistoryByUser(c *gin.Context)
	GetLoginHistoryList(c *gin.Context)
}

type LoginHistoryController struct {
	LoginHistorySvc service.ILoginHistoryService `inject:"LoginHistorySvc"`
}

func (l *LoginHistoryController) GetLoginHistory(c *gin.Context) {
	panic("implement me")
}

func (l *LoginHistoryController) GetLoginHistoryByUser(c *gin.Context) {
	panic("implement me")
}

func (l *LoginHistoryController) GetLoginHistoryList(c *gin.Context) {
	panic("implement me")
}

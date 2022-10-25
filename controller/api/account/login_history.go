package account

import (
	"github.com/cnpythongo/goal/service/account"
	"github.com/gin-gonic/gin"
)

type ILoginHistoryController interface {
	GetUserLoginHistory(c *gin.Context)
}

type LoginHistoryController struct {
	LoginHistorySvc account.ILoginHistoryService `inject:"LoginHistorySvc"`
}

func (l *LoginHistoryController) GetUserLoginHistory(c *gin.Context) {
	panic("implement me")
}

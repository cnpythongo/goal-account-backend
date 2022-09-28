package liveness

import (
	"github.com/cnpythongo/goal/pkg/response"
	"github.com/gin-gonic/gin"
)

type ILivenessController interface {
	Ping(c *gin.Context)
}

type LivenessController struct {
}

func (l *LivenessController) Ping(c *gin.Context) {
	response.SuccessJsonResp(c, "hello world", nil)
}

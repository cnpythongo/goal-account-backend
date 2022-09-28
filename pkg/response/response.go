package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonResp(c *gin.Context, code int, extends interface{}) {
	statusCode := http.StatusOK
	data := gin.H{
		"code": code,
		"msg":  GetCodeMsg(code),
	}
	if extends != nil {
		ex := extends.(map[string]interface{})
		for key := range ex {
			data[key] = ex[key]
		}
	}
	c.JSON(statusCode, data)
}

func SuccessJsonResp(c *gin.Context, result interface{}, extends interface{}) {
	if extends != nil {
		ex := extends.(map[string]interface{})
		ex["result"] = result
	} else {
		extends = map[string]interface{}{
			"result": result,
		}
	}
	jsonResp(c, SuccessCode, extends)
}

func FailJsonResp(c *gin.Context, code int, extends interface{}) {
	jsonResp(c, code, extends)
}

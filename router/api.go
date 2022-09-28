package router

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	
	"github.com/cnpythongo/goal/controller/api"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/common/config"
	"github.com/cnpythongo/goal/pkg/common/log"
	"github.com/cnpythongo/goal/pkg/liveness"
)

func InitAPIRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)

	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: model.GetDB()},
		&inject.Object{Value: log.GetLogger()},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}

	userController := api.InjectUserController(injector)
	liveController := liveness.InjectLivenessController(injector)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveController.Ping)
	// api
	userGroup := route.Group("/api/account")
	userGroup.GET("/me", userController.GetUserByUuid)
	userGroup.GET("/users/:uid", userController.GetUserByUuid)
	return route
}

package router

import (
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"

	"github.com/cnpythongo/goal/controller/admin/injectors"
	"github.com/cnpythongo/goal/controller/liveness"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/common/config"
	"github.com/cnpythongo/goal/pkg/common/log"
	"github.com/cnpythongo/goal/router/middleware"
)

func InitAdminRouters(cfg *config.Configuration) *gin.Engine {
	route := initDefaultRouter(cfg)

	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: model.GetDB()},
		&inject.Object{Value: log.GetLogger()},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}

	userController := injectors.InjectUserController(injector)
	liveController := liveness.InjectLivenessController(injector)

	// common test api
	apiGroup := route.Group("/api")
	apiGroup.GET("/ping", liveController.Ping)
	// admin api
	adminGroup := route.Group("/api/account")
	adminGroup.POST("/login", userController.Login)

	adminGroup.Use(middleware.JWTAuth())
	adminGroup.GET("/users", userController.GetUserList)
	adminGroup.POST("/users", userController.CreateUser)
	adminGroup.PATCH("/users", userController.UpdateUsers)
	adminGroup.DELETE("/users", userController.DeleteUsers)

	adminGroup.GET("/users/:uid", userController.GetUserByUuid)
	adminGroup.PATCH("/users/:uid", userController.UpdateOneUser)
	return route
}

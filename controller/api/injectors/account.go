package injectors

import (
	"github.com/facebookgo/inject"

	accountCtrl "github.com/cnpythongo/goal/controller/api/account"
	"github.com/cnpythongo/goal/model"
	"github.com/cnpythongo/goal/pkg/common/log"
	accountRepo "github.com/cnpythongo/goal/repository/account"
	accountSvc "github.com/cnpythongo/goal/service/account"
)

func InjectUserController(injector inject.Graph) accountCtrl.UserController {
	var ctl accountCtrl.UserController
	err := injector.Provide(
		&inject.Object{Value: &accountRepo.UserRepository{}, Name: "UserRepo"},
		&inject.Object{Value: &accountSvc.UserService{}, Name: "UserSvc"},
		&inject.Object{Value: &ctl},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}
	if err := injector.Populate(); err != nil {
		panic("inject fatal: " + err.Error())
	}
	return ctl
}

func InjectUserProfileController(injector inject.Graph) accountCtrl.UserProfileController {
	var ctl accountCtrl.UserProfileController
	err := injector.Provide(
		&inject.Object{Value: &accountRepo.UserProfileRepository{}, Name: "UserProfileRepo"},
		&inject.Object{Value: &accountSvc.UserProfileService{}, Name: "UserProfileSvc"},
		&inject.Object{Value: &ctl},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}
	if err := injector.Populate(); err != nil {
		panic("inject fatal: " + err.Error())
	}
	return ctl
}

func InjectLoginHistoryController(injector inject.Graph) accountCtrl.LoginHistoryController {
	var ctl accountCtrl.LoginHistoryController
	err := injector.Provide(
		&inject.Object{Value: model.GetDB()},
		&inject.Object{Value: log.GetLogger()},
		&inject.Object{Value: &accountRepo.LoginHistoryRepository{}, Name: "LoginHistoryRepo"},
		&inject.Object{Value: &accountSvc.LoginHistoryService{}, Name: "LoginHistorySvc"},
		&inject.Object{Value: &ctl},
	)
	if err != nil {
		panic("inject fatal: " + err.Error())
	}
	if err := injector.Populate(); err != nil {
		panic("inject fatal: " + err.Error())
	}
	return ctl
}

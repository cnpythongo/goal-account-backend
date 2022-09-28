package liveness

import (
	"github.com/facebookgo/inject"
)

func InjectLivenessController(injector inject.Graph) LivenessController {
	var ctl LivenessController
	err := injector.Provide(
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

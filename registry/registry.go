package registry

import (
	controller "academy-go-q32021/interface/controller"
)

type registry struct {
}

type Registry interface {
	NewAppController() controller.App
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.App {
	return controller.App{
		User: r.NewUserController(),
	}
}

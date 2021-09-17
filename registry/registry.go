package registry

import (
	controller "academy-go-q32021/interface/controller"
	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controller.App
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppController() controller.App {
	return controller.App{
		User: r.NewUserController(),
	}
}
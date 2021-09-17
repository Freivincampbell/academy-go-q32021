package registry

import (
	"academy-go-q32021/interface/controller"
	ip "academy-go-q32021/interface/presenter"
	ir "academy-go-q32021/interface/repository"
	"academy-go-q32021/usecase/interactor"
	up "academy-go-q32021/usecase/presenter"
	ur "academy-go-q32021/usecase/repository"
)

func (r *registry) NewUserController() controller.User {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.User {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.User {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.User {
	return ip.NewUserPresenter()
}

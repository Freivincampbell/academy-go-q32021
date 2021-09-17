package interactor

import (
	"academy-go-q32021/domain/model"
	"academy-go-q32021/usecase/presenter"
	"academy-go-q32021/usecase/repository"
)

type user struct {
	UserRepository repository.User
	UserPresenter  presenter.User
}

type User interface {
	Get(u []*model.User) ([]*model.User, error)
}

func NewUserInteractor(r repository.User, p presenter.User) User {
	return &user{r, p}
}

func (us *user) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

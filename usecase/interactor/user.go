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
	ReadUsers() ([]*model.User, error)
	ReadUsersByKey(k string) ([]*model.CustomCSV, error)
	GetUsers(u []*model.User) ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
}

func NewUserInteractor(r repository.User, p presenter.User) User {
	return &user{r, p}
}

func (us *user) ReadUsers() ([]*model.User, error) {
	file, err := us.UserRepository.ReadUsers()
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseReadUsers(file), nil
}

func (us *user) ReadUsersByKey(k string) ([]*model.CustomCSV, error) {
	key, err := us.UserRepository.ReadUsersByKey(k)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseReadUsersByKey(key), nil
}

func (us *user) GetUsers(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.GetUsers(u)
	if err != nil {
		return nil, err
	}

	u, err = us.UserPresenter.ResponseGetUsers(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us *user) GetUserById(id int) (*model.User, error) {
	u, err := us.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseGetUserById(u), nil
}

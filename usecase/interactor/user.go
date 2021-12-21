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
	GetUsersConcurrently(itemType string, items, itemsWorker int) ([]*model.User, error)
}

func NewUserInteractor(r repository.User, p presenter.User) User {
	return &user{r, p}
}

// ReadUsers get users from CSV
func (us *user) ReadUsers() ([]*model.User, error) {
	file, err := us.UserRepository.ReadUsers()
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseReadUsers(file), nil
}

// ReadUsersByKey get users by CSV KEY
func (us *user) ReadUsersByKey(k string) ([]*model.CustomCSV, error) {
	key, err := us.UserRepository.ReadUsersByKey(k)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseReadUsersByKey(key), nil
}

// GetUsers Get users from API
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

// GetUserById Get user by ID
func (us *user) GetUserById(id int) (*model.User, error) {
	u, err := us.UserRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseGetUserById(u), nil
}

// GetUsersConcurrently Get users concurrently
func (us *user) GetUsersConcurrently(itemType string, items, itemsWorker int) ([]*model.User, error) {
	u, err := us.UserRepository.GetUsersConcurrently(itemType, items, itemsWorker)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseGetUsersConcurrently(u), nil
}

package interactor

import (
	"academy-go-q32021/usecase/presenter"
	"academy-go-q32021/usecase/repository"
)

type user struct {
	UserRepository repository.User
	UserPresenter  presenter.User
}

type User interface {
	ReadUsers(f string) (string, error)
}

func NewUserInteractor(r repository.User, p presenter.User) User {
	return &user{r, p}
}

func (us *user) ReadUsers(f string) (string, error) {
	file, err := us.UserRepository.ReadUsers(f)
	if err != nil {
		return "", err
	}

	return us.UserPresenter.ResponseReadUsers(file), nil
}

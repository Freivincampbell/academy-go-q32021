package controller

import (
	"academy-go-q32021/usecase/interactor"
)

type user struct {
	user interactor.User
}

type User interface {
	ReadUsers(c Context) error
}

func NewUserController(us interactor.User) User {
	return &user{us}
}

func (uc *user) ReadUsers(c Context) error {
	var f string

	f, err := uc.user.ReadUsers(f)
	if err != nil {
		return err
	}

	return c.File(f)
}

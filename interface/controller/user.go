package controller

import (
	"academy-go-q32021/usecase/interactor"
	"net/http"
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
	f := "./public/data.csv"

	f, err := uc.user.ReadUsers(f)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.File(f)
}

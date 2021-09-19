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
	ReadUsersByKey(c Context) error
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
func (uc *user) ReadUsersByKey(c Context) error {
	k := c.QueryParam("key")

	f, err := uc.user.ReadUsersByKey(k)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.File(f)
}

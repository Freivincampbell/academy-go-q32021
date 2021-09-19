package controller

import (
	"academy-go-q32021/domain/model"
	"academy-go-q32021/usecase/interactor"
	"net/http"
)

type user struct {
	user interactor.User
}

type User interface {
	ReadUsers(c Context) error
	ReadUsersByKey(c Context) error
	GetUsers(c Context) error
}

func NewUserController(us interactor.User) User {
	return &user{us}
}

func (uc *user) ReadUsers(c Context) error {
	f := "./public/data.csv"

	u, err := uc.user.ReadUsers(f)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *user) ReadUsersByKey(c Context) error {
	k := c.QueryParam("key")

	u, err := uc.user.ReadUsersByKey(k)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}

func (uc *user) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.user.GetUsers(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

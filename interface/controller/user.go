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
	GetUsers(c Context) error
}

func NewUserController(us interactor.User) User {
	return &user{us}
}

func (uc *user) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.user.Get(u)
	if err != nil {
		return err
	}


	return c.JSON(http.StatusOK, u)
}
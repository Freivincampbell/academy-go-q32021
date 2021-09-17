package presenter

import (
	"academy-go-q32021/domain/model"
)

type user struct {
}

type User interface {
	ResponseUsers(us []*model.User) []*model.User
}

func NewUserPresenter() User {
	return &user{}
}

func (up *user) ResponseUsers(us []*model.User) []*model.User {
	//TODO: Transform data

	return us
}
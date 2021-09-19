package presenter

import "academy-go-q32021/domain/model"

type user struct {
}

type User interface {
	ResponseReadUsers(u []*model.User) []*model.User
	ResponseReadUsersByKey(u []*model.CustomCSV) []*model.CustomCSV
}

func NewUserPresenter() User {
	return &user{}
}

func (up *user) ResponseReadUsers(u []*model.User) []*model.User {
	return u
}
func (up *user) ResponseReadUsersByKey(u []*model.CustomCSV) []*model.CustomCSV {
	return u
}

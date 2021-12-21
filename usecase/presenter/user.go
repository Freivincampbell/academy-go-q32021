package presenter

import "academy-go-q32021/domain/model"

type User interface {
	ResponseReadUsers(u []*model.User) []*model.User
	ResponseReadUsersByKey(u []*model.CustomCSV) []*model.CustomCSV
	ResponseGetUsers(u []*model.User) ([]*model.User, error)
	ResponseGetUserById(u *model.User) *model.User
}

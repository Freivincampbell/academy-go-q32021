package repository

import "academy-go-q32021/domain/model"

type User interface {
	ReadUsers() ([]*model.User, error)
	ReadUsersByKey(k string) ([]*model.CustomCSV, error)
	GetUsers(u []*model.User) ([]*model.User, error)
	GetUserById(id int) (*model.User, error)
}

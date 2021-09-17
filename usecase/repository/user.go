package repository

import "academy-go-q32021/domain/model"

type User interface {
	FindAll(u []*model.User) ([]*model.User, error)
}

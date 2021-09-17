package presenter

import "academy-go-q32021/domain/model"

type User interface {
	ResponseUsers(u []*model.User) []*model.User
}

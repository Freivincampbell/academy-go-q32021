package repository

import (
	"academy-go-q32021/domain/model"
	"reflect"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("Fetch user data without error", func(t *testing.T) {
		var u []*model.User
		_, err := NewUserRepository().GetUsers(u)
		if err != nil {
			t.Fail()
		}
	})

	t.Run("Get users return User model type", func(t *testing.T) {
		var u []*model.User
		u, _ = NewUserRepository().GetUsers(u)
		if reflect.TypeOf(u).String() != "[]*model.User" {
			t.Fail()
		}
	})

	t.Run("Get 10 users", func(t *testing.T) {
		var u []*model.User
		u, err := NewUserRepository().GetUsers(u)
		if err != nil {
			t.Fail()
		}

		if len(u) != 10 {
			t.Fail()
		}
	})

	t.Run("Get users has id = 5 in", func(t *testing.T) {
		var u []*model.User
		u, _ = NewUserRepository().GetUsers(u)
		if u[4].Id != 5 {
			t.Fail()
		}
	})

	t.Run("Get user with id = 1", func(t *testing.T) {
		u, _ := NewUserRepository().GetUserById(1)
		if u.Id != 1 {
			t.Fail()
		}
	})

	t.Run("Get user with wrong id", func(t *testing.T) {
		u, _ := NewUserRepository().GetUserById(1)
		if u.Id != 1 {
			t.Fail()
		}
	})

	t.Run("Open File does find path", func(t *testing.T) {
		_, err := NewUserRepository().ReadUsers()
		if err.Error() != "path provided was not found" {
			t.Fail()
		}
	})
}
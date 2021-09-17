package repository

import (
	"academy-go-q32021/domain/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type user struct {
}

type User interface {
	FindAll(u []*model.User) ([]*model.User, error)
}

func NewUserRepository() User {
	return &user{}
}

func (ur *user) FindAll(u []*model.User) ([]*model.User, error) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	err = json.NewDecoder(response.Body).Decode(&u)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return u, nil
}

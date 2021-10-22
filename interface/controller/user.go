package controller

import (
	"academy-go-q32021/domain/model"
	"academy-go-q32021/usecase/interactor"
	"net/http"
	"strconv"
)

type user struct {
	user interactor.User
}

type User interface {
	ReadUsers(c Context) error
	ReadUsersByKey(c Context) error
	GetUsers(c Context) error
	GetUserById(c Context) error
	GetUsersConcurrently(c Context) error
}

func NewUserController(us interactor.User) User {
	return &user{us}
}

// ReadUsers Read all user from CSV
func (uc *user) ReadUsers(c Context) error {
	u, err := uc.user.ReadUsers()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}

// ReadUsersByKey Get data from CSV key
func (uc *user) ReadUsersByKey(c Context) error {
	k := c.QueryParam("key")

	u, err := uc.user.ReadUsersByKey(k)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}

// GetUsers Get all users from an API call
func (uc *user) GetUsers(c Context) error {
	var u []*model.User

	u, err := uc.user.GetUsers(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

// GetUserById Get a user from an API call by ID
func (uc *user) GetUserById(c Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	u, err := uc.user.GetUserById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

// GetUsersConcurrently Get users from a CSV concurrently
func (uc *user) GetUsersConcurrently(c Context) error {
	items, err := strconv.Atoi(c.QueryParam("items"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	itemsWorker, err := strconv.Atoi(c.QueryParam("items_per_workers"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if items < itemsWorker  {
		return c.JSON(http.StatusBadRequest, "Items can not be lower than items for worker")
	}

	itemType := c.QueryParam("type")
	if itemType != "odd" && itemType != "even" {
		return c.JSON(http.StatusBadRequest, "invalid type provided")
	}

	u, err := uc.user.GetUsersConcurrently(itemType, items, itemsWorker)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, u)
}

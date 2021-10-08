package presenter

import (
	"academy-go-q32021/domain/model"
	"encoding/csv"
	"os"
	"strconv"
)

type user struct{}

type User interface {
	ResponseReadUsers(u []*model.User) []*model.User
	ResponseReadUsersByKey(u []*model.CustomCSV) []*model.CustomCSV
	ResponseGetUsers(u []*model.User) ([]*model.User, error)
	ResponseGetUserById(u *model.User) *model.User
}

var CSVFILE = "./public/data.csv"

func NewUserPresenter() User {
	return &user{}
}

func (up *user) ResponseReadUsers(u []*model.User) []*model.User {
	return u
}

func (up *user) ResponseReadUsersByKey(u []*model.CustomCSV) []*model.CustomCSV {
	return u
}

func (up *user) ResponseGetUsers(u []*model.User) ([]*model.User, error) {
	err := storeInCSV(u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (up *user) ResponseGetUserById(u *model.User) *model.User {
	return u
}

func storeInCSV(us []*model.User) error {
	csvFile, err := os.Create(CSVFILE)

	if err != nil {
		return err
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			return
		}
	}(csvFile)

	writer := csv.NewWriter(csvFile)

	var row []string
	row = append(row, "Id", "Name", "Username", "Email", "Phone", "Website")

	err = writer.Write(row)
	if err != nil {
		return err
	}

	row = nil
	for _, u := range us {
		row = append(row, strconv.Itoa(u.Id))
		row = append(row, u.Name)
		row = append(row, u.Username)
		row = append(row, u.Email)
		row = append(row, u.Phone)
		row = append(row, u.Website)
		err := writer.Write(row)
		if err != nil {
			return err
		}

		row = nil
	}

	writer.Flush()

	return nil
}

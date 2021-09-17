package repository

import (
	"academy-go-q32021/domain/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

	storeInCSV(u)

	return u, nil
}

func storeInCSV(us []*model.User) {
	csvFile, err := os.Create("./public/data.csv")

	if err != nil {
		log.Fatalln(err)
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(csvFile)

	writer := csv.NewWriter(csvFile)

	var row []string
	row = append(row, "Id", "Name", "Username", "Email", "Phone", "Website")
	err = writer.Write(row)
	if err != nil {
		log.Fatalln(err)
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
			log.Fatalln(err)
		}
		row = nil
	}

	writer.Flush()
}

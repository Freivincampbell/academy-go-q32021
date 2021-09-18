package repository

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type user struct {
}

type User interface {
	ReadUsers(f string) (string, error)
}

func NewUserRepository() User {
	return &user{}
}

func (ur *user) ReadUsers(f string) (string, error) {
	csvFile, err := openFile(f)
	if err != nil {
		return "", err
	}

	err = validateCSV(csvFile)
	if err != nil {
		return "", err
	}

	return f, nil
}

func openFile(f string) (*os.File, error) {
	csvFile, err := os.Open(f)
	if err != nil {
		err = fmt.Errorf("path provided was not found")
		return nil, err
	}
	return csvFile, nil
}

func validateCSV(csvFile *os.File) error {
	r := csv.NewReader(csvFile)
	for {
		_, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			err = fmt.Errorf("wrong number of fields or wrong format")
			return err
		}
	}

	return nil
}

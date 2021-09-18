package repository

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("Should fail if the file path is wrong", func(t *testing.T) {
		f := "./public/dataa.csv"
		_, err := NewUserRepository().ReadUsers(f)
		comp := fmt.Errorf("path provided was not found")

		if err.Error() != comp.Error() {
			t.Fail()
		}
	})
	t.Run("Should fail if csv File does not have correct format", func(t *testing.T) {
		f := "./public/invalid.csv"
		_, err := NewUserRepository().ReadUsers(f)
		comp := fmt.Errorf("wrong number of fields or wrong format")
		if err.Error() != comp.Error() {
			t.Fail()
		}
	})

	t.Run("Should return csv path if csv File has correct format", func(t *testing.T) {
		filePath := "./public/data.csv"
		f, err := NewUserRepository().ReadUsers(filePath)
		println(err.Error())
		if f != filePath && err != nil {
			t.Fail()
		}
	})
}

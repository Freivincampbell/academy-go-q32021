package repository

var filePath = "./public/data.csv"

type user struct {
}

type User interface {
	ReadUsers(f string) (string, error)
}

func NewUserRepository() User {
	return &user{}
}

func (ur *user) ReadUsers(f string) (string, error) {
	f = filePath
	return f, nil
}

package repository

type User interface {
	ReadUsers(f string) (string, error)
}

package repository

type User interface {
	ReadUsers(f string) (string, error)
	ReadUsersByKey(k string) (string, error)
}

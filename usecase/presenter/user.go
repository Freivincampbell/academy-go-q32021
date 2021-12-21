package presenter

type User interface {
	ResponseReadUsers(f string) string
	ResponseReadUsersByKey(k string) string
}

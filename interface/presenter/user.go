package presenter

type user struct {
}

type User interface {
	ResponseReadUsers(f string) string
	ResponseReadUsersByKey(f string) string
}

func NewUserPresenter() User {
	return &user{}
}

func (up *user) ResponseReadUsers(f string) string {
	return f
}
func (up *user) ResponseReadUsersByKey(k string) string {
	return k
}

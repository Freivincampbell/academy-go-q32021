package controller

type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	File(file string) (err error)
	QueryParam(s string) string
}

package infrastruct

import "net/http"

type CustomError struct {
	msg  string
	Code int
}

func NewError(msg string, code int) *CustomError {
	return &CustomError{
		msg:  msg,
		Code: code,
	}
}

func (c *CustomError) Error() string {
	return c.msg
}

var (
	ErrorInternalServerError = NewError("внутренняя ошибка сервера", http.StatusInternalServerError)
)

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
	ErrorInternalServerError = NewError("Внутренняя ошибка сервера", http.StatusInternalServerError)
	ErrorBadRequest          = NewError("Плохие входные данные", http.StatusBadRequest)
	ErrorUserNotFound        = NewError("Результаты по вашему запросу не найдены", http.StatusOK)
	ErrorDataIsInvalid       = NewError("Некорректная дата, формата ввода даты рождения: \"2020-12-30\"", http.StatusBadRequest)
)

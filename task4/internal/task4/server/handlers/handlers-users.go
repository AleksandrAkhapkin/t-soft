package handlers

import (
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/infrastruct"
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/logger"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

//Получить весь справочник
//(возможно добавление квери параметра minAge задающего минимальный возраст возвращаемых пользователей (включительно))
func (h *Handlers) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var err error

	minAge := r.FormValue("minAge")
	if minAge == "" {
		//получаем всех пользователей, если значение minAge не задано
		users, err := h.srv.GetAllUsers()
		if err != nil {
			apiErrorEncode(w, err)
			return
		}

		apiResponseEncoder(w, users)
		return
	}

	//преобразуем квери параметр в инт
	minAgeInt, err := strconv.Atoi(minAge)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err with parse minAge in handler GetAllUsers"))
		apiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	//получаем пользователей c минимальным возрастом minAge (включительно)
	users, err := h.srv.GetUsersWithMinAge(minAgeInt)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	apiResponseEncoder(w, users)
}

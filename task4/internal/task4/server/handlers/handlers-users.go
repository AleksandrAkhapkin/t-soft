package handlers

import (
	"encoding/json"
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/types"
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/infrastruct"
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/logger"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"time"
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

//Получить пользователя по ID
func (h *Handlers) GetUserByID(w http.ResponseWriter, r *http.Request) {

	//получаем айди пользователя
	userID, err := strconv.Atoi(mux.Vars(r)["userID"])
	if err != nil {
		apiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	//получаем пользователя
	user, err := h.srv.GetUserByID(userID)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	apiResponseEncoder(w, user)
}

//Создать пользователя
func (h *Handlers) MakeUser(w http.ResponseWriter, r *http.Request) {

	//декодирум новые полученные значения
	newUser := types.User{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		apiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	//проверяем что значения не пустые
	if newUser.Birthday == "" || newUser.Name == "" {
		apiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	//проверяем что мы можем распарсить полученную дату
	if _, err := time.Parse(time.RFC3339, newUser.Birthday); err != nil {
		apiErrorEncode(w, infrastruct.ErrorDataIsInvalid)
		return
	}

	//Создаем пользователя
	userID, err := h.srv.MakeUser(&newUser)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}

	newRegister := &types.NewRegister{
		Message: "Пользователь успешно зарегистрирован",
		ID:      userID,
	}

	apiResponseEncoder(w, newRegister)
}

//Изменить пользователя по ID
func (h *Handlers) PutUserByID(w http.ResponseWriter, r *http.Request) {

	//получаем айди пользователя
	userID, err := strconv.Atoi(mux.Vars(r)["userID"])
	if err != nil {
		apiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	//декодирум новые полученные значения
	newUser := types.User{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		apiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}
	newUser.ID = userID

	//проверяем что новые значения не пустые
	if newUser.Birthday == "" || newUser.Name == "" {
		apiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	//проверяем что мы можем распарсить полученную дату
	if _, err = time.Parse(time.RFC3339, newUser.Birthday); err != nil {
		apiErrorEncode(w, infrastruct.ErrorDataIsInvalid)
		return
	}

	//изменяем данные
	err = h.srv.PutUserByID(&newUser)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}
}

//Удалить пользователя по ID
func (h *Handlers) DelUserByID(w http.ResponseWriter, r *http.Request) {

	//получаем айди пользователя
	userID, err := strconv.Atoi(mux.Vars(r)["userID"])
	if err != nil {
		apiErrorEncode(w, infrastruct.ErrorBadRequest)
		return
	}

	//удаляем пользователя
	err = h.srv.DelUserByID(userID)
	if err != nil {
		apiErrorEncode(w, err)
		return
	}
}

package server

import (
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/server/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(h *handlers.Handlers) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.Methods(http.MethodGet).Path("/ping").HandlerFunc(h.Ping)

	//Получить весь справочник
	//(возможно добавление квери параметра minAge задающего минимальный возраст возвращаемых пользователей (включительно))
	router.Methods(http.MethodGet).Path("/users").HandlerFunc(h.GetAllUsers)

	//Получить пользователя по указанному [ID?]
	router.Methods(http.MethodGet).Path("/users/{idStudent:[0-9]+}").HandlerFunc(h.GetUserByID)

	return router

}

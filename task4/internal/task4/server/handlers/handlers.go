package handlers

import (
	"encoding/json"
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/service"
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/infrastruct"
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/logger"
	"net/http"
)

type Handlers struct {
	srv *service.Service
}

func NewHandlers(srv *service.Service) *Handlers {
	return &Handlers{
		srv: srv,
	}
}

func (h *Handlers) Ping(w http.ResponseWriter, _ *http.Request) {

	_, _ = w.Write([]byte("pong"))
}

func apiErrorEncode(w http.ResponseWriter, err error) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if customError, ok := err.(*infrastruct.CustomError); ok {
		w.WriteHeader(customError.Code)
	}

	result := struct {
		Err string `json:"error"`
	}{
		Err: err.Error(),
	}

	if err = json.NewEncoder(w).Encode(result); err != nil {
		logger.LogError(err)
	}
}

func apiResponseEncoder(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		logger.LogError(err)
	}
}

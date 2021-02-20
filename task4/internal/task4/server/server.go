package server

import (
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/server/handlers"
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/logger"
	"github.com/pkg/errors"
	"net/http"
)

func StartServer(handlers *handlers.Handlers, port string) {

	router := NewRouter(handlers)
	logger.LogInfo("Start service in port " + port)
	if err := http.ListenAndServe(port, router); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewRouter"))
	}
}

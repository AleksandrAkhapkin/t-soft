package main

import (
	"flag"
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/client/postgres"
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/server"
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/server/handlers"
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/service"
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/types/config"
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/logger"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//флаг - путь до конфига
	configPath := new(string)
	flag.StringVar(configPath, "config-path", "config/config.yaml", "path to yaml config")
	flag.Parse()

	cnfFile, err := os.Open(*configPath)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with os.Open"))
	}

	cnf := config.Config{}
	if err := yaml.NewDecoder(cnfFile).Decode(&cnf); err != nil {
		logger.LogFatal(errors.Wrap(err, "err with Decode config"))
	}

	pq, err := postgres.NewPostgres(cnf.PostgresDsn)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewPostgres"))
	}

	srv := service.NewService(pq)

	handls := handlers.NewHandlers(srv)

	//Корректное закрытие базы при получении сигнала
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		_ = <-sigChan
		logger.LogInfo("Finish service")
		if err := pq.Close(); err != nil {
			logger.LogFatal(err)
		}
		os.Exit(0)
	}()

	server.StartServer(handls, cnf.ServerPort)
}

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
)

func main() {

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

	srv, err := service.NewService(pq)
	if err != nil {
		logger.LogFatal(errors.Wrap(err, "err with NewService"))
	}

	handls := handlers.NewHandlers(srv)

	server.StartServer(handls, cnf.ServerPort)
}

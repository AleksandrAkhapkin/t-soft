/*Задание 4.
Реализовать REST-API для работы с таблицей "tb_user" базы данных PostgreSQL.
Вывод осуществить по протоколу http в формате json-документа.

Подключение к базе данных: 79.104.55.66:7002
База данных: "hr-test"
Структура таблицы tb_user:
id serial8 not null,
name varchar(255) null,
birthday date null,
age int null,
is_male bool null

Функции API:
1. Получить весь справочник
метод: GET
путь: "/users"

2. Получить пользователей с возрастом более метод: GET
путь: "/users?minAge="

3. Получить пользователя по указанному [ID?] метод: GET
путь: "/users/"

4. Создать пользователя
метод: POST
путь: "/users"
тело запроса:
{
"name": string,
"birthday": date,
"isMale": bool
}
Значение поля "age" должно быть рассчитано, как полное значение лет от даты рождения до текущей даты.

5. Изменить пользователя по указанному метод: PUT
путь: "/users/"
тело запроса:
{
"name": string,
"birthday": date,
"isMale": bool
}
Значение поля "age" должно быть рассчитано как полное значение лет от даты рождения до текущей даты.

6. Удалить пользователя по указанному метод: DELETE
путь: "/users/"*/
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

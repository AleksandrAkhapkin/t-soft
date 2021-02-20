package service

import (
	"database/sql"
	"github.com/AleksandrAkhapkin/testTNS/task4/internal/task4/types"
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/infrastruct"
	"github.com/AleksandrAkhapkin/testTNS/task4/pkg/logger"
	"github.com/pkg/errors"
	"time"
)

//Получить весь справочник
func (s *Service) GetAllUsers() ([]types.User, error) {

	//получаем всех пользователей
	users, err := s.p.GetAllUsers()
	if err != nil {
		if err != sql.ErrNoRows {
			logger.LogError(errors.Wrap(err, "err in service GetAllUsers "))
			return nil, infrastruct.ErrorInternalServerError
		}
		return nil, infrastruct.ErrorUserNotFound
	}

	//задаем возраст пользователей
	for i, _ := range users {
		users[i].Age, err = userAgeCalculator(users[i].Birthday)
		if err != nil {
			logger.LogError(errors.Wrap(err, "err in service GetAllUsers "))
			return nil, infrastruct.ErrorInternalServerError
		}
	}

	return users, nil
}

//Получить пользователей с возрастом более minAge (включительно)
func (s *Service) GetUsersWithMinAge(minAge int) ([]types.User, error) {

	//считаем максимальную дату рождения которая подходит для заданного возраста
	now := time.Now()
	maxBDay := now.AddDate(-1*minAge, 0, 0).Format("2006-01-02")

	//получаем пользователей с датой рождения которая подходит для заданного возраста
	users, err := s.p.GetUserWithMaxBDay(maxBDay)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.LogError(errors.Wrap(err, "err in service GetUserWithMinBDay "))
			return nil, infrastruct.ErrorInternalServerError
		}
		return nil, infrastruct.ErrorUserNotFound
	}

	//задаем возраст пользователей
	for i, _ := range users {
		users[i].Age, err = userAgeCalculator(users[i].Birthday)
		if err != nil {
			logger.LogError(errors.Wrap(err, "err in service GetUserWithMinBDay "))
			return nil, infrastruct.ErrorInternalServerError
		}
	}

	return users, nil
}

//Получить пользователя по айди
func (s *Service) GetUserByID(userID int) (*types.User, error) {

	//получаем пользователя по айди
	user, err := s.p.GetUserByID(userID)
	if err != nil {
		if err != sql.ErrNoRows {
			logger.LogError(errors.Wrap(err, "err in service GetUserByID "))
			return nil, infrastruct.ErrorInternalServerError
		}
		return nil, infrastruct.ErrorUserNotFound
	}

	//задаем возраст
	user.Age, err = userAgeCalculator(user.Birthday)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in service GetUserByID "))
		return nil, infrastruct.ErrorInternalServerError
	}

	return user, nil
}

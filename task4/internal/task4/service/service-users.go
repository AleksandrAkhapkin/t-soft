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

//Создать нового пользователя, возвращает его ID
func (s *Service) MakeUser(newUser *types.User) (int64, error) {

	var err error

	//задаем возраст
	newUser.Age, err = userAgeCalculator(newUser.Birthday)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in service MakeUser "))
		return 0, infrastruct.ErrorBadRequest
	}

	//Создаем нового пользователя
	userID, err := s.p.MakeUser(newUser)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in service MakeUser "))
		return 0, infrastruct.ErrorInternalServerError
	}

	return userID, nil
}

//Изменить пользователя по айди
func (s *Service) PutUserByID(newUser *types.User) error {

	var err error

	//задаем возраст
	newUser.Age, err = userAgeCalculator(newUser.Birthday)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in service PutUserByID "))
		return infrastruct.ErrorBadRequest
	}

	//изменяем пользователя по айди
	if err = s.p.PutUserByID(newUser); err != nil {
		logger.LogError(errors.Wrap(err, "err in service PutUserByID "))
		return infrastruct.ErrorInternalServerError
	}

	return nil
}

//Удалить пользователя по айди
func (s *Service) DelUserByID(userID int) error {

	//удаляем пользователя по айди
	if err := s.p.DelUserByID(userID); err != nil {
		logger.LogError(errors.Wrap(err, "err in service DelUserByID "))
		return infrastruct.ErrorInternalServerError
	}

	return nil
}

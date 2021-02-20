package service

import (
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
		err = errors.Wrap(err, "err in GetAllUsers ")
		logger.LogError(err)
		return nil, infrastruct.ErrorInternalServerError
	}

	//задаем возраст пользователей
	for i, _ := range users {
		users[i].Age, err = userAgeCalculator(users[i].Birthday)
		if err != nil {
			return nil, errors.Wrap(err, "err in GetAllUsers ")
		}
	}

	return users, nil
}

//Принимает дату рождения, возвращает текущий возраст
func userAgeCalculator(bday string) (int, error) {

	now := time.Now()

	birthday, err := time.Parse(time.RFC3339, bday)
	if err != nil {
		return 0, errors.Wrap(err, "in userAgeCalculator with time.Parse")
	}

	years := now.Year() - birthday.Year()
	if now.YearDay() < birthday.YearDay() {
		years--
	}

	return years, nil
}

//Получить пользователей с возрастом более minAge (включительно)
func (s *Service) GetUsersWithMinAge(minAge int) ([]types.User, error) {

	//считаем максимальную дату рождения которая подходит для заданного возраста
	now := time.Now()
	maxBDay := now.AddDate(-1*minAge, 0, 0).Format("2006-01-02")

	//получаем пользователей с датой рождения которая подходит для заданного возраста
	users, err := s.p.GetUserWithMaxBDay(maxBDay)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in GetUserWithMinBDay"))
		return nil, infrastruct.ErrorInternalServerError
	}

	//задаем возраст пользователей
	for i, _ := range users {
		users[i].Age, err = userAgeCalculator(users[i].Birthday)
		if err != nil {
			return nil, errors.Wrap(err, "err in GetUsersWithMinAge ")
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

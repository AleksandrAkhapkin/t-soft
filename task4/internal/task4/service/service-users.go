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

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

	users, err := s.p.GetAllUsers()
	if err != nil {
		err = errors.Wrap(err, "err in GetAllUsers ")
		logger.LogError(err)
		return nil, infrastruct.ErrorInternalServerError
	}

	now := time.Now()
	for i, _ := range users {
		birthday, err := time.Parse(time.RFC3339, users[i].Birthday)
		if err != nil {
			err = errors.Wrap(err, "err in GetAllUsers with time.Parse")
			logger.LogError(err)
			return nil, infrastruct.ErrorInternalServerError
		}

		years := now.Year() - birthday.Year()
		if now.YearDay() < birthday.YearDay() {
			years--
		}

		users[i].Age = years
	}

	return users, nil
}

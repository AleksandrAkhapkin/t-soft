package service

import (
	"github.com/pkg/errors"
	"time"
)

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

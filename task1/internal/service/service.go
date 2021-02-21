package service

import (
	"encoding/json"
	"github.com/AleksandrAkhapkin/testTNS/task1/internal/types"
	"github.com/AleksandrAkhapkin/testTNS/task1/pkg/logger"
	"github.com/pkg/errors"
	"math/rand"
	"net/http"
	"time"
)

//Функция создаёт и вращает срез из 10 случайно сгенерированных элементов human
func GenerateTenRandomHumans() ([]types.Human, error) {

	var err error
	humans := make([]types.Human, 10)
	for i, _ := range humans {
		rand.Seed(time.Now().UnixNano())
		humans[i].Age = rand.Intn(100)
		humans[i].Name, err = randomName()
		if err != nil {
			return nil, errors.Wrap(err, "in randomName ")
		}
	}

	return humans, nil
}

//Функция фильтрует полученный срез по указанному значению (включительно), возвращает результирующий новый срез
func SortHumansByAge(minAge uint, humans []types.Human) []types.Human {

	sortHumans := make([]types.Human, 0, len(humans))
	for _, v := range humans {
		if v.Age >= int(minAge) {
			sortHumans = append(sortHumans, v)
		}
	}

	return sortHumans
}

//Функция генерирует случайные имена
func randomName() (string, error) {

	//Обращаемся к API и получаем случайные имена :)
	res, err := http.Get("https://api.randomdatatools.ru")
	if err != nil {
		return "", errors.Wrap(err, "with http.Get")
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			logger.LogError(errors.Wrap(err, "err in  randomName with res.Body.Close"))
			return
		}
	}()

	name := struct {
		FirstName string `json:"FirstName"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&name); err != nil {
		return "", errors.Wrap(err, "with NewDecoder")
	}

	return name.FirstName, nil
}

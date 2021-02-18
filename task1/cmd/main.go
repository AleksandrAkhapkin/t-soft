/*1. Работа со срезами.
1.1 Создать структуру human, состоящую из двух полей - имени и возраста.
1.2 Создать функцию, возвращающую срез из 10 случайно сгенерированных элементов human.
1.3 Создать функцию, принимающую 2 параметра - срез human и числовое значение минимального возраста. Функция фильтрует полученный срез по указанному значению (включительно), возвращает результирующий новый срез.
1.4 Вывести результат работы двух функций в консоль.*/
package main

import (
	"fmt"
	"github.com/AleksandrAkhapkin/testTNS/task1/internal/service"
	"github.com/AleksandrAkhapkin/testTNS/task1/pkg/logger"
	"github.com/pkg/errors"
)

func main() {

	humans, err := service.GenerateTenRandomHumans()
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in GenerateTenRandomHumans"))
		return
	}

	fmt.Println("Сгенерированные люди:")
	for i, v := range humans {
		fmt.Printf("№%d:\tВозраст: %d,\tИмя: %s\n", i+1, v.Age, v.Name)
	}

	var minAge uint
	fmt.Print("\nВведите минимальный возраст для сортировки (включительно): ")
	_, err = fmt.Scanf("%d", &minAge)
	if err != nil {
		logger.LogError(errors.Wrap(err, "err in main with Scanf"))
		return
	}

	humans = service.SortHumansByAge(minAge, humans)

	fmt.Printf("Люди c возрастом %d (включительно) и старше:", minAge)
	if len(humans) == 0 {
		fmt.Println(" не найдены.")
		return
	}
	for i, v := range humans {
		fmt.Printf("\n№%d:\tВозраст: %d,\tИмя: %s", i+1, v.Age, v.Name)
	}

	fmt.Println()
}

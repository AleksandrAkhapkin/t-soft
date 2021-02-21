/*Задание 3
Указывается число int64
Два параллельных потока последовательно перебирают все числа от 1 до указанного (включительно).
У каждого числа считается количество установленных бит и результат суммируется с общим результатом.
Когда все числа будут обработаны, вывести общий результат, т.е. количество установленных бит у чисел от 1 до указанного.
Каждый поток не должен обрабатывать число, уже обрабатывающееся или обработанное другим потоком.*/
package main

import (
	"fmt"
	"math/bits"
	"sync"
	"time"
)

func main() {

	//данная реализация, как мне показалось, не совсем соответствуют заданию,
	//но при работе с большими числами (например 10000000000), показывает гораздо более высокую скорость выполнения

	var a int64
	var sum int64
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	fmt.Print("Введите число: ")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println(err)
		return
	}

	ts := time.Now()

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(mu *sync.Mutex, wg *sync.WaitGroup, i int) {
			sumRoutine := int64(0)
			positionInRoutine := int64(0 + i)
			defer wg.Done()
			for positionInRoutine <= a {
				//fmt.Printf("routine %d use num %d\n", i+1, position)
				sumRoutine += int64(bits.OnesCount64(uint64(positionInRoutine)))
				positionInRoutine += 2
			}
			mu.Lock()
			sum += sumRoutine
			mu.Unlock()
		}(mu, wg, i)
	}

	wg.Wait()

	fmt.Println("Время выполнения: ", time.Since(ts))

	fmt.Println("Сумма установленных байтов: ", sum)
}

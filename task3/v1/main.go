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

	var a int64
	var sum int64
	var position int64
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
			defer wg.Done()
			mu.Lock()
			for position < a {
				position++
				//fmt.Printf("routine %d use num %d\n", i+1, position)
				sum += int64(bits.OnesCount64(uint64(position)))
			}
			mu.Unlock()
		}(mu, wg, i)
	}

	wg.Wait()

	fmt.Println("Время выполнения: ", time.Since(ts))

	fmt.Println("Сумма установленных байтов: ", sum)
}

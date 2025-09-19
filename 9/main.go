package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup // использую вейтгруппу для ожидания всего вывода горутины
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch2) // закрываем второй канал, после того, как отправили все даныее из канала 1 и он закрылся

		for val := range ch1 {
			ch2 <- val * 2 // пишу в канал 2 квадраты полученных числе из канала 1
		}
	}()

	wg.Add(1)
	go func() {
		for val := range ch2 {
			fmt.Println(val) // выводим в stdout полученное вычисленное значение
		}

		wg.Done()
	}()

	var arr []int

	for i := 0; i < 100; i++ { // наполняем массив
		arr = append(arr, i)
	}

	for _, val := range arr { // отправляем числа из массива в первый канал
		ch1 <- val
	}

	close(ch1) // закрываем канал 1 после передачи всех значений
	wg.Wait()
}

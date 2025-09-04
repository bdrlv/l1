package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ошибка некорректного количества аргументов. Укажите количество воркеров в качестве аргумента.")
		os.Exit(1)
	}

	wn, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("некорректный формат аргумента, %v. Значение должно быть положительным целочисенным и больше, либо равное 1.", err)
		os.Exit(1)
	}

	if wn <= 0 {
		fmt.Println("значение колечства воркеров должно быть >= 1")
		os.Exit(1)
	}

	fmt.Printf("L1.3 - Запущено воркеров: %d\n", wn)

	ch := make(chan int)

	for i := 0; i < wn; i++ {
		go func(worker_id int, data <-chan int) {
			for val := range data {
				fmt.Printf("Воркер: %d, Полученное значение: %d\n", worker_id, val)
			}
		}(i, ch)
	}

	for {
		ch <- time.Now().Nanosecond()
		time.Sleep(100 * time.Millisecond)
	}
}

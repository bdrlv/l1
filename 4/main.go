package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("ошибка некорректного количества аргументов. Укажите количество воркеров в качестве аргумента.")
		os.Exit(1)
	}

	wn, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Printf("некорректный формат аргумента, %v. Значение должно быть положительным целочисенным и больше, либо равное 1.", err)
		os.Exit(1)
	}

	if wn <= 0 {
		log.Println("значение колечства воркеров должно быть >= 1")
		os.Exit(1)
	}

	log.Printf("L1.4 - Запущено воркеров: %d\n", wn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan int)

	for i := 0; i < wn; i++ {
		go func(ctx context.Context, worker_id int, data <-chan int) {
			for {
				select {
				case val, ok := <-data:
					if !ok {
						log.Println("Канал закрыт")
						return
					}
					log.Printf("Воркер: %d, Полученное значение: %d\n", worker_id, val)
				case <-ctx.Done():
					log.Println("Команда завершения работы")
					return
				}

			}
		}(ctx, i, ch)
	}

	for {
		select {
		case <-ctx.Done():
			return // если получаем отмену контектста, то выходим
		default: // пока нет отмены → продолжаем слать данные
			ch <- time.Now().Nanosecond()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

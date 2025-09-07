package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, syscall.SIGINT) // ловим системный вызов завершения
	go func() {                          // запускаем горутину, которая будет ожидать получения сискола
		<-ctrlc // ожидаем получения, не блокируя выполнения
		cancel()
	}()

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
					log.Printf("Команда завершения работы воркера %d", worker_id)
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

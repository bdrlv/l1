package main

import (
	"context"
	"log"
	"time"
)

func main() {
	log.Println("L1.5 Запущено")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	deadline := time.After(3 * time.Second) // создаем длительность, после которой наступает дедлайн для завершения приложения

	ch := make(chan int)

	go func(ctx context.Context, data <-chan int) {
		for {
			select {
			case val, ok := <-data:
				if !ok {
					log.Println("Канал закрыт")
					return
				}
				log.Printf("Полученное значение: %d\n", val)
			case <-deadline:
				log.Println("Время работы программы истекло")
				cancel()
			case <-ctx.Done():
				return
			}

		}
	}(ctx, ch)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- time.Now().Nanosecond()
			time.Sleep(100 * time.Millisecond)
		}
	}
}

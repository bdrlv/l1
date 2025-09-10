package main

import (
	"context"
	"log"
	"time"
)

func main() {

	ctxBG := context.Background()
	dataChan := make(chan int)

	// Остановка по усовию
	// Пример: Пока не будет получено 20 значений из канала
	limitN := 20
	ctx1, cancel1 := context.WithCancel(ctxBG)
	defer cancel1()
	sig1 := make(chan int)
	go stopByCondition(ctx1, dataChan, limitN, sig1)

	for {
		select {
		case <-ctx1.Done():
			return
		case <-sig1:
			cancel1()
		default:
			dataChan <- time.Now().Nanosecond()
			time.Sleep(100 * time.Nanosecond)
		}
	}
}

func stopByCondition(ctx context.Context, ch <-chan int, n int, sig chan<- int) {
	for {
		select {
		case <-ctx.Done():
			log.Println("Остановка горутины stopByCondition")
			return
		case val, ok := <-ch:
			if !ok {
				log.Println("stopByCondition: канал закрты")
				return
			}
			n--
			log.Println(val, n)
			if n == 0 {
				sig <- 1
			}

		}
	}
}

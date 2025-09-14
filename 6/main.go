package main

import (
	"context"
	"log"
	"time"
)

func main() {

	// Остановка по уведомлению через канал (вообще сделал гибрит с условием)
	stopByChannel()
	// Остановка по условию
	stopByCondition()
}
func stopByChannel() {
	// Пример: Пока не будет получено 20 значений из канала
	n := 20 // лимит получаемых значений
	ctxBG := context.Background()
	dataChan := make(chan int)
	ctx, cancel1 := context.WithCancel(ctxBG)
	defer cancel1()
	sig := make(chan int, 1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Остановка горутины stopByChannel")
				return
			case val, ok := <-dataChan:
				if !ok {
					log.Println("stopByChannel: канал закрыт")
					return
				}
				n--
				log.Printf("stopByChannel получила %d, осталось получить %d значений.\n", val, n)
				if n == 0 {
					sig <- 1
				}

			}
		}
	}()

	// Передача данных
	for {
		select {
		case <-ctx.Done():
			return
		case <-sig:
			cancel1()
		default:
			dataChan <- time.Now().Nanosecond()
			time.Sleep(100 * time.Nanosecond)
		}
	}
}

func stopByCondition() {

	dataChan := make(chan int)
	n := 20
	chanClosed := make(chan struct{}) // Будем использовать для обратной связи из горутины, чтобы выйти из функции

	go func() {
		for n > 0 { // Основное условие, при срабатывании которого горутина будет выходить из цикла и посылать обратную связь
			val, ok := <-dataChan
			if !ok {
				log.Println("stopByCondition: канал закрыт")
				return
			}
			n--
			log.Printf("stopByCondition получила %d, осталось получить %d значений.\n", val, n)

		}
		log.Println("Остановка горутины stopByCondition")
		close(chanClosed)
	}()

	for {
		select {
		case <-chanClosed: // выходим из функции при получении сигнала обратной связи
			return
		case dataChan <- time.Now().Nanosecond(): // Вынес в кейс, т.к. оставляя в default получаем дедлок
			time.Sleep(100 * time.Nanosecond)
		}
	}
}

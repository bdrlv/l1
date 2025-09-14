package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {

	// Остановка по уведомлению через канал (вообще сделал гибрит с условием)
	stopByNotificationChannel()
	// Остановка по условию
	stopByCondition()
	// Остановка при закрытии канала данных
	stopByChannelClosing()
}
func stopByNotificationChannel() {
	var wg sync.WaitGroup

	n := 10 // лимит получаемых значений
	ctxBG := context.Background()
	dataChan := make(chan int)
	ctx, cancel1 := context.WithCancel(ctxBG)
	defer cancel1()
	sig := make(chan int, 1)

	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				log.Println("Остановка горутины stopByNotificationChannel")
				return
			case val, ok := <-dataChan:
				if !ok {
					log.Println("stopByNotificationChannel: канал закрыт")
					return
				}
				n--
				log.Printf("stopByNotificationChannel получила %d, осталось получить %d значений.\n", val, n)
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
			wg.Wait()
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
	var wg sync.WaitGroup

	dataChan := make(chan int)
	n := 10
	chanClosed := make(chan struct{}) // Будем использовать для обратной связи из горутины, чтобы выйти из функции

	wg.Add(1)
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
		wg.Done()
	}()

	for {
		select {
		case <-chanClosed: // выходим из функции при получении сигнала обратной связи
			wg.Wait()
			return
		case dataChan <- time.Now().Nanosecond(): // Вынес в кейс, т.к. оставляя в default получаем дедлок
			time.Sleep(100 * time.Nanosecond)
		}
	}
}

func stopByChannelClosing() {
	var wg sync.WaitGroup

	dataChan := make(chan int)

	wg.Add(1)
	go func() {
		for val := range dataChan {
			log.Printf("stopByChannelClosing получила %d\n", val)
		}
		log.Println("Канал datChan горутины stopByChannelClosing закрыт")
		log.Println("Остановка горутины stopByChannelClosing")
		wg.Done()
	}()

	for n := 0; n < 10; n++ {
		dataChan <- n
		time.Sleep(100 * time.Nanosecond)
	}
	log.Println("Закрываем канал datChan горутины stopByChannelClosing")
	close(dataChan)
	wg.Wait() // Использую wg.Wait(), чтобы функция дождалась вывода данных горутиной и корректно
}

package main

import (
	"log"
	"sync"
)

func main() {
	var mu sync.Mutex     // для инхронизации использую мьютакс
	var wg sync.WaitGroup // для корректного ожидания завершения горутин

	data := make(map[int]int)
	for i := 0; i < 10; i++ { // запускаю 10 горутин для примера
		wg.Add(1)
		go func() { // горутина заполняет ключ мапы в цикле
			for j := 0; j < 100; j++ {
				mu.Lock() // блокирую мапу
				log.Printf("key: %d, val: %d", i, j)
				data[i] = j // выполняю запись
				mu.Unlock() // разблокирую мапу
			}
			wg.Done() // вычитаем горутину из вейтгруппы, т.к. ее выполнение закончилось
		}()
	}

	wg.Wait()         // дожидаемся выполнения горутин
	log.Println(data) // смотрим результат
}

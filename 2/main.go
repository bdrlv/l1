package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	wg.Add(len(input))

	for _, item := range input {
		go CalculateSquareAndPrintResult(item, &wg)
	}

	wg.Wait()
}

func CalculateSquareAndPrintResult(item int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Квадрат для %v равен %v\n", item, item*item)
}

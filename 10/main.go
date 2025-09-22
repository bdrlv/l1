package main

import (
	"fmt"
	"math"
)

func main() {
	input := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	splitArrayIntoGroups(groups, input)

	for key, val := range groups {
		fmt.Printf("%d %v\n", key, val)
	}
}

func splitArrayIntoGroups(resultMap map[int][]float64, data []float64) {
	for _, item := range data {
		// отсечь дробную часть int() или math.Trunc()
		// выяснить знак >= 0 или < 0
		groupKey := getDigitGroup(item)

		// добавляем значение в мапу
		resultMap[groupKey] = append(resultMap[groupKey], item)

	}
}

func getDigitGroup(item float64) int {
	if item >= 0 {
		return int(item/10) * 10 // делю на 10 до отсечения дробной части, чтобы просто получить ровные десятки
	}
	return int(math.Ceil(item/10)) * 10 // при отрицательном значении числа дополнительно кругляю вверх, что соответсвовать условию (диапазон -20 включает значения от -20 до -29.9)
}

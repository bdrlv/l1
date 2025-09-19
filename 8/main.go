package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input64, output int64

	if len(os.Args) != 4 {
		fmt.Println("ошибка некорректного количества аргументов. Укажите число, позицию изменяемого бита, значение бита")
		os.Exit(1)
	}

	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("некорректный формат аргумента, %v. Значение должно быть положительным целочисенным и больше, либо равное 0.", err)
		os.Exit(1)
	}

	input64 = int64(input)

	targetBit, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("некорректный формат аргумента, %v. Значение должно быть положительным целочисенным и больше, либо равное 0.", err)
		os.Exit(1)
	}
	if targetBit < 0 || targetBit > 63 { // здесь проверяем, не выходит ли позиция изменяемого бита за диапазон int64
		fmt.Println("позиция целевого бита должна быть >= 0 и <= 63")
		os.Exit(1)
	}

	targetBitValue, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Printf("некорректный формат аргумента, %v. Значение должно быть положительным целочисенным и больше, либо равное 0.", err)
		os.Exit(1)
	}
	if targetBitValue < 0 || targetBitValue > 1 {
		fmt.Println("значение целевого бита должно быть 1 || 0")
		os.Exit(1)
	}

	// fmt.Println(input64, targetBit, targetBitValue)

	fmt.Printf("Ввод  - десятичное: %d, двоичное: %b\n", input64, input64)

	if targetBitValue == 1 { // проверяем, какое значение бита мы хотим установить
		output = input64 | (1 << uint(targetBit)) // устанавливаем в 1 с помощью OR
	} else {
		output = input64 &^ (1 << uint(targetBit)) // устанавливаем в 0 с помощью AND NOT
	}

	fmt.Printf("Вывод - Десятичное: %d, двоичное: %b\n", output, output)
}

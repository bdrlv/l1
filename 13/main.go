package main

import "fmt"

func main() {
	a := 2
	b := 3

	ver1(a, b)
	ver2(a, b)

}

func ver1(a, b int) {
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("Результат обмена через сложение/вычитание: a=%d, b=%d\n", a, b)
}

func ver2(a, b int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("Результат обмена через XOR: a=%d, b=%d\n", a, b)
}

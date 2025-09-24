package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	fmt.Println(intersection(a, b))
}

func intersection(a []int, b []int) []int {
	var result []int

	for _, itemA := range a {
		for _, itemB := range b {
			if itemA == itemB {
				result = append(result, itemA)
				break
			}
		}
	}

	return result
}

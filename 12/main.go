package main

import "fmt"

func main() {
	input := []string{"cat", "cat", "dog", "cat", "tree"}
	var result []string

	set := make(map[string]bool)
	for _, item := range input {
		set[item] = true
	}

	for key, val := range set {
		if val {
			result = append(result, key)
		}
	}

	// fmt.Println(input)
	fmt.Println(result)

}

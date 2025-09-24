package main

import "testing"

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{name: "Пример пересечения по заданию",
			a:        []int{1, 2, 3},
			b:        []int{2, 3, 4},
			expected: []int{2, 3},
		},
		{
			name:     "одно множество пустое",
			a:        []int{1, 2, 3},
			b:        []int{},
			expected: []int{},
		},
		{
			name:     "один элемент общий",
			a:        []int{10, 20, 30},
			b:        []int{30, 40, 50},
			expected: []int{30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := intersection(tt.a, tt.b)

			if !equal(result, tt.expected) {
				t.Errorf("пересечение %v, %v = %v, а ожидается %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) { // крайний случай
		return false
	}

	intersectsCounters := make(map[int]int) // считаем, какое число сколько раз встречается

	for _, valA := range a {
		intersectsCounters[valA]++
	}

	for _, valB := range b {
		intersectsCounters[valB]--
		if intersectsCounters[valB] < 0 { // если какое-то значение будет меньше 0, значит элементы встречаются с разной частотой
			return false
		}
	}

	return true
}

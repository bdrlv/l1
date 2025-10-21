package main

import "testing"

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		b        []int
		expected []int
	}{
		{name: "пример пересечения по заданию",
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
		{
			name:     "оба пустых",
			a:        []int{},
			b:        []int{},
			expected: []int{},
		},
		{
			name:     "идентичные заполненные",
			a:        []int{1, 2, 3},
			b:        []int{1, 2, 3},
			expected: []int{1, 2, 3},
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
	if len(a) != len(b) {
		return false
	}

	intersectsCounters := make(map[int]int)

	for _, valA := range a {
		intersectsCounters[valA]++
	}

	for _, valB := range b {
		intersectsCounters[valB]--
		if intersectsCounters[valB] < 0 {
			return false
		}
	}

	return true
}

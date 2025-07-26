package main

import (
	"fmt"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{
			name:     "Генерация 100 элементов",
			size:     100,
			expected: 100,
		},
		{
			name:     "Генерация 1 элемента",
			size:     1,
			expected: 1,
		},
		{
			name:     "Нулевой размер",
			size:     0,
			expected: 0,
		},
		{
			name:     "Отрицательный размер",
			size:     -5,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateRandomElements(tt.size)

			if len(result) != tt.expected {
				t.Errorf("generateRandomElements(%d) = слайс длиной %d, ожидалось %d",
					tt.size, len(result), tt.expected)
			}

			// Проверяем, что все числа положительные
			for i, value := range result {
				if value <= 0 {
					t.Errorf("Элемент %d равен %d, ожидалось положительное число", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Обычный слайс",
			input:    []int{1, 5, 3, 9, 2},
			expected: 9,
		},
		{
			name:     "Один элемент",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Все элементы одинаковые",
			input:    []int{5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "Максимум в начале",
			input:    []int{10, 1, 2, 3},
			expected: 10,
		},
		{
			name:     "Максимум в конце",
			input:    []int{1, 2, 3, 10},
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			if result != tt.expected {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Обычный слайс",
			input:    []int{1, 5, 3, 9, 2, 7, 4, 6},
			expected: 9,
		},
		{
			name:     "Один элемент",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "Пустой слайс",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Маленький слайс",
			input:    []int{1, 2, 3},
			expected: 3,
		},
		{
			name:     "Большой слайс",
			input:    make([]int, 1000),
			expected: 999,
		},
	}

	for i := range tests[4].input {
		tests[4].input[i] = i
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.input)
			if result != tt.expected {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMaximumVsMaxChunks(t *testing.T) {
	testSizes := []int{100, 1000, 10000}

	for _, size := range testSizes {
		t.Run(fmt.Sprintf("Сравнение на %d элементах", size), func(t *testing.T) {
			elements := generateRandomElements(size)

			maxSingle := maximum(elements)
			maxMultiple := maxChunks(elements)

			if maxSingle != maxMultiple {
				t.Errorf("Результаты не совпадают: maximum() = %d, maxChunks() = %d",
					maxSingle, maxMultiple)
			}
		})
	}
}

func BenchmarkMaximum(b *testing.B) {
	elements := generateRandomElements(1000000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maximum(elements)
	}
}

func BenchmarkMaxChunks(b *testing.B) {
	elements := generateRandomElements(1000000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maxChunks(elements)
	}
}

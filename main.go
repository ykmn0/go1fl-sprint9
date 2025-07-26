package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)

	for i := 0; i < size; i++ {
		elements[i] = rand.Intn(1000000) + 1
	}

	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data {
		if value > max {
			max = value
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	if len(data) == 1 {
		return data[0]
	}

	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		return maximum(data)
	}

	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		go func(chunkIndex int) {
			defer wg.Done()

			startIndex := chunkIndex * chunkSize
			endIndex := startIndex + chunkSize

			if chunkIndex == CHUNKS-1 {
				endIndex = len(data)
			}

			chunk := data[startIndex:endIndex]
			maxValues[chunkIndex] = maximum(chunk)
		}(i)
	}

	wg.Wait()

	// Находим максимум среди максимумов
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	elements := generateRandomElements(SIZE)
	fmt.Printf("Сгенерировано %d элементов\n", len(elements))

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(elements)
	elapsed := time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(elements)
	elapsed = time.Since(start).Milliseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}

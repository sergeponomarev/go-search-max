package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000 // 100 миллионов элементов
	CHUNKS = 8           // количество частей для параллельной обработки
)

// generateRandomElements генерирует слайс случайных целых положительных чисел заданного размера.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{} // Обработка крайнего случая: размер <= 0
	}

	data := make([]int, size)
	for i := range size {
		// rand.Int() возвращает неотрицательное псевдослучайное число int
		data[i] = rand.Int()
	}
	return data
}

// maximum находит максимальное число в слайсе.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0 // Обработка крайнего случая: пустой слайс
	}

	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks находит максимальное значение, разделяя слайс на CHUNKS частей
// и обрабатывая каждую часть в отдельной горутине.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0 // Обработка крайнего случая: пустой слайс
	}

	chunkSize := len(data) / CHUNKS

	// Если элементов меньше, чем CHUNKS, нет смысла делить, вызываем обычный maximum
	if chunkSize == 0 {
		return maximum(data)
	}

	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize

		// Последний чанк должен захватить все оставшиеся элементы (хвост),
		// если длина слайса не делится на CHUNKS без остатка
		if i == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		// Запускаем анонимную функцию асинхронно, передавая ей срез и индекс для записи результата
		go func(chunk []int, index int) {
			defer wg.Done()
			maxValues[index] = maximum(chunk)
		}(data[start:end], i)
	}

	// Ожидаем завершения всех 8 горутин
	wg.Wait()

	// Находим максимум среди максимумов каждого чанка
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел...\n", SIZE)

	// Генерация данных
	data := generateRandomElements(SIZE)
	fmt.Println("Генерация завершена.")

	fmt.Println("Ищем максимальное значение в один поток")
	startSingle := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(startSingle).Microseconds() // Замер в микросекундах

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n\n", maxSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startMulti := time.Now()
	maxMulti := maxChunks(data)
	elapsedMulti := time.Since(startMulti).Microseconds() // Замер в микросекундах

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxMulti, elapsedMulti)
}

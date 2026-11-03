package main

import (
	"testing"
)

// ===== Тесты generateRandomElements =====

func TestGenerateRandomElements_NormalSize(t *testing.T) {
	data := generateRandomElements(100)
	if len(data) != 100 {
		t.Errorf("Ожидалась длина 100, получено %d", len(data))
	}
}

func TestGenerateRandomElements_ZeroOrNegativeSize(t *testing.T) {
	if len(generateRandomElements(0)) != 0 {
		t.Errorf("Ожидался пустой слайс при size=0")
	}
	if len(generateRandomElements(-5)) != 0 {
		t.Errorf("Ожидался пустой слайс при size<0")
	}
}

func TestGenerateRandomElements_PositiveValues(t *testing.T) {
	data := generateRandomElements(1000)
	for i, v := range data {
		if v < 0 {
			t.Errorf("Элемент на позиции %d отрицательный: %d", i, v)
		}
	}
}

// ===== Тесты maximum =====

func TestMaximum_EmptySlice(t *testing.T) {
	if got := maximum([]int{}); got != 0 {
		t.Errorf("Для пустого слайса ожидался 0, получено %d", got)
	}
}

func TestMaximum_SingleElement(t *testing.T) {
	if got := maximum([]int{42}); got != 42 {
		t.Errorf("Ожидался 42, получено %d", got)
	}
}

func TestMaximum_MultipleElements(t *testing.T) {
	if got := maximum([]int{1, 5, 3, 99, 2, 7}); got != 99 {
		t.Errorf("Ожидался 99, получено %d", got)
	}
}

func TestMaximum_AllElementsEqual(t *testing.T) {
	data := []int{7, 7, 7, 7, 7}
	expected := 7
	if got := maximum(data); got != expected {
		t.Errorf("Для слайса с одинаковыми элементами ожидалось %d, получено %d", expected, got)
	}
}

// ===== Тесты maxChunks =====

func TestMaxChunks_EmptySlice(t *testing.T) {
	if got := maxChunks([]int{}); got != 0 {
		t.Errorf("Для пустого слайса ожидался 0, получено %d", got)
	}
}

func TestMaxChunks_SmallerThanChunks(t *testing.T) {
	// Слайс меньше, чем CHUNKS (8), должен корректно отработать
	if got := maxChunks([]int{10, 50, 5}); got != 50 {
		t.Errorf("Ожидался 50, получено %d", got)
	}
}

func TestMaxChunks_MatchesSingleThread(t *testing.T) {
	// Главный тест: результат многопоточного поиска должен совпадать с однопоточным
	data := generateRandomElements(50_000)
	expected := maximum(data)
	actual := maxChunks(data)

	if expected != actual {
		t.Errorf("Расхождение результатов! Ожидался %d, получено %d", expected, actual)
	}
}

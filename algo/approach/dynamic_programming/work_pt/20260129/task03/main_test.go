package main

import (
	"testing"
)

func TestCountWays(t *testing.T) {
	tests := []struct {
		n        int
		k        int
		expected int
	}{
		{1, 1, 1},  // только старт, уже на конечной клетке
		{2, 1, 1},  // 1 -> 2
		{2, 2, 1},  // 1 -> 2 (только прыжок на 1)
		{3, 1, 1},  // 1 -> 2 -> 3
		{3, 2, 2},  // 1->2->3 или 1->3
		{3, 3, 2},  // 1->2->3 или 1->3 (не 3!)
		{4, 2, 3},  // 1-2-3-4, 1-2-4, 1-3-4
		{4, 3, 4},  // 1-2-3-4, 1-2-4, 1-3-4, 1-4
		{5, 2, 5},  // последовательность Фибоначчи
		{8, 2, 21}, // пример из условия
		{5, 3, 7},  // проверка
		{5, 4, 8},  // проверка
		{5, 5, 8},  // проверка
	}

	for _, tt := range tests {
		result := countWays(tt.n, tt.k)
		if result != tt.expected {
			t.Errorf("countWays(%d, %d) = %d, want %d", tt.n, tt.k, result, tt.expected)
		}
	}
}

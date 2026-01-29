package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		N        int
		expected int
	}{
		{"N=1", 1, 0},
		{"N=2", 2, 1},
		{"N=3", 3, 1},
		{"N=4", 4, 2},
		{"N=5", 5, 3},
		{"N=6", 6, 2},
		{"N=7", 7, 3},
		{"N=8", 8, 3},
		{"N=9", 9, 2},
		{"N=10", 10, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ops, path := solve(tt.N)

			// Проверка количества операций
			if ops != tt.expected {
				t.Errorf("solve(%d) ops = %d, want %d", tt.N, ops, tt.expected)
			}

			// Проверка длины пути
			if len(path) != ops+1 {
				t.Errorf("solve(%d) path length = %d, but ops+1 = %d", tt.N, len(path), ops+1)
			}

			// Проверка начала и конца пути
			if len(path) > 0 {
				if path[0] != 1 {
					t.Errorf("solve(%d) path must start with 1, got %d", tt.N, path[0])
				}
				if path[len(path)-1] != tt.N {
					t.Errorf("solve(%d) path must end with %d, got %d", tt.N, tt.N, path[len(path)-1])
				}
			}

			// Проверка корректности каждого шага
			for i := 1; i < len(path); i++ {
				prev := path[i-1]
				curr := path[i]
				valid := false

				// Проверка допустимых операций
				switch curr {
				case prev + 1: // +1
					valid = true
				case prev * 2: // ×2
					valid = true
				case prev * 3: // ×3
					valid = true
				}

				if !valid {
					t.Errorf("solve(%d) invalid step from %d to %d", tt.N, prev, curr)
				}
			}
		})
	}
}

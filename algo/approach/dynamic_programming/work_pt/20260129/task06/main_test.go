package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		N, M     int
		expected int64
	}{
		// Тесты из условия
		{"N=1 M=1", 1, 1, 1},
		{"N=4 M=4", 4, 4, 2},
		{"N=2 M=3", 2, 3, 1},

		// N=3 M=2: (1,1)→(3,2)
		{"N=3 M=2", 3, 2, 1},

		// N=3 M=4: невозможно достичь (3,4)
		{"N=3 M=4", 3, 4, 0},

		// N=5 M=6: 3 пути (описано в комментариях)
		{"N=5 M=6", 5, 6, 3},

		// Другие тесты
		{"N=2 M=2", 2, 2, 0},
		{"N=6 M=5", 6, 5, 3},
		{"N=7 M=7", 7, 7, 6},
		{"N=50 M=50", 50, 50, 2459931987630},

		// Дополнительные проверки
		// N=4 M=3: невозможно (4,3) с данными ходами
		{"N=4 M=3", 4, 3, 0},
		// N=5 M=4: невозможно достичь (5,4)
		{"N=5 M=4", 5, 4, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solve(tt.N, tt.M)
			if result != tt.expected {
				t.Errorf("solve(%d, %d) = %d, want %d",
					tt.N, tt.M, result, tt.expected)
			}
		})
	}
}

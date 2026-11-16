package main

import "testing"

func Test_solve(t *testing.T) {
	tests := []struct {
		name string
		N    int
		want int
	}{
		{
			name: "N=1",
			N:    1,
			want: 8, // 1,2,3,4,5,6,7,9 (без 0 и 8)
		},
		{
			name: "N=2",
			N:    2,
			want: 16,
		},
		{
			name: "N=3",
			N:    3,
			want: 34,
		},
		{
			name: "N=4",
			N:    4,
			want: 72,
		},
		{
			name: "N=5",
			N:    5,
			want: 156,
		},
		{
			name: "N=10",
			N:    10,
			want: 6824,
		},
		{
			name: "N=20",
			N:    20,
			want: 13477272,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.N); got != tt.want {
				t.Errorf("solve(%d) = %d, want %d", tt.N, got, tt.want)
			}
		})
	}
}

func Benchmark_solve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(20) // максимальный размер
	}
}

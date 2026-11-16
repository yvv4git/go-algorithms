package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		want    int
		wantErr bool
	}{
		{name: "n=0", n: 0, want: 1, wantErr: false},
		{name: "n=1", n: 1, want: 1, wantErr: false},
		{name: "n=2", n: 2, want: 2, wantErr: false},
		{name: "n=3", n: 3, want: 3, wantErr: false},
		{name: "n=4", n: 4, want: 5, wantErr: false},
		{name: "n=5", n: 5, want: 8, wantErr: false},
		{name: "n=10", n: 10, want: 89, wantErr: false},
		{name: "negative", n: -1, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fib(tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("fib(%d) error = %v, wantErr = %v", tt.n, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fib(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestFibV2(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "n=0", n: 0, want: 1},
		{name: "n=1", n: 1, want: 1},
		{name: "n=5", n: 5, want: 8},
		{name: "n=10", n: 10, want: 89},
		{name: "n=20", n: 20, want: 10946},
		{name: "n=30", n: 30, want: 1346269},
		{name: "n=40", n: 40, want: 165580141},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := make(map[int]int)
			got := fibV2(tt.n, cache)
			if got != tt.want {
				t.Errorf("fibV2(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestFibV3(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "n=0", n: 0, want: 1},
		{name: "n=1", n: 1, want: 1},
		{name: "n=2", n: 2, want: 2},
		{name: "n=3", n: 3, want: 3},
		{name: "n=4", n: 4, want: 5},
		{name: "n=5", n: 5, want: 8},
		{name: "n=6", n: 6, want: 13},
		{name: "n=7", n: 7, want: 21},
		{name: "n=10", n: 10, want: 89},
		{name: "n=20", n: 20, want: 10946},
		{name: "n=30", n: 30, want: 1346269},
		{name: "negative n", n: -5, want: 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fibV3(tt.n)
			if got != tt.want {
				t.Errorf("fibV3(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestFibV3_Research(t *testing.T) {
	n := 10

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	fmt.Println("Заполняем DP-таблицу:")
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	fmt.Println("Готовая DP-таблица:")
	for i, v := range dp {
		fmt.Printf("  dp[%d] = F(%d) = %d\n", i, i, v)
	}
	fmt.Printf("Ответ: F(%d) = %d\n", n, dp[n])
}

func TestFibV3_ConsistentWithRecursion(t *testing.T) {
	for n := 0; n <= 15; n++ {
		want, err := fib(n)
		if err != nil {
			t.Fatalf("fib(%d) returned error: %v", n, err)
		}
		got := fibV3(n)
		if got != want {
			t.Errorf("fibV3(%d) = %d, fib(%d) = %d — mismatch", n, got, n, want)
		}
	}
}

func TestFibV2_ConsistentWithV3(t *testing.T) {
	for n := 0; n <= 40; n++ {
		cache := make(map[int]int)
		want := fibV2(n, cache)
		got := fibV3(n)
		if got != want {
			t.Errorf("fibV3(%d) = %d, fibV2(%d) = %d — mismatch", n, got, n, want)
		}
	}
}

func BenchmarkFibV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache := make(map[int]int)
		fibV2(30, cache)
	}
}

func BenchmarkFibV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibV3(30)
	}
}

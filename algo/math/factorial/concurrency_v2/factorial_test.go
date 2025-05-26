package main

import "testing"

func TestFactorialConcurrent(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Факториал 0",
			args: args{n: 0},
			want: 1,
		},
		{
			name: "Факториал 1",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "Факториал 2",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "Факториал 3",
			args: args{n: 3},
			want: 6,
		},
		{
			name: "Факториал 4",
			args: args{n: 4},
			want: 24,
		},
		{
			name: "Факториал 5",
			args: args{n: 5},
			want: 120,
		},
		{
			name: "Факториал 6",
			args: args{n: 6},
			want: 720,
		},
		{
			name: "Факториал 10",
			args: args{n: 10},
			want: 3628800,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactorialConcurrent(tt.args.n); got != tt.want {
				t.Errorf("FactorialConcurrent(%d) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}

// Benchmark
// go test -bench .

func BenchmarkFactorialConcurrent10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialConcurrent(10)
	}
}

func BenchmarkFactorialConcurrent15(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialConcurrent(15)
	}
}

func BenchmarkFactorialConcurrent20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialConcurrent(20)
	}
}

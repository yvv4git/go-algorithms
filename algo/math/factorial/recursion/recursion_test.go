package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorialByRecursion(t *testing.T) {
	testCases := []struct {
		InValue  int
		Expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.Expected, FactorialByRecursion(tc.InValue))
	}
}

func BenchmarkFactorialByRecursion(b *testing.B) {
	benchmarks := []struct {
		Name    string
		InValue int
	}{
		{"Pass 0", 0},
		{"Pass 1", 1},
		{"Pass 2", 2},
		{"Pass 3", 3},
		{"Pass 4", 4},
		{"Pass 5", 5},
		{"Pass 6", 6},
		{"Pass 7", 7},
		{"Pass 8", 8},
		{"Pass 9", 9},
		{"Pass 10", 10},
	}

	for _, bm := range benchmarks {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FactorialByRecursion(bm.InValue)
			}
		})
	}
}

// Benchmark
// go test -bench .
// go test -bench=. -run ^$

func BenchmarkFactorialConcurrent10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialByRecursion(10)
	}
}

func BenchmarkFactorialConcurrent15(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialByRecursion(15)
	}
}

func BenchmarkFactorialConcurrent20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FactorialByRecursion(20)
	}
}

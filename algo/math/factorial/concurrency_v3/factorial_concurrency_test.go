package main

import (
	"fmt"
	"testing"
)

// 5!=1х2х3х4х5=120.
func TestFactorialCuncurrency(t *testing.T) {
	numbers := []int{0, 1, 5, 7, 10, 20}
	factorialChannel := make(chan FactorialResult, len(numbers))

	for _, num := range numbers {
		go factorialConcurrency(num, factorialChannel)
	}

	for i := 0; i < len(numbers); i++ {
		fact := <-factorialChannel
		fmt.Printf("Factorial of %d: %d\n", fact.Num, fact.Factorial)
	}
}

// Benchmark
// go test -bench .
// go test -bench=. -run ^$

func BenchmarkFactorialConcurrent10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan FactorialResult, 1)
		go factorialConcurrency(10, ch)
		<-ch
	}
}

func BenchmarkFactorialConcurrent15(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan FactorialResult, 1)
		go factorialConcurrency(15, ch)
		<-ch
	}
}

func BenchmarkFactorialConcurrent20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan FactorialResult, 1)
		go factorialConcurrency(20, ch)
		<-ch
	}
}

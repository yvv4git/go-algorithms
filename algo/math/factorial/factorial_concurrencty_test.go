package factorial

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

type FactorialResult struct {
	Num       int
	Factorial int64
}

func factorialConcurrency(n int, ch chan FactorialResult) {
	var result int64 = 1

	for i := 1; i <= n; i++ {
		result *= int64(i)
	}

	ch <- FactorialResult{Num: n, Factorial: result}
}

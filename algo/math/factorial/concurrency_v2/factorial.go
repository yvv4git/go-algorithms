package main

import (
	"fmt"
	"sync"
)

func FactorialConcurrent(n int) int {
	if n == 0 || n == 1 {
		return 1 // Базовый случай(база рекурсии)
	}

	var result int
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		result = n * FactorialConcurrent(n-1)
		wg.Done()
	}()

	wg.Wait()
	return result
}

func main() {
	n := 5
	fmt.Printf("Вычисление факториала числа %d (с использованием горутины):\n", n)

	result := FactorialConcurrent(n)

	fmt.Printf("Результат: %d\n", result)
}

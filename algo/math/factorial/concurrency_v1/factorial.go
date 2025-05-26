package main

import (
	"fmt"
	"time"
)

type request struct {
	n      int
	result chan int
}

func FactorialConcurrentV2(n int) int {
	// Создаем общий канал для всех запросов
	workerChan := make(chan request)

	// Запускаем воркер, который обрабатывает запросы
	go factorialWorker(workerChan)

	// Создаем канал для получения результата
	resultChan := make(chan int)

	// Отправляем запрос воркеру
	workerChan <- request{n, resultChan}

	// Ждем результат
	return <-resultChan
}

func factorialWorker(workerChan chan request) {
	for req := range workerChan {
		if req.n == 0 || req.n == 1 {
			req.result <- 1
			continue
		}

		// Создаем канал для подзадачи
		subResultChan := make(chan int)

		// Отправляем подзадачу в тот же воркер
		workerChan <- request{req.n - 1, subResultChan}

		// Получаем результат подзадачи, умножаем и отправляем обратно
		subResult := <-subResultChan
		req.result <- req.n * subResult
	}
}

func main() {
	n := 5
	fmt.Printf("Вычисляем факториал числа %d (конкурентно):\n", n)
	result := FactorialConcurrentV2(n)
	fmt.Printf("Результат: %d\n", result)

	// Ждем немного, чтобы воркер завершил обработку (необязательно, просто для демонстрации)
	time.Sleep(100 * time.Millisecond)
}

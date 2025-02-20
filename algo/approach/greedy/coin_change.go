//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// Функция размена монет с использованием жадного алгоритма
func coinChange(coins []int, amount int) map[int]int {
	/*
		METHOD: Greedy algorithm
		1. Сортируем номиналы монет по убыванию.
		2. На каждом шаге берём максимальное количество самой крупной доступной монеты.
		3. Повторяем, пока сумма не будет покрыта.

		TIME COMPLEXITY: O(nlogn) (из-за сортировки)
		SPACE COMPLEXITY: O(n) (для хранения результата)
	*/
	// Сортируем монеты по убыванию номинала
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	// Словарь для хранения количества монет каждого номинала
	change := make(map[int]int)

	for _, coin := range coins {
		if amount >= coin {
			count := amount / coin // Берём максимум возможных монет этого номинала
			amount -= count * coin // Вычитаем из оставшейся суммы
			change[coin] = count   // Записываем количество использованных монет
		}
	}

	// Если сумма не может быть разменяна (например, нет 1-копеечных монет), возвращаем пустой результат
	if amount > 0 {
		return nil
	}

	return change
}

// Дано несколько номиналов монет. Нужно разменять сумму S с минимальным количеством монет.
func main() {
	// Номиналы монет
	coins := []int{1, 5, 10, 25}
	amount := 63 // Размениваем 63 копейки

	// Вызываем жадный алгоритм
	result := coinChange(coins, amount)

	// Выводим результат
	if result == nil {
		fmt.Println("Невозможно разменять данную сумму!")
	} else {
		fmt.Println("Размен с минимальным количеством монет:")
		for coin, count := range result {
			fmt.Printf("Монета %d копеек: %d шт.\n", coin, count)
		}
	}
}

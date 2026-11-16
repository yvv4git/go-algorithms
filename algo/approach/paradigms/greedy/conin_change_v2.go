//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// TASK: Имеется определенное количество монет разных номиналов. Необходимо выдать сдачу или вернуть nil, если это невозможно.
func coinChange(coins []int, amount int) map[int]int {
	/*
		APPROACH: Greedy
		1. Начинаем с самой крупной доступной монеты.
		2. Берем максимальное количество этой монеты.
		3. Переходим к следующей меньшей монете.
		4. Повторяем процесс, пока сдача не будет полностью выдана.

		TIME COMPLEXITY: O(n * m)
		- Сортировка монет занимает O(m log m), где m — количество номиналов монет.
		- Основной цикл проходит по всем m номиналам.
		- В худшем случае для каждой монеты выполняется деление и вычитание O(n) раз.
		- Итоговая сложность O(m log m + m * n), но так как сортировка выполняется один раз и обычно меньше, чем основной цикл, это упрощается до O(n * m).

		SPACE COMPLEXITY: O(m)
		- Мы используем дополнительную память для хранения карты `change`, где в худшем случае будет O(m) записей (по одной на каждый номинал).
		- Остальные переменные занимают O(1) памяти.
		- Итоговая сложность O(m).
	*/
	sort.Sort(sort.Reverse(sort.IntSlice(coins))) // Сортируем монеты по убыванию

	change := make(map[int]int) // Хранит {монета : количество}
	remaining := amount         // Оставшаяся сумма для размена

	for _, coin := range coins {
		if remaining >= coin {
			count := remaining / coin // Сколько таких монет можем взять
			change[coin] = count
			remaining -= count * coin
		}
	}

	// Если размен невозможен
	if remaining > 0 {
		fmt.Println("Невозможно разменять сумму с данными монетами")
		return nil
	}

	return change
}

func main() {
	coins := []int{1, 5, 10, 25} // Доступные номиналы монет
	amount := 63                 // Сумма для размена

	result := coinChange(coins, amount)

	if result != nil {
		fmt.Println("Размен возможен:")
		for coin, count := range result {
			fmt.Printf("%d: %d шт.\n", coin, count)
		}
	} else {
		fmt.Println("Размен невозможен")
	}
}

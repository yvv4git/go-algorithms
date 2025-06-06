# Алгоритм Dynamic Programming (Динамическое программирование)

## Введение

Динамическое программирование (Dynamic Programming, DP) — это метод решения задач, основанный на разбиении задачи на подзадачи, 
решение которых используется для построения решения всей задачи. Этот подход применяется, 
когда задача имеет свойство перекрывающихся подзадач, то есть когда решение подзадач может быть использовано несколько раз.

## Принцип работы

1. **Разбиение на подзадачи**: Исходная задача разбивается на подзадачи, которые могут быть решены независимо.
2. **Мемоизация (или табуляция)**: Результаты решения подзадач сохраняются в таблице, чтобы избежать повторного вычисления тех же подзадач.
3. **Объединение решений**: Решение исходной задачи строится на основе решений подзадач.

### Когда использовать Dynamic Programming?

- Когда задача может быть разбита на перекрывающиеся подзадачи.
- Когда подзадачи решаются несколько раз в процессе вычислений, и можно избежать повторных вычислений, сохраняя результаты промежуточных вычислений.

## Пример: Задача о рюкзаке (0/1 Knapsack Problem)

### Описание задачи:

Дано множество предметов, каждый из которых имеет вес и стоимость. Нужно выбрать набор предметов, которые можно положить в рюкзак ограниченной вместимости, так чтобы суммарная стоимость этих предметов была максимальной.

### Принцип решения с использованием динамического программирования:

1. Разбиваем задачу на подзадачи, где каждая подзадача решается для определенной вместимости рюкзака и набора предметов.
2. Используем таблицу для хранения решений этих подзадач.
3. Для каждой подзадачи принимаем решение: включить ли текущий предмет в рюкзак или не включать.

### Пример реализации на Go:

```go
package main

import "fmt"

// Функция для решения задачи о рюкзаке с использованием динамического программирования
func knapsack(weights []int, values []int, capacity int) int {
	n := len(weights)
	// Создаем таблицу для хранения решений подзадач
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	// Заполняем таблицу dp
	for i := 1; i <= n; i++ {
		for w := 1; w <= capacity; w++ {
			if weights[i-1] <= w {
				// Если текущий предмет помещается в рюкзак
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-weights[i-1]]+values[i-1])
			} else {
				// Если текущий предмет не помещается в рюкзак
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// Ответ находится в правом нижнем углу таблицы dp
	return dp[n][capacity]
}

// Вспомогательная функция для нахождения максимального значения
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Пример предметов
	weights := []int{2, 3, 4, 5}
	values := []int{3, 4, 5, 6}
	capacity := 5

	// Находим максимальную стоимость, которую можно унести
	result := knapsack(weights, values, capacity)

	// Выводим результат
	fmt.Printf("Максимальная стоимость: %d\n", result)
}

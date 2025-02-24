# Алгоритм подсчета элементов в массиве (Counting)


## Описание

Алгоритм подсчета элементов используется для нахождения частоты появления каждого уникального элемента в массиве. 
Это полезно для различных задач, таких как подсчет повторений или анализ распределения элементов.

## Принцип работы

1. **Инициализация карты**: Мы создаем пустую карту (`map`), где ключами будут элементы массива, а значениями — количество вхождений этих элементов.
2. **Подсчет элементов**: Проходим по каждому элементу массива и увеличиваем счетчик соответствующего элемента в карте.
3. **Результат**: В конце получается карта, где каждый элемент из массива связан с количеством его вхождений.

## Время работы

- **Время работы**: Время работы алгоритма составляет `O(n)`, где `n` — количество элементов в массиве, так как каждый элемент проходит только один раз.
- **Пространственная сложность**: Пространственная сложность составляет `O(k)`, где `k` — количество уникальных элементов в массиве, так как для каждого уникального элемента мы храним его частоту в карте.

## Пример реализации на Go

```go
package main

import "fmt"

// Функция для подсчета элементов в массиве
func countElements(arr []int) map[int]int {
	// Создаем пустую map для подсчета
	countMap := make(map[int]int)

	// Проходим по каждому элементу в массиве
	for _, element := range arr {
		// Увеличиваем значение элемента на 1
		countMap[element]++
	}

	// Возвращаем map с подсчитанными элементами
	return countMap
}

func main() {
	// Пример использования
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	counts := countElements(numbers)

	// Выводим результаты
	for key, value := range counts {
		fmt.Printf("Элемент %d встречается %d раз(а)\n", key, value)
	}
}

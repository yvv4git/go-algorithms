//go:build ignore

package main

import (
	"fmt"
	"sort"
)

// Структура для хранения предметов
type Item struct {
	weight float64 // Вес предмета
	value  float64 // Ценность предмета
	ratio  float64 // Ценность за единицу веса
}

// Функция для решения задачи о дробном рюкзаке с использованием жадного алгоритма
func fractionalKnapsack(items []Item, capacity float64) float64 {
	/*
		METHOD: Greedy algorithm
		1. Вычисляем "ценность за единицу веса" (value/weight) для каждого предмета.
		2. Сортируем предметы по убыванию этого показателя.
		3. Добавляем в рюкзак предметы, пока есть место, а если не влезает – берем дробную часть.

		TIME COMPLEXITY: O(nlogn) (из-за сортировки)
		SPACE COMPLEXITY: O(1) (используем фиксированное количество памяти)
	*/

	// Вычисляем ratio (ценность за единицу веса)
	for i := range items {
		items[i].ratio = items[i].value / items[i].weight
	}

	// Сортируем предметы по убыванию ratio
	sort.Slice(items, func(i, j int) bool {
		return items[i].ratio > items[j].ratio
	})

	totalValue := 0.0    // Общая ценность в рюкзаке
	currentWeight := 0.0 // Текущий вес рюкзака

	// Проходим по предметам, добавляя их в рюкзак
	for _, item := range items {
		if currentWeight+item.weight <= capacity {
			// Берем предмет целиком
			currentWeight += item.weight
			totalValue += item.value
		} else {
			// Берем только часть предмета
			remainingWeight := capacity - currentWeight
			totalValue += item.ratio * remainingWeight
			break // Рюкзак заполнен
		}
	}

	return totalValue
}

// Дано N предметов с весом wᵢ и ценностью vᵢ. Нужно максимизировать ценность в рюкзаке ёмкостью W,
// при этом можно брать дробные части предметов.
func main() {
	// Пример предметов (вес, ценность)
	items := []Item{
		{weight: 10, value: 60},
		{weight: 20, value: 100},
		{weight: 30, value: 120},
	}

	capacity := 50.0 // Вместимость рюкзака

	// Вызываем жадный алгоритм
	maxValue := fractionalKnapsack(items, capacity)

	// Выводим результат
	fmt.Printf("Максимальная ценность рюкзака: %.2f\n", maxValue)
}

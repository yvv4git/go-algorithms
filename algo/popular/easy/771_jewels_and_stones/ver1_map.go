package main

import (
	"fmt"
)

func numJewelsInStones(jewels string, stones string) int {
	/*
		METHOD: Map
		Описание метода:
		Мы используем карту (map) для хранения драгоценностей. Это позволяет нам быстро проверять,
		является ли каждый символ в строке stones драгоценностью. Мы заполняем карту символами из
		строки jewels, а затем перебираем строку stones и проверяем каждый символ на наличие в карте.

		TIME COMPLEXITY: O(n)
		- Заполнение карты драгоценностями: O(m), где m — длина строки jewels.
		- Проверка каждого символа в строке stones: O(n), где n — длина строки stones.
		Таким образом, общая временная сложность составляет O(m + n), что можно упростить до O(n),
		если предположить, что n обычно больше или равно m.

		SPACE COMPLEXITY: O(n)
		- Карта для хранения драгоценностей: O(m), где m — длина строки jewels.
		Таким образом, общая пространственная сложность составляет O(m), что можно упростить до O(n),
		если предположить, что n обычно больше или равно m.
	*/
	// Создаем карту (map) для хранения драгоценностей
	jewelMap := make(map[rune]bool)

	// Заполняем карту драгоценностями
	for _, jewel := range jewels {
		jewelMap[jewel] = true
	}

	// Инициализируем счетчик для подсчета драгоценных камней
	count := 0

	// Перебираем каждый символ в строке stones
	for _, stone := range stones {
		// Проверяем, является ли текущий камень драгоценностью
		if jewelMap[stone] {
			// Если да, увеличиваем счетчик
			count++
		}
	}

	// Количество драгоценных камней
	return count
}

func main() {
	// Пример использования функции
	jewels := "aA"
	stones := "aAAbbbb"
	result := numJewelsInStones(jewels, stones)
	fmt.Println(result) // Вывод: 3
}

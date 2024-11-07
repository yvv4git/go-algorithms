package main

func numJewelsInStonesV2(jewels string, stones string) int {
	/*
		METHOD: Brute Force
		Мы используем два вложенных цикла: внешний цикл перебирает каждый символ в строке stones,
		а внутренний цикл проверяет, содержится ли этот символ в строке jewels. Если символ содержится
		в jewels, увеличиваем счетчик.

		TIME COMPLEXITY: O(n * m)
		- Внешний цикл: O(n), где n — длина строки stones.
		- Внутренний цикл: O(m), где m — длина строки jewels.
		Таким образом, общая временная сложность составляет O(n * m).

		SPACE COMPLEXITY: O(1)
		- Дополнительная память не используется, поэтому пространственная сложность составляет O(1).
	*/
	// Инициализируем счетчик для подсчета драгоценных камней
	count := 0

	// Перебираем каждый символ в строке stones
	for _, stone := range stones {
		// Внутренний цикл для проверки, содержится ли текущий камень в строке jewels
		for _, jewel := range jewels {
			if stone == jewel {
				// Если да, увеличиваем счетчик
				count++
				break // Прерываем внутренний цикл, так как камень уже найден
			}
		}
	}

	// Возвращаем количество драгоценных камней
	return count
}

// func main() {
// 	// Пример использования функции
// 	jewels := "aA"
// 	stones := "aAAbbbb"
// 	result := numJewelsInStones(jewels, stones)
// 	fmt.Println(result) // Вывод: 3
// }

package main

func minCostToMoveChips(position []int) int {
	/*
		МЕТОД: Linear Scan
			Мы можем посчитать количество фишек на четных и нечетных позициях в массиве.
			Если количество фишек на четных позициях меньше, чем на нечетных,
			мы можем переместить все фишки с четных позиций на одну и ту же позицию и вернуть количество фишек на четных позициях.
			В противном случае, мы можем переместить все фишки с нечетных позиций на одну и ту же позицию и вернуть количество фишек на нечетных позициях.

			СЛОЖНОСТЬ ПО ВРЕМЕНИ: O(n)
			СЛОЖНОСТЬ ПО ПАМЯТИ: O(1)
	*/
	even := 0 // Счетчик фишек на четных позициях
	odd := 0  // Счетчик фишек на нечетных позициях

	// Проходим по всем позициям и считаем количество фишек на четных и нечетных позициях
	for _, pos := range position {
		if pos%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	// Возвращаем минимальное количество фишек, которые нужно переместить
	if even < odd {
		return even
	}

	return odd
}

func main() {
	position := []int{1, 2, 3}             // Пример массива позиций фишек
	result := minCostToMoveChips(position) // Вычисляем минимальную стоимость перемещения
	println(result)                        // Выводим результат
}

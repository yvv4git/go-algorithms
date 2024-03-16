package main

import "fmt"

// Функция для поиска элемента, встречающегося один раз
func singleNumber(nums []int) int {
	/*
		METHOD: Bitwise
		Для решения этой задачи используется подход, основанный на битовых операциях.
		Битовые операции используются для отслеживания вхождений каждого числа.
		Для этого используются две переменные ones и twos, которые используются для отслеживания чисел,
		встречающихся один и два раза соответственно.

		В процессе прохода по массиву чисел для каждого числа выполняются следующие операции:
		1. ones обновляется с использованием XOR (исключающее ИЛИ) для текущего числа и отрицания twos.
		Это гарантирует, что биты в ones будут установлены только для чисел, которые встречаются один раз.

		2. twos обновляется аналогичным образом, но используется уже обновленное значение ones.
		Это гарантирует, что биты в twos будут установлены только для чисел, которые встречаются два раза.

		Таким образом, в конце процесса в ones будет только один бит, соответствующий числу,
		которое встречается один раз, и все остальные биты будут сброшены.

		TIME COMPLEXITY: O(n), где n - количество элементов в массиве, так как мы проходим по каждому элементу ровно один раз.

		SPACE COMPLEXITY: O(1), так как мы используем несколько переменных для отслеживания вхождений элементов, и это не зависит от размера входных данных.
	*/
	// Инициализируем переменные для отслеживания вхождений
	ones, twos := 0, 0

	// Проходим по каждому элементу в массиве
	for _, num := range nums {
		// Обновляем переменную ones, учитывая текущий элемент
		ones = (ones ^ num) & ^twos
		// Обновляем переменную twos, учитывая текущий элемент
		twos = (twos ^ num) & ^ones
	}

	// В конце концов, в ones будет только один бит, соответствующий числу,
	// которое встречается один раз, и все остальные биты будут сброшены.
	return ones
}

func main() {
	// Пример использования функции
	nums := []int{2, 2, 3, 2}
	fmt.Println(singleNumber(nums)) // Вывод: 3
}

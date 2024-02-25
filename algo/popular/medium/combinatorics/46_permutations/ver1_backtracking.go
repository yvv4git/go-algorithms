package main

import "fmt"

// Функция для генерации всех перестановок
func permute(nums []int) [][]int {
	/*
		METHOD: Backtracking
		TIME COMPLEXITY: O(n*n!), где n - количество элементов в исходном массиве. Это связано с тем, что для каждого элемента мы генерируем все возможные перестановки, и таких перестановок n! (факториал от n), где n - количество элементов в массиве.
		SPACE COMPLEXITY: O(n*n!), также из-за хранения всех возможных перестановок. Каждая перестановка занимает O(n) места, и их n!, поэтому общее количество занимаемого места составляет O(n*n!).

		Подход, используемый для решения задачи "46. Permutations", называется "обратным отслеживанием" или "backtracking".
		Этот подход основан на рекурсивном переборе всех возможных комбинаций элементов.

		В данном случае, мы начинаем с пустого набора и постепенно добавляем элементы в набор.
		Когда набор становится достаточно большим, мы добавляем его в результат.
		Если мы добавили слишком много элементов, мы откатываемся назад и пытаемся добавить другие элементы.
		Этот процесс продолжается до тех пор, пока мы не перебрали все возможные комбинации.

		Backtracking является мощным инструментом для решения задач, в которых требуется перебор всех возможных комбинаций или вариантов.
		Он широко используется в различных алгоритмах, таких как генерация всех возможных путей в лабиринте,
		решение судоку или другие задачи, где требуется перебор всех возможных вариантов.
	*/
	// Инициализация результирующего среза для хранения всех перестановок
	result := make([][]int, 0)
	// Вызов вспомогательной функции для генерации перестановок
	backtrack(&result, nums, 0)
	return result
}

// Вспомогательная функция для генерации перестановок с использованием обратного отслеживания
func backtrack(result *[][]int, nums []int, start int) {
	// Если начальный индекс равен длине среза, то перестановка завершена
	if start == len(nums) {
		// Создание копии текущей перестановки и добавление ее в результирующий срез
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
	} else {
		// Проход по всем возможным индексам для текущего уровня рекурсии
		for i := start; i < len(nums); i++ {
			// Меняем местами текущий элемент с элементом на позиции start
			nums[start], nums[i] = nums[i], nums[start]
			// Рекурсивный вызов для следующего уровня рекурсии
			backtrack(result, nums, start+1)
			// Возвращение исходного порядка элементов для следующей итерации
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
}

func main() {
	// Пример использования функции permute
	nums := []int{1, 2, 3}
	permutations := permute(nums)
	fmt.Println(permutations)
}
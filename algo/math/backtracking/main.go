package main

import "fmt"

func backtrack(candidates []int, target int, temp []int, result *[][]int, start int) {
	/*
		METHOD: Backtracking
		Time complexity: O(n ^ T), где n - количество элементов в наборе, t - это целевое значение.

		Почему так?
		В худшем случае, когда каждое число может быть использовано несколько раз, мы можем вызвать функцию backtrack для каждого числа в наборе.
		Это происходит T раз, потому что мы вызываем функцию для каждого числа в наборе T раз. Таким образом, мы получаем N^T в худшем случае.

		Обратите внимание, что это очень грубая оценка, так как в реальности не все комбинации будут рассматриваться.
		Например, если мы имеем набор [1, 2, 3] и целевое значение 6, только комбинации [1, 2, 3] будет рассматриваться,
		а комбинации [1, 2, 1, 2] и [2, 2, 2] не будут рассматриваться, поскольку они превышают целевое значение.
	*/
	if target < 0 {
		return
	} else if target == 0 {
		*result = append(*result, append([]int{}, temp...))
	} else {
		for i := start; i < len(candidates); i++ {
			temp = append(temp, candidates[i])
			backtrack(candidates, target-candidates[i], temp, result, i)
			temp = temp[:len(temp)-1]
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var temp []int
	backtrack(candidates, target, temp, &result, 0)
	return result
}

func main() {
	/*
		Backtracking может быть использован для решения различных задач, например, для генерации всех возможных комбинаций элементов.
		Вот пример на языке Go, где мы генерируем все возможные комбинации из заданного набора элементов.

		В этом примере мы используем функцию backtrack для генерации всех возможных комбинаций, которые дают сумму target.
		Функция backtrack использует рекурсию для проверки всех возможных комбинаций.
		Если сумма текущей комбинации больше target, мы "откатываемся" (backtrack) и пробуем другую комбинацию.
		Если сумма текущей комбинации равна target, мы добавляем ее в результат.
	*/
	candidates := []int{2, 3, 6, 7} // Сумма каких чисел из набора даст target?
	target := 7
	fmt.Println(combinationSum(candidates, target))
}

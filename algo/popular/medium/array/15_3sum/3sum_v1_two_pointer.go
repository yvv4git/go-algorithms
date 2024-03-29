package _5_3sum

import "sort"

func threeSumV1(nums []int) [][]int {
	/*
		METHOD: Sort & two pointer.
		TIME COMPLEXITY: O(n^2)
		SPACE COMPLEXITY: O(1)

		Асимптотическая сложность этого алгоритма составляет O(n^2), где n - это количество элементов в массиве.

		Это обусловлено двумя основными операциями:
		1. Сортировка массива: Это происходит за O(n log n) времени.
		2. Два указателя: Это происходит за O(n^2) времени, так как для каждого элемента мы проходим по всем остальным элементам.

		Таким образом, общий асимптотический сложность составляет O(n log n) + O(n^2),
		который можно упростить до O(n^2), так как O(n log n) << O(n^2) при больших n.


		Решить эту задачу можно несколькими способами. Ниже представлены некоторые из них:
		1. Брутфорс.
		Это самый простой способ решения задачи, но он не эффективен для больших массивов, поскольку его временная сложность составляет O(n^3).

		2. Хэш-таблицы.
		Вы можете использовать хэш-таблицы для хранения суммы двух элементов и их индексов.
		Этот подход имеет временную сложность O(n^2), но он эффективнее брутфорса.

		3. Два указателя.
		Этот подход использует две переменные, одна из которых начинается с начала массива, а другая - с конца.
		Если сумма элементов меньше нуля, то увеличиваем левый указатель, если сумма больше нуля, то уменьшаем правый указатель.
		Если сумма равна нулю, то мы нашли одну из троек. Этот подход также имеет временную сложность O(n^2),
		но он эффективнее брутфорса и хэш-таблиц.

		4. Бинарный поиск.
		Вы можете использовать бинарный поиск для поиска пары элементов, сумма которых дает требуемую сумму.
		Этот подход имеет временную сложность O(n log n), но он эффективнее брутфорса и хэш-таблиц.

		5. Сортировка и два указателя.
		Вы можете отсортировать массив и использовать два указателя для нахождения троек.
		Этот подход имеет временную сложность O(n^2), но он эффективнее брутфорса и хэш-таблиц.
	*/
	sort.Ints(nums) // O(n log n)
	var result [][]int

	for i := 0; i < len(nums)-2; i++ { // O(n^2)
		if i > 0 && nums[i] == nums[i-1] { // Проверяем, что нет повторяющихся элементов
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

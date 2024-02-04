package top_k_frequent

func topKFrequent(nums []int, k int) []int {
	/*
		TASK: Найти k чаще всего встречающихся элемента в массиве целых чисел.
		Дан целочисленный массив и целое число k. Необходимо вернуть k часто встречающихся элементов. Порядок чисел в ответе не имеет значения.
		Гарантируется, что ответ уникальный. То есть несколько элементов не могут встречаться одинаковое количество раз.

		EXAMPLE:
		Ввод: nums = [1,1,1,2,2,3], k = 2
		Вывод: [1,2]

		Ввод: nums = [1], k = 1
		Вывод: [1]

		Times complexity: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	dict := make(map[int]int)

	for _, num := range nums {
		dict[num]++
	}

	arr := [][]int{}
	for key, value := range dict {
		arr = append(arr, []int{key, value})
	}

	n := len(arr)
	quickSelect(arr, 0, n-1, n-k)

	res := []int{}
	for i := n - k; i < n; i++ {
		res = append(res, arr[i][0])
	}

	return res
}

func quickSelect(arr [][]int, start, end, k int) {
	if start < end {
		pivot := partition(arr, start, end)
		if pivot == k {
			return
		} else if pivot < k {
			quickSelect(arr, pivot+1, end, k)
		} else {
			quickSelect(arr, start, pivot-1, k)
		}
	}
}

func partition(arr [][]int, start, end int) int {
	idx := start - 1
	pivot := end

	for i := start; i < end; i++ {
		if arr[i][1] < arr[pivot][1] {
			idx++
			arr[idx], arr[i] = arr[i], arr[idx]
		}
	}
	idx++
	arr[idx], arr[pivot] = arr[pivot], arr[idx]
	return idx
}

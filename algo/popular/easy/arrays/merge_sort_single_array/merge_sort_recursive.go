package merge_sort_single_array

func mergeSort(items []int) []int {
	/*
		Сложность: O(n*log(n))
		Память: O(1) - мы ничего никуда не копируем, используем имеющуюся память.
	*/
	if len(items) < 2 {
		return items
	}

	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])

	return merge(first, second)
}

func merge(a []int, b []int) []int {
	var final []int
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}

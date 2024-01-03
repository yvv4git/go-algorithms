package count

func countingSort(list []int) {
	if len(list) == 0 {
		return
	}

	min := 0
	max := 0

	for i := 0; i < len(list); i++ {
		if min > list[i] {
			min = list[i]
		}

		if max < list[i] {
			max = list[i]
		}
	}

	count := make([]int, max-min+1) // среди чисел еще присутствует ноль, потому надо прибавить 1
	for _, x := range list {
		count[x-min]++
	}

	z := 0
	for i, c := range count {
		if c > 0 {
			list[z] = i + min // используем закономерность, что index в мапе мы строили относительно min элемента
			z++
		}
	}
}

package _252_meeting_rooms

//type Interval struct {
//	Start int
//	End   int
//}

func canAttendMeetingsV2(intervals []Interval) bool {
	/*
		METHOD: Loop after sort.
		TIME COMPLEXITY: O(n log n)
		SPACE COMPLEXITY: O(1)

		В этом коде мы используем быструю сортировку для сортировки начал и концов встреч.
		Затем мы проверяем, есть ли пересечения в отсортированных массивах.
		Если есть, то это означает, что человек не может принять все встречи.

		Обратите внимание, что временная сложность этого алгоритма также будет O(n log n),
		так как мы используем быструю сортировку, которая имеет сложность O(n log n).
		Однако, в отличие от предыдущего алгоритма, этот алгоритм не требует дополнительной памяти для хранения отсортированных начал и концов встреч.
	*/
	startTimes := make([]int, len(intervals))
	endTimes := make([]int, len(intervals))

	for i, interval := range intervals {
		startTimes[i] = interval.Start
		endTimes[i] = interval.End
	}

	quickSort(startTimes, 0, len(startTimes)-1)
	quickSort(endTimes, 0, len(endTimes)-1)

	for i := 1; i < len(startTimes); i++ {
		if startTimes[i] < endTimes[i-1] {
			return false
		}
	}

	return true
}

// Рекурсивный алгоритм быстрой сортировки.
func quickSort(arr []int, low, high int) {
	/*
		TIME COMPLEXITY: O(n log n)

		Асимптотическая сложность алгоритма быстрой сортировки в среднем и наилучшем случае составляет O(n log n).
		Это означает, что время сортировки растет линейно пропорционально логарифму количества элементов в массиве.

		Однако, в худшем случае (когда самый маленький или самый большой элемент всегда выбирается в качестве опорного),
		асимптотическая сложность может достигать O(n^2). Это происходит, когда на каждом шаге мы разделяем массив на две части,
		одна из которых содержит n-1 элементов, а другая - 0 элементов.
	*/
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

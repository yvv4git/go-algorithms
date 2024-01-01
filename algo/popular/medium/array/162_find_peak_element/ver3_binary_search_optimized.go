package _62_find_peak_element

// Функция для поиска пика в массиве
func findPeakV3(nums []int) int {
	start, end := 0, len(nums)-1

	for start < end {
		mid := start + (end-start)/2

		// Если следующий элемент больше текущего, то пик находится в правой половине
		if nums[mid] < nums[mid+1] {
			start = mid + 1
		} else {
			// Иначе пик находится в левой половине
			end = mid
		}
	}

	// В конце цикла start и end будут указывать на один и тот же элемент, который является пиком
	return start
}

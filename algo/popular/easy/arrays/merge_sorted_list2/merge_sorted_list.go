package main

// Моя собственная реализация объединения двух отсортированных массивов.
func mergeTwoSortedList(num1, num2 []int) (result []int) {
	len1, len2 := len(num1), len(num2)
	i := 0
	j := 0

	result = make([]int, 0, len1+len2)
	for i < len1 && j < len2 {
		if num1[i] <= num2[j] {
			result = append(result, num1[i])
			i++
		} else {
			result = append(result, num2[j])
			j++
		}
	}

	// Надо понимать, что если i==len1, значит все элементы из num1 были использованы.
	// А, если j==len2, значит все элементы из num2 были использованы.
	// Но, когда i<len1 или j<len2, значит в одном из них остались элементы, которые можно докинуть в конец результирующего слайса.
	if i < len1 {
		result = append(result, num1[i:]...)
	} else if j < len2 {
		result = append(result, num2[j:]...)
	}

	return
}

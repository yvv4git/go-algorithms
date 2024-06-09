package main

func maxAreaV2(height []int) int {
	/*
		METHOD: Scanning
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		// Вычисляем текущую площадь
		currentArea := (right - left) * min(height[left], height[right])
		// Обновляем максимальную площадь, если текущая больше
		maxArea = max(maxArea, currentArea)

		// Перемещаем указатель, где высота меньше
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func main() {
//	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
//	fmt.Println(maxArea(height)) // Вывод: 49
//}

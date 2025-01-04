package main

func maxPowerV2(s string) int {
	/*
		METHOD: Sliding Window
		В текущем решении используется подход с двумя указателями (left и right), которые представляют границы текущего окна.
		Мы итерируемся по строке и обновляем границы окна на основе текущего символа.
		Если текущий символ отличается от предыдущего, мы обновляем левую границу окна.
		Затем мы вычисляем длину текущего окна (right - left + 1) и обновляем максимальную длину, если она больше предыдущей.
		Мы продолжаем итерироваться до конца строки.
		В конце мы возвращаем максимальную длину последовательных символов.

		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	left, right := 0, 0
	maxLen := 1

	for right < len(s) {
		if right > 0 && s[right] != s[right-1] {
			left = right
		}
		maxLen = max(maxLen, right-left+1)
		right++
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

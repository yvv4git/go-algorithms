package main

func reverseStrV2(s string, k int) string {
	/*
		METHOD: Iterative

		TIME COMPLEXITY: O(n), так как мы проходим по каждому символу в строке ровно один раз.

		SPACE COMPLEXITY: O(n), где n - количество символов в строке.
	*/
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n; i += 2 * k {
		for left, right := i, min(i+k-1, n-1); left < right; left, right = left+1, right-1 {
			runes[left], runes[right] = runes[right], runes[left]
		}
	}
	return string(runes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//func main() {
//	fmt.Println(reverseStrV2("abcdefg", 2)) // Вывод: "bacdfeg"
//}

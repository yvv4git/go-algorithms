package main

func repeatedNTimesV4(nums []int) int {
	/*
		METHOD: XOR
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	xor := 0
	for _, num := range nums {
		xor ^= num
	}

	for _, num := range nums {
		if xor^num == 0 {
			return num
		}
	}

	return -1
}

// func main() {
// 	nums := []int{1, 2, 3, 3}
// 	result := repeatedNTimesV4(nums)
// 	fmt.Println("Повторяющийся элемент:", result) // Вывод: Повторяющийся элемент: 3
// }

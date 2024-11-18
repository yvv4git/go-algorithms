package main

func repeatedNTimesV3(nums []int) int {
	/*
		METHOD: Two Pointers
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	for i := 0; i < len(nums); i++ {
		for j := 1; j <= 3 && i+j < len(nums); j++ {
			if nums[i] == nums[i+j] {
				return nums[i]
			}
		}
	}
	return -1
}

// func main() {
//     nums := []int{1, 2, 3, 3}
//     result := repeatedNTimesV3(nums)
//     fmt.Println("Повторяющийся элемент:", result) // Вывод: Повторяющийся элемент: 3
// }

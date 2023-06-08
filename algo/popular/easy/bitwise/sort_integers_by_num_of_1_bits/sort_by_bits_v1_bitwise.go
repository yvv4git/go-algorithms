package sort_integers_by_num_of_1_bits

import "sort"

const maskOneBits = 14

func sortByBitsV1(arr []int) []int {
	/*
		Method: Bitwise
		Time complexity: O(n * 14 + n)
		Space complexity: O1) - так как мы просто обновляем значения в массиве

		Explanation:
		Стандартная сортировка имеет сложность O(n * log(n)).
		- Плюс к этому есть ограничение 10^4 = 10_000.
		- Количество бит в числе максимум 2^14, так как 10_000 < 2^14. Так как 2^13=8_192, а 2^14=16_384, т.е. от 0-13 бит могут быть единицами.
	*/

	// O(n)
	for i := 0; i < len(arr); i++ {
		arr[i] += countOnes(arr[i]) << maskOneBits
	}

	sort.Ints(arr) // // n log(n)

	// O(n)
	for i := 0; i < len(arr); i++ {
		arr[i] &= (1 << maskOneBits) - 1 // (1 << 14) - 1          = 0000011111111111111
	}

	return arr
}

func countOnes(n int) int {
	c := 0
	for i := 0; i < maskOneBits; i++ {
		c += (n >> i) & 1
	}
	return c
}

//go:build ignore

package main

import "fmt"

// APPROACH: "Решение через битовые операции (Bit Manipulation)"
// - Для каждого бита (0..31) считаем, сколько раз он установлен во всех числах массива.
// - Если бит встречается более n/2 раз, он есть в majority element.
// - Собираем majority element по битам.
// - Гарантируется, что majority element существует.
//
// TIME COMPLEXITY: O(32n) ≈ O(n)
// SPACE COMPLEXITY: O(1)
func majorityElement(nums []int) int {
	n := len(nums)
	result := 0
	for i := 0; i < 32; i++ {
		bitCount := 0
		for _, num := range nums {
			if (num>>i)&1 == 1 {
				bitCount++
			}
		}
		if bitCount > n/2 {
			result |= (1 << i)
		}
	}
	return result
}

func main() {
	examples := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
	}
	for _, nums := range examples {
		res := majorityElement(nums)
		fmt.Printf("Input: %v\n", nums)
		fmt.Printf("Majority element: %d\n\n", res)
	}
}

package main

import "fmt"

func mergeTwoSortedList(num1, num2 []int) (result []int) {
	i, j := len(num1)-1, len(num2)-1
	k := i + j - 1
	fmt.Printf("%v %v \n", num1, num2)

	result = make([]int, k+1)
	for k > 0 {
		if num1[i] > num2[j] {
			//copy(result[0:1], num1[i])
			copy(result[0:1], num1[i-1:i])
			i--
		} else if num1[i] < num2[j] {
			copy(result[0:1], num1[i-1:i])
			j--
		}

		k--
	}

	fmt.Printf("\n%d %d \n", i, j)
	return
}

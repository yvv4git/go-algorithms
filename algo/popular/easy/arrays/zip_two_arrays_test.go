package arrays

import (
	"fmt"
	"testing"
)

// Задача: написать реализацию функции zip, чтобы получить слайс слайсов [[1 4] [2 5] [3 6]]
// Обрати внимание, что входящие массивы могут быть разной длины. Соответственно, результат надо генерить
// по длине меньшего из них.
func TestZip(t *testing.T) {
	s1, s2 := []int{1, 2, 3}, []int{4, 5, 6, 7, 8}
	fmt.Println(zip(s1, s2)) // [[1 4] [2 5] [3 6]]
}

func zip(s1 []int, s2 []int) [][]int {
	var result [][]int
	var tmp []int

	minLen := 0
	if len(s1) > len(s2) {
		minLen = len(s2)
	} else {
		minLen = len(s1)
	}

	for i := 0; i < minLen-1; i++ {
		tmp = nil
		tmp = append(tmp, s1[i])
		tmp = append(tmp, s2[i])
		result = append(result, tmp)
	}

	return result
}

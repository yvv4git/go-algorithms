package ex02

import (
	"fmt"
	"testing"
)

func Test_combinations(t *testing.T) {
	result := combinations([]int{1, 2, 3})
	for i, perm := range result {
		fmt.Printf("[%d] %#v \n", i, perm)
	}
}

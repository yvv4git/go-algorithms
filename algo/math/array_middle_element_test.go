package math

import "testing"

func middleIdx(left, right int) int {
	/*
		Уникальный способ, позволяющий предотвратить выход за пределы массива.
	*/
	return left + (right-left)/2
}

func TestMiddleElementOfArray001(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5}

	left := a[0]
	right := len(a) - 1
	middle := middleIdx(left, right) // Такая конструкция позволяет предотвратить выход за пределы массива

	t.Logf("middle=%d a[%d]=%d", middle, middle, a[middle])
}

func TestMiddleElementOfArray002(t *testing.T) {
	a := [...]int{1, 2, 3, 4, 5, 6}

	left := a[0]
	right := len(a) - 1
	middle := middleIdx(left, right) // Такая конструкция позволяет предотвратить выход за пределы массива

	t.Logf("middle=%d a[%d]=%d", middle, middle, a[middle])
}

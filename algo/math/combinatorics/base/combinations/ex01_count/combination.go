package ex01_count

import (
	"fmt"
)

// CombinationsCnt - Сочетания.
// Возможно не самый подходящий перевод combinations для сочитаний.
// Сочетания - это когда у нас есть набор из n элементов и мы
// будем выбирать из них(n) m различных объектов. Это похоже на размещение, но
// в отличие от размещения порядок не важен.
// Т.е. в размещении [1, 2] и [2, 1] это разные результирующие наборы.
// А в сочетании это одно и то же.
// В итоге количество сочетаний меньше, чем количество размещений.
// Формула: n!/(n-m)! * m!
func CombinationsCnt(n int, m int) int {
	return CalcByLoop(n) / (CalcByLoop(n-m) * CalcByLoop(m))
}

// Сгенерировать все перестановки среза.
// Так же символы могут повторяться, например: "aaa".
func TransposeSliceRunes(a []rune, width int) {
	var permutate func(prefix []rune)
	permutate = func(prefix []rune) {
		for i := 0; i < len(a); i++ {
			if len(prefix) >= width {
				fmt.Println(string(prefix))
				return
			}
			prefix := append(prefix, a[i])
			permutate(prefix)
		}
	}

	permutate([]rune{})
}

// CalcByLoop is used for calculate factorial without recursion.
func CalcByLoop(n int) int {
	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}

	return result
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	fmt.Println(countSequences(n))
}

// CountSequences возвращает количество последовательностей из нулей и единиц
// длины n, в которых нет трёх единиц подряд.
func countSequences(n int) int {
	/*
		METHOD: Dynamic Programming (с 3 состояниями)
		Суть задачи: Подсчёт бинарных последовательностей длины n, где нет трёх единиц подряд.
		dp[i][j] - количество последовательностей длины i, заканчивающихся на j единиц (j=0,1,2).

		TIME COMPLEXITY: O(n) - один проход, на каждом шаге O(1) операций
		SPACE COMPLEXITY: O(1) - храним только 3 значения (zeroEnd, oneEnd, twoEnd)
	*/

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2 // "0", "1"
	}
	if n == 2 {
		return 4 // "00", "01", "10", "11"
	}

	// dp[i][j] - количество последовательностей длины i,
	// заканчивающихся на j последовательных единиц (j = 0, 1, 2)
	// j=0 означает, что последний символ - 0
	// j=1 означает, что последний символ - 1 (одна единица в конце)
	// j=2 означает, что последние два символа - 11 (две единицы в конце)

	// Инициализация для i=2
	zeroEnd := 2 // "00", "01"
	oneEnd := 1  // "10"
	twoEnd := 1  // "11"

	for i := 3; i <= n; i++ {
		newZeroEnd := zeroEnd + oneEnd + twoEnd // добавляем 0 к любой последовательности
		newOneEnd := zeroEnd                    // добавляем 1 к последовательности, заканчивающейся на 0
		newTwoEnd := oneEnd                     // добавляем 1 к последовательности, заканчивающейся на 1

		zeroEnd, oneEnd, twoEnd = newZeroEnd, newOneEnd, newTwoEnd
	}

	return zeroEnd + oneEnd + twoEnd
}

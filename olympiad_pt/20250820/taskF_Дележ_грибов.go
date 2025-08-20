//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}

	// Исходная радость
	originalJoy := 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			originalJoy += a[i]
		} else {
			originalJoy -= a[i]
		}
	}

	maxJoy := originalJoy

	// Находим лучший обмен между четной и нечетной позицией
	// Для этого ищем минимальный элемент на четных позициях и максимальный на нечетных
	minEven := 1001
	maxOdd := -1

	for i := 0; i < n; i++ {
		if i%2 == 0 && a[i] < minEven {
			minEven = a[i]
		}
		if i%2 == 1 && a[i] > maxOdd {
			maxOdd = a[i]
		}
	}

	// Если можно улучшить обменом четного и нечетного
	if maxOdd > minEven {
		improvement := 2 * (maxOdd - minEven)
		if originalJoy+improvement > maxJoy {
			maxJoy = originalJoy + improvement
		}
	}

	fmt.Fprintln(out, maxJoy)
}

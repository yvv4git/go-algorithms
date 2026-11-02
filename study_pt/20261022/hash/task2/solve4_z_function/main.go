package main

import (
	"bufio"
	"fmt"
	"os"
)

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	left, right := 0, 0
	for i := 1; i < n; i++ {
		if i <= right {
			z[i] = min(right-i+1, z[i-left])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > right {
			left = i
			right = i + z[i] - 1
		}
	}
	return z
}

func minCyclicShiftV4(x, y string) int {
	/*
		TASK: Проверить, является ли y циклическим сдвигом x вправо,
		и найти минимальный размер сдвига или -1, если сдвига нет.
		Циклический сдвиг вправо на k позиций - это перенос последних k символов
		в начало строки, а остальные символы сдвигаются вправо.
		Строку можно представить как кольцо символов, которое поворачивают на k позиций.
		Примеры для x = "abcde" (n = 5):
		k = 0: "abcde" - сдвига нет;
		k = 1: "eabcd" - последняя буква "e" уходит в начало;
		k = 2: "deabc" - последние "de" уходят в начало;
		k = 3: "cdeab" - последние "cde" уходят в начало;
		k = 5: "abcde" - полный оборот, снова исходная строка.
		Если y = "abced", то это не циклический сдвиг, ответ -1.

		METHOD: Z-function (Z-функция)
		Z-функция строки - это массив Z, где Z[i] равен длине наибольшего
		общего префикса всей строки и её суффикса, начинающегося в позиции i.
		Проще говоря, Z[i] показывает, сколько первых символов строки совпадает
		с символами начиная с позиции i.
		Строим строку t = y + "#" + x + x и считаем её Z-функцию.
		Разделитель "#" не даёт совпадению выйти за границу y.
		Для позиции i внутри части x + x условие Z[i] >= n означает,
		что y встречается в x + x с позиции pos = i - (n + 1).
		Запоминаем максимальное pos не больше n, ответ k = n - pos.

		TIME COMPLEXITY: O(n)
		- Z-функция строки длины 3n + 1 считается за один проход.

		SPACE COMPLEXITY: O(n)
		- Строка t длины 3n + 1 и массив Z длины 3n + 1.
	*/
	n := len(x)
	if len(y) != n {
		return -1
	}
	if n == 0 {
		return 0
	}

	t := y + "#" + x + x
	z := zFunction(t)

	offset := n + 1
	maxPos := -1
	for i := offset; i < len(t); i++ {
		if z[i] >= n {
			pos := i - offset
			if pos <= n && pos > maxPos {
				maxPos = pos
			}
		}
	}

	if maxPos == -1 {
		return -1
	}
	return n - maxPos
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var x, y string
	if _, err := fmt.Fscan(in, &x, &y); err != nil {
		return
	}

	fmt.Fprintln(out, minCyclicShiftV4(x, y))
}

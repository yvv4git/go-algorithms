package main

import (
	"bufio"
	"fmt"
	"os"
)

func prefixFunction(s string) []int {
	pi := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := pi[i-1]
		for j > 0 && s[i] != s[j] {
			j = pi[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		pi[i] = j
	}
	return pi
}

func minCyclicShiftV3(x, y string) int {
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

		METHOD: KMP (алгоритм Кнута - Морриса - Пратта)
		Префикс-функция строки y - это массив pi, где pi[i] равен длине
		наибольшего собственного префикса y[0..i], который одновременно
		является суффиксом y[0..i].
		Проще говоря, pi[i] показывает, сколько первых символов строки можно
		оставить при несовпадении, чтобы не начинать сравнение заново.
		KMP ищет y в z = x + x за один проход: при несовпадении он откатывается
		по pi[matched-1] и продолжает с уже совпавшей части.
		Для вхождения с концом в позиции i начало равно pos = i - n + 1.
		Запоминаем максимальное pos не больше n, ответ k = n - pos.

		TIME COMPLEXITY: O(n)
		- Построение префикс-функции O(n).
		- Проход по z длины 2n с амортизированными откатами O(n).

		SPACE COMPLEXITY: O(n)
		- Префикс-функция длины n и строка z длины 2n.
	*/
	n := len(x)
	if len(y) != n {
		return -1
	}
	if n == 0 {
		return 0
	}

	pi := prefixFunction(y)
	z := x + x

	matched := 0
	maxPos := -1
	for i := 0; i < len(z); i++ {
		for matched > 0 && z[i] != y[matched] {
			matched = pi[matched-1]
		}
		if z[i] == y[matched] {
			matched++
		}
		if matched == n {
			pos := i - n + 1
			if pos <= n && pos > maxPos {
				maxPos = pos
			}
			matched = pi[matched-1]
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

	fmt.Fprintln(out, minCyclicShiftV3(x, y))
}

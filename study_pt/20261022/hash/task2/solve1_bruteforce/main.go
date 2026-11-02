package main

import (
	"bufio"
	"fmt"
	"os"
)

func minCyclicShift(x, y string) int {
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

		METHOD: Bruteforce (полный перебор сдвигов)
		Полный перебор - это проверка всех возможных вариантов по очереди.
		Перебираем k от 0 до n-1 и сравниваем y со сдвигом x на k позиций.
		Сравнение идёт посимвольно без создания новой строки:
		символ y[i] должен совпасть с x[(n-k+i) mod n].
		Первый подходящий k и есть минимальный сдвиг.

		TIME COMPLEXITY: O(n^2)
		- n возможных сдвигов, каждый сравнение занимает до O(n).

		SPACE COMPLEXITY: O(1)
		- Дополнительные строки не создаются, только счётчики и флаги.
	*/
	n := len(x)
	if len(y) != n {
		return -1
	}
	if n == 0 {
		return 0
	}

	for k := range n {
		matches := true
		for i := range n {
			if y[i] != x[(n-k+i)%n] {
				matches = false
				break
			}
		}
		if matches {
			return k
		}
	}
	return -1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var x, y string
	if _, err := fmt.Fscan(in, &x, &y); err != nil {
		return
	}

	fmt.Fprintln(out, minCyclicShift(x, y))
}

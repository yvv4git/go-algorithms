package main

import (
	"bufio"
	"fmt"
	"os"
)

type Transaction struct {
	ID int
	RS []int
	WS []int
}

func selectEmpty(_ []Transaction) []int {
	/*
		TASK: Выбрать последовательность транзакций без конфликтов как можно большей длины.
		Транзакция - это операция, которая читает локации памяти RS_i и записывает локации памяти WS_i.
		Конфликт - это ситуация, когда более ранняя транзакция записывает в локацию, которую более поздняя транзакция читает.
		Нужно вернуть номера транзакций в порядке выполнения так, чтобы таких конфликтов не было.
		Пример: tx1 читает {1} и пишет {2}, tx2 читает {2} и пишет {3}, tx3 читает {3} и пишет {1}.
		Последовательность (1, 3) безопасна, потому что tx1 пишет в 2, а tx3 читает 3.
		Последовательность (1, 2) небезопасна, потому что tx1 пишет в 2, а tx2 читает 2.
		Оптимальное решение не требуется, любой корректный ответ допустим.
		Если вернуть пустую последовательность, конфликтов нет, но качество ответа нулевое.

		METHOD: Empty sequence (пустая последовательность)
		Пустая последовательность не содержит ни одной транзакции, значит в ней нет пар транзакций.
		Если пар нет, то не может быть и пары, где ранняя транзакция пишет, а поздняя читает ту же локацию.
		Поэтому ответ всегда корректен, но его длина равна 0.

		TIME COMPLEXITY: O(1)
		- Результат не зависит от входа, создаётся пустой список.

		SPACE COMPLEXITY: O(1)
		- Дополнительная память не используется, кроме пустого результата.
	*/
	return []int{}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return
	}

	transactions := make([]Transaction, 0, n)
	for i := 1; i <= n; i++ {
		var r, w int
		if _, err := fmt.Fscan(in, &r, &w); err != nil {
			break
		}
		tx := Transaction{ID: i, RS: make([]int, r), WS: make([]int, w)}
		for j := 0; j < r; j++ {
			fmt.Fscan(in, &tx.RS[j])
		}
		for j := 0; j < w; j++ {
			fmt.Fscan(in, &tx.WS[j])
		}
		transactions = append(transactions, tx)
	}

	ans := selectEmpty(transactions)
	fmt.Fprintln(out, len(ans))
}

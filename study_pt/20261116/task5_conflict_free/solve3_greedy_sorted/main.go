package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Transaction struct {
	ID int
	RS []int
	WS []int
}

func selectGreedySorted(transactions []Transaction, memoryLocations int) []int {
	/*
		TASK: Выбрать последовательность транзакций без конфликтов как можно большей длины.
		Транзакция - это операция, которая читает локации памяти RS_i и записывает локации памяти WS_i.
		Конфликт - это ситуация, когда более ранняя транзакция записывает в локацию, которую более поздняя транзакция читает.
		Нужно вернуть номера транзакций в порядке выполнения так, чтобы таких конфликтов не было.
		Пример: tx1 читает {1} и пишет {2}, tx2 читает {2} и пишет {3}, tx3 читает {3} и пишет {1}.
		Последовательность (1, 3) безопасна, потому что tx1 пишет в 2, а tx3 читает 3.
		Последовательность (1, 2) небезопасна, потому что tx1 пишет в 2, а tx2 читает 2.
		Если транзакция только читает уже записанную локацию позже, это конфликт, а запись после чтения конфликтом не считается.
		Оптимальное решение не требуется, любой корректный ответ допустим.

		METHOD: Greedy sorted append (жадное добавление с сортировкой)
		Жадный алгоритм строит ответ шаг за шагом и на каждом шаге делает безопасный выбор.
		Сначала сортируем транзакции по эвристике: меньше размер RS, при равенстве меньше размер WS, затем меньше номер.
		Транзакция с маленьким RS реже читает из уже записанных локаций, а маленький WS меньше мешает будущим транзакциям.
		После сортировки используем тот же безопасный проход: пропускаем транзакцию, если её RS пересекается с уже записанными локациями.
		Если пересечений нет, добавляем номер транзакции в ответ и помечаем все локации из WS.
		Эвристика может улучшить длину ответа, но корректность обеспечивается проверкой конфликтов.

		TIME COMPLEXITY: O(S + n log n)
		- S - суммарный размер всех множеств RS и WS, сортировка номеров транзакций занимает O(n log n).

		SPACE COMPLEXITY: O(m + n)
		- Нужен массив written длины m+1, копия списка транзакций и список выбранных номеров.
	*/
	items := append([]Transaction(nil), transactions...)
	sort.Slice(items, func(i, j int) bool {
		if len(items[i].RS) != len(items[j].RS) {
			return len(items[i].RS) < len(items[j].RS)
		}
		if len(items[i].WS) != len(items[j].WS) {
			return len(items[i].WS) < len(items[j].WS)
		}
		return items[i].ID < items[j].ID
	})

	written := make([]bool, memoryLocations+1)
	order := make([]int, 0, len(items))

	for _, tx := range items {
		conflict := false
		for _, loc := range tx.RS {
			if loc >= 0 && loc < len(written) && written[loc] {
				conflict = true
				break
			}
		}
		if conflict {
			continue
		}
		order = append(order, tx.ID)
		for _, loc := range tx.WS {
			if loc >= 0 && loc < len(written) {
				written[loc] = true
			}
		}
	}

	return order
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

	ans := selectGreedySorted(transactions, m)
	fmt.Fprintln(out, len(ans))
	for i, id := range ans {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, id)
	}
	if len(ans) > 0 {
		fmt.Fprintln(out)
	}
}

package _073_time_needed_to_buy_tickets

// Функция timeRequiredToBuy возвращает минимальное количество времени, необходимое для того, чтобы все люди купили билеты.
// tickets - массив, где tickets[i] - количество билетов, которое нужно купить человеку i.
// k - индекс человека, который стоит в начале очереди.
func timeRequiredToBuyV2(tickets []int, k int) (res int) {
	/*
		METHOD: Loop
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)

		Хороший метод решения задачи за O(n).
		Смысл в том, что не обязательно проходить несколько итераций по списку билетов(tickets), а можно
		использовать закономерность.
		При индексах до k, можно сразу вычесть минимальное значение, так как столько раз придется пройтись по кругу
		снимая по 1 из tickets.
	*/

	// Проходим по всем билетам в очереди
	for i, v := range tickets {
		if i <= k {
			res += min(v, tickets[k])
		} else {
			res += min(v, tickets[k]-1)
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

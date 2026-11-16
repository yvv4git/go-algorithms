package main

import (
	"bufio"
	"fmt"
	"os"
)

func groupCost(freq []int64, start, count int) int64 {
	var cost int64
	for i := range count {
		cost += int64(i+1) * freq[start+i]
	}
	return cost
}

func searchBest(freq []int64, index, remainingButtons, start int, current, best []int, currentCost, bestCost int64) int64 {
	if index == len(current)-1 {
		copy(current[index:], []int{remainingButtons})
		total := currentCost + groupCost(freq, start, remainingButtons)
		if bestCost == -1 || total < bestCost {
			copy(best, current)
			return total
		}
		return bestCost
	}

	for count := 0; count <= remainingButtons; count++ {
		current[index] = count
		cost := groupCost(freq, start, count)
		bestCost = searchBest(freq, index+1, remainingButtons-count, start+count, current, best, currentCost+cost, bestCost)
	}
	return bestCost
}

func distributeEmojis(freq []int64, buttons int) []int {
	/*
		TASK: Разложить N эмодзи по M кнопкам так, чтобы суммарное число нажатий было минимальным.
		Эмодзи идут в фиксированном порядке, переставлять их нельзя.
		На одной кнопке эмодзи лежат подряд: первый выбирается одним нажатием, второй - двумя, третий - тремя и так далее.
		Если на кнопку положили k эмодзи с частотами f1, f2, ..., fk, вклад кнопки равен f1 + 2*f2 + ... + k*fk.
		Нужно вернуть массив b длины M, где b_i - число эмодзи на i-й кнопке, и сумма b_i равна N.
		Пример: N = 5, M = 3, частоты [1, 3, 5, 1, 1].
		Разбиение (1, 1, 3) даёт 1*1 + 3*1 + 5*1 + 1*2 + 1*3 = 14 нажатий.
		Разбиение (1, 2, 2) даёт 1*1 + 3*1 + 5*2 + 1*1 + 1*2 = 17 нажатий, это хуже.
		Если N = 0, все кнопки пустые и ответ состоит из нулей.
		Если M = 1, все N эмодзи лежат на одной кнопке.
		Если M >= N, можно дать каждой кнопке не больше одного эмодзи, и каждый частый эмодзи выберется одним нажатием.

		METHOD: Bruteforce (полный перебор разбиений)
		Полный перебор - это проверка всех возможных вариантов по очереди.
		Разбиение определяется размерами групп b1, ..., bM, где каждое b_i >= 0 и b1 + ... + bM = N.
		Рекурсивно выбираем размер каждой кнопки слева направо, пока не распределим все эмодзи.
		Для каждого полного варианта считаем суммарную стоимость и запоминаем вариант с минимальной стоимостью.
		При равной стоимости можно вернуть любой вариант, поэтому сохраняем первый найденный минимум.

		TIME COMPLEXITY: O(C(N+M-1, M-1) * N)
		- Число разбиений N эмодзи на M групп равно C(N+M-1, M-1), для каждого считаем стоимость за O(N) суммарно.

		SPACE COMPLEXITY: O(N + M)
		- Храним текущий и лучший варианты длины M и используем рекурсию глубины M.
	*/
	n := len(freq)
	best := make([]int, buttons)
	current := make([]int, buttons)
	if buttons == 0 {
		return best
	}
	searchBest(freq, 0, n, 0, current, best, 0, -1)
	return best
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	if _, err := fmt.Fscan(in, &n, &m); err != nil {
		return
	}

	freq := make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &freq[i])
	}

	ans := distributeEmojis(freq, m)
	for i, v := range ans {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}

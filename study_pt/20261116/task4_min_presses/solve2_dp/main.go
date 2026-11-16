package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func distributeEmojisDP(freq []int64, buttons int) []int {
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

		METHOD: Dynamic programming (динамическое программирование)
		Динамическое программирование - это способ решать задачу, запоминая ответы для меньших частей.
		Пусть dp[j][i] - минимальная стоимость, если первые i эмодзи распределены по первым j кнопкам.
		Последняя кнопка может содержать хвост из k эмодзи, поэтому dp[j][i] = min(dp[j-1][i-k] + cost(i-k+1, i)).
		Здесь cost(l, r) = a_l + 2*a_{l+1} + ... + (r-l+1)*a_r - стоимость одной кнопки с этим хвостом.
		База: dp[0][0] = 0, а невозможные состояния имеют бесконечную стоимость.
		В prev[j][i] храним k - сколько эмодзи положили на j-ю кнопку, чтобы потом восстановить ответ.

		TIME COMPLEXITY: O(N^2 * M)
		- Для каждого состояния (j, i) перебираем длину последней группы k от 0 до i.

		SPACE COMPLEXITY: O(N * M)
		- Храним таблицы dp и prev размеров (M+1) x (N+1).
	*/
	n := len(freq)
	const inf int64 = math.MaxInt64 / 4

	dp := make([][]int64, buttons+1)
	prev := make([][]int, buttons+1)
	for j := 0; j <= buttons; j++ {
		dp[j] = make([]int64, n+1)
		prev[j] = make([]int, n+1)
		for i := 0; i <= n; i++ {
			dp[j][i] = inf
		}
	}
	dp[0][0] = 0

	for j := 1; j <= buttons; j++ {
		for i := 0; i <= n; i++ {
			var tailCost int64
			var tailSum int64
			for k := 0; k <= i; k++ {
				if k > 0 {
					tailSum += freq[i-k]
					tailCost += tailSum
				}
				prevCost := dp[j-1][i-k]
				if prevCost == inf {
					continue
				}
				candidate := prevCost + tailCost
				if candidate < dp[j][i] {
					dp[j][i] = candidate
					prev[j][i] = k
				}
			}
		}
	}

	ans := make([]int, buttons)
	i := n
	for j := buttons; j >= 1; j-- {
		k := prev[j][i]
		ans[j-1] = k
		i -= k
	}
	return ans
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

	ans := distributeEmojisDP(freq, m)
	for i, v := range ans {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}

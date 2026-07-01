// B. Цензура
// Подход: Строим Ахо-Корасик над запрещёнными строками, достраиваем
// суффиксные ссылки и полную таблицу переходов.  Вершины, в которых
// заканчивается запрещённая строка (или достижимые по fail из такой),
// помечаются как плохие.  DP[pos][v] — число строк длины pos, после
// прочтения которых автомат находится в вершине v и ни одна запрещённая
// подстрока не встречена.  Ответ — сумма DP[m][v] по всем неплохим v.
// Time complexity:  O(p*L + K*m) — построение автомата (p <= 10, L <= 10)
//                   и DP по m <= 100 шагам, каждый по K <= 101 состоянию.
// Space complexity: O(K*n + K*m) — автомат (K <= 101 вершин, n <= 100 букв)
//                   и таблица DP.

package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, p int
	fmt.Fscan(in, &n, &m, &p)

	var alpha string
	fmt.Fscan(in, &alpha)

	// Отображаем символ алфавита в индекс 0..n-1
	idx := make(map[byte]int)
	for i := 0; i < n; i++ {
		idx[alpha[i]] = i
	}

	// Вершина автомата Ахо-Корасик
	type node struct {
		next []int // переходы по каждому символу алфавита (полные)
		fail int   // суффиксная ссылка
		bad  bool  // true, если в этой вершине заканчивается запрещённая строка
	}

	trie := []node{{next: make([]int, n)}}

	// Вставка запрещённой строки в бор
	add := func(s string) {
		v := 0
		for i := 0; i < len(s); i++ {
			c := idx[s[i]]
			if trie[v].next[c] == 0 {
				trie[v].next[c] = len(trie)
				trie = append(trie, node{next: make([]int, n)})
			}
			v = trie[v].next[c]
		}
		trie[v].bad = true
	}

	for i := 0; i < p; i++ {
		var s string
		fmt.Fscan(in, &s)
		add(s)
	}

	// BFS — строим суффиксные ссылки и заполняем недостающие переходы
	q := make([]int, 0)
	for c := 0; c < n; c++ {
		if trie[0].next[c] != 0 {
			child := trie[0].next[c]
			trie[child].fail = 0
			q = append(q, child)
		} else {
			trie[0].next[c] = 0 // корень переходит сам в себя
		}
	}

	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for c := 0; c < n; c++ {
			u := trie[v].next[c]
			if u != 0 {
				// У вершины есть прямой потомок по c
				f := trie[v].fail
				trie[u].fail = trie[f].next[c]
				if trie[trie[u].fail].bad {
					trie[u].bad = true
				}
				q = append(q, u)
			} else {
				// Прямого потомка нет — переход через fail
				trie[v].next[c] = trie[trie[v].fail].next[c]
			}
		}
	}

	// DP по длине строки и состоянию автомата
	sz := len(trie)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, sz)
	}
	dp[0][0] = 1 // пустая строка, корень

	for pos := 0; pos < m; pos++ {
		for v := range sz {
			if trie[v].bad || dp[pos][v] == 0 {
				continue
			}
			for c := 0; c < n; c++ {
				nv := trie[v].next[c]
				if trie[nv].bad {
					continue
				}
				dp[pos+1][nv] = (dp[pos+1][nv] + dp[pos][v]) % MOD
			}
		}
	}

	ans := 0
	for v := range sz {
		if !trie[v].bad {
			ans = (ans + dp[m][v]) % MOD
		}
	}
	fmt.Fprintln(out, ans)
}

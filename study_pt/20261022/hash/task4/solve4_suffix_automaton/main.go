package main

import (
	"bufio"
	"fmt"
	"os"
)

type samState struct {
	next   [26]int
	link   int
	length int
}

func newSAMState(length int) samState {
	state := samState{link: -1, length: length}
	for i := range state.next {
		state.next[i] = -1
	}
	return state
}

func countDistinctSubstringsV4(s string) int {
	/*
		TASK: Посчитать количество различных подстрок строки S.
		Подстрока S[i..j] - это непрерывный кусок строки с i-го по j-й символ включительно.
		Одинаковые куски считаются один раз.
		Пример: S = "abc" - подстроки "a", "b", "c", "ab", "bc", "abc", ответ 6.
		Пример: S = "aaa" - различные подстроки "a", "aa", "aaa", ответ 3.
		Пример: S = "abab" - различные подстроки "a", "b", "ab", "ba", "aba", "bab", "abab", ответ 7.

		METHOD: Suffix automaton (суффиксный автомат)
		Суффиксный автомат - это минимальный автомат, который принимает все подстроки строки.
		У каждого состояния v есть len[v] (длина самых длинных строк) и link[v] (суффиксная ссылка).
		Состояние v отвечает за группу подстрок длины от len[link[v]] + 1 до len[v].
		Поэтому вклад состояния равен len[v] - len[link[v]] различных подстрок.
		Сумма вкладов всех состояний, кроме начального, и есть ответ.

		TIME COMPLEXITY: O(n)
		- Каждое состояние и переход добавляются амортизированно за O(1).

		SPACE COMPLEXITY: O(n)
		- Число состояний суффиксного автомата не больше 2n.
	*/
	states := make([]samState, 0, 2*len(s)+1)
	states = append(states, newSAMState(0))
	last := 0

	for i := 0; i < len(s); i++ {
		c := int(s[i] - 'a')
		current := len(states)
		states = append(states, newSAMState(states[last].length+1))

		p := last
		for p != -1 && states[p].next[c] == -1 {
			states[p].next[c] = current
			p = states[p].link
		}

		if p == -1 {
			states[current].link = 0
		} else {
			q := states[p].next[c]
			if states[p].length+1 == states[q].length {
				states[current].link = q
			} else {
				clone := len(states)
				cloneState := newSAMState(states[p].length + 1)
				cloneState.next = states[q].next
				cloneState.link = states[q].link
				states = append(states, cloneState)

				for p != -1 && states[p].next[c] == q {
					states[p].next[c] = clone
					p = states[p].link
				}
				states[q].link = clone
				states[current].link = clone
			}
		}

		last = current
	}

	distinct := 0
	for v := 1; v < len(states); v++ {
		distinct += states[v].length - states[states[v].link].length
	}
	return distinct
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var s string
	if _, err := fmt.Fscan(in, &s); err != nil {
		return
	}

	fmt.Fprintln(out, countDistinctSubstringsV4(s))
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type asker func(x, y int) bool

func pFromOrder(order []int, n int) []int {
	p := make([]int, n)
	for value, pos := range order {
		p[pos-1] = value + 1
	}
	return p
}

func parsePI(s string, n int) []int {
	pi := make([]int, 0, n)
	cur := 0
	has := false
	for _, c := range s {
		if c >= '0' && c <= '9' {
			cur = cur*10 + int(c-'0')
			has = true
		} else if has {
			pi = append(pi, cur)
			cur = 0
			has = false
		}
	}
	if has {
		pi = append(pi, cur)
	}
	return pi
}

func countComp(m int) int {
	if m <= 1 {
		return 0
	}
	return countComp(m/2) + countComp((m+1)/2) + m - 1
}

// recoverPermutationV3 восстанавливает перестановку, укладываясь в Q запросов.
func recoverPermutationV3(n, k, q int, ask asker) []int {
	maxQueries := q
	if maxQueries < 0 {
		maxQueries = 0
	}
	queriesUsed := 0

	cMax := countComp(n)
	if cMax == 0 {
		p := make([]int, n)
		for i := range p {
			p[i] = i + 1
		}
		return p
	}

	// Гарантированный путь: худший случай solve2 — (k+1)*C + k.
	if (k+1)*cMax+k <= maxQueries {
		return recoverPermutationAdaptive(n, k, q, ask)
	}

	override := map[[2]int]bool{}

	askFresh := func(x, y int) bool {
		if queriesUsed >= maxQueries {
			return x < y
		}
		queriesUsed++
		return ask(x, y)
	}

	var mergeSort func(arr []int, cmp func(int, int) bool) []int
	mergeSort = func(arr []int, cmp func(int, int) bool) []int {
		if len(arr) <= 1 {
			r := make([]int, len(arr))
			copy(r, arr)
			return r
		}
		mid := len(arr) / 2
		left := mergeSort(arr[:mid], cmp)
		right := mergeSort(arr[mid:], cmp)
		res := make([]int, 0, len(arr))
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if cmp(left[i], right[j]) {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		}
		res = append(res, left[i:]...)
		res = append(res, right[j:]...)
		return res
	}

	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = i + 1
	}

	// Этап 1: majority-сортировка. Бюджет делим так, чтобы хватило
	// на проверку кэша (~cMax) и остаток на всплытие.
	repeats := 1
	if 3*cMax <= maxQueries/2 {
		repeats = 3
	} else if 2*cMax <= maxQueries/2 {
		repeats = 2
	}
	askMaj := func(x, y int) bool {
		yes, no := 0, 0
		res := x < y
		for i := 0; i < repeats; i++ {
			if queriesUsed >= maxQueries {
				break
			}
			queriesUsed++
			if ask(x, y) {
				yes++
			} else {
				no++
			}
			rem := repeats - i - 1
			if yes > no+rem {
				res = true
				break
			}
			if no > yes+rem {
				res = false
				break
			}
			if i == repeats-1 {
				if yes > no {
					res = true
				} else if no > yes {
					res = false
				}
			}
		}
		override[[2]int{x, y}] = res
		return res
	}
	order = mergeSort(order, askMaj)
	if os.Getenv("DEBUG_V3") != "" && os.Getenv("PI") != "" {
		pi := parsePI(os.Getenv("PI"), n)
		if len(pi) == n {
			inv := 0
			for i := 0; i+1 < n; i++ {
				if pi[order[i]-1] > pi[order[i+1]-1] {
					inv++
				}
			}
			fmt.Fprintf(os.Stderr, "stage1 queries=%d adjInv=%d\n", queriesUsed, inv)
		}
	}

	// Этапы 2–3 (проверка кэша + пересортировка) — только если бюджет
	// позволяет. При тесном бюджете (cMax > Q/2) сразу чиним порядок
	// бинарными вставками: там обычно 1–k ошибок, это O(k log n).
	if cMax <= maxQueries/2 {
		bubbleReserve := 2 * (n - 1)
		if bubbleReserve > maxQueries/4 {
			bubbleReserve = maxQueries / 4
		}
		for key, val := range override {
			if queriesUsed+bubbleReserve >= maxQueries {
				break
			}
			x, y := key[0], key[1]
			fresh := askFresh(x, y)
			if fresh != val {
				override[key] = fresh
			}
		}

		pos := make([]int, n+1)
		for i, v := range order {
			pos[v] = i
		}
		answerSort := func(x, y int) bool {
			if v, ok := override[[2]int{x, y}]; ok {
				return v
			}
			if v, ok := override[[2]int{y, x}]; ok {
				return !v
			}
			if queriesUsed < maxQueries {
				v := askFresh(x, y)
				override[[2]int{x, y}] = v
				return v
			}
			return pos[x] < pos[y]
		}
		order = mergeSort(order, answerSort)
	}

	// Этап 4: ремонт инверсий. При тесном бюджете чистое всплытие не успевает
	// (элемент может быть далеко не на месте). При нахождении инверсии
	// вынимаем элемент и вставляем бинарным поиском — O(log n) вместо O(n).
	minClean := 1
	if k > 0 {
		minClean = k + 1
		if maxPasses := (maxQueries - queriesUsed) / max(n-1, 1); minClean > maxPasses && maxPasses > 0 {
			minClean = maxPasses
		}
		// При очень тесном бюджете один чистый проход лучше, чем ноль.
		if minClean < 1 {
			minClean = 1
		}
	}
	cleanStreak := 0
	for pass := 0; ; pass++ {
		changed := false
		for i := 0; i+1 < n; {
			if queriesUsed >= maxQueries {
				break
			}
			a, b := order[i], order[i+1]
			if askFresh(a, b) {
				i++
				continue
			}
			// a > b: b стоит правее, чем нужно. Ищем позицию для b
			// бинарным поиском по order[0..i] и вставляем туда (ротация).
			// Если бюджет кончился до конца поиска — НЕ двигаем массив
			// (неполный lo ломает порядок), просто выходим из ремонта.
			lo, hi := 0, i+1
			complete := true
			for lo < hi {
				if queriesUsed >= maxQueries {
					complete = false
					break
				}
				mid := (lo + hi) / 2
				if askFresh(order[mid], b) {
					lo = mid + 1
				} else {
					hi = mid
				}
			}
			if !complete {
				return pFromOrder(order, n)
			}
			// сдвигаем order[lo..i] вправо, b -> order[lo]
			for j := i + 1; j > lo; j-- {
				order[j] = order[j-1]
			}
			order[lo] = b
			changed = true
			if lo > 0 {
				i = lo - 1
			} else {
				i = 0
			}
		}
		if !changed {
			cleanStreak++
			if cleanStreak >= minClean {
				break
			}
		} else {
			cleanStreak = 0
		}
		if queriesUsed >= maxQueries {
			break
		}
		if pass >= 4*k+n+100 {
			break
		}
	}

	return pFromOrder(order, n)
}

// recoverPermutationAdaptive — гарантированная стратегия, худший случай (k+1)*C+k.
func recoverPermutationAdaptive(n, k, q int, ask asker) []int {
	liesLeft := k
	if liesLeft < 0 {
		liesLeft = 0
	}
	queriesUsed := 0
	maxQueries := q
	if maxQueries < 0 {
		maxQueries = 0
	}

	less := func(x, y int) bool {
		yes := 0
		no := 0
		for {
			if yes > liesLeft {
				liesLeft -= no
				return true
			}
			if no > liesLeft {
				liesLeft -= yes
				return false
			}
			if queriesUsed >= maxQueries {
				if yes == no {
					return x < y
				}
				return yes > no
			}
			queriesUsed++
			if ask(x, y) {
				yes++
			} else {
				no++
			}
		}
	}

	order := make([]int, n)
	for i := 0; i < n; i++ {
		order[i] = i + 1
	}

	var mergeSort func([]int) []int
	mergeSort = func(arr []int) []int {
		if len(arr) <= 1 {
			r := make([]int, len(arr))
			copy(r, arr)
			return r
		}
		mid := len(arr) / 2
		left := mergeSort(arr[:mid])
		right := mergeSort(arr[mid:])
		res := make([]int, 0, len(arr))
		i, j := 0, 0
		for i < len(left) && j < len(right) {
			if less(left[i], right[j]) {
				res = append(res, left[i])
				i++
			} else {
				res = append(res, right[j])
				j++
			}
		}
		res = append(res, left[i:]...)
		res = append(res, right[j:]...)
		return res
	}

	sortedPos := mergeSort(order)
	p := make([]int, n)
	for value, pos := range sortedPos {
		p[pos-1] = value + 1
	}
	return p
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k, q int
	if _, err := fmt.Fscan(reader, &n, &k, &q); err != nil {
		return
	}

	ask := func(x, y int) bool {
		fmt.Fprintf(writer, "? %d %d\n", x, y)
		writer.Flush()
		var resp string
		if _, err := fmt.Fscan(reader, &resp); err != nil {
			os.Exit(0)
		}
		if resp == "-1" {
			os.Exit(0)
		}
		return resp == "YES"
	}

	p := recoverPermutationV3(n, k, q, ask)
	fmt.Fprint(writer, "!")
	for _, v := range p {
		fmt.Fprintf(writer, " %d", v)
	}
	fmt.Fprintln(writer)
}

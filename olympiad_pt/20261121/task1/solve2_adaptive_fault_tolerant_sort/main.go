package main

import (
	"bufio"
	"fmt"
	"os"
)

type asker func(x, y int) bool

func recoverPermutationV2(n, k, q int, ask asker) []int {
	liesLeft := max(k, 0)

	less := func(x, y int) bool {
		yes := 0
		no := 0
		for {
			if yes-no > liesLeft {
				liesLeft -= no
				return true
			}
			if no-yes > liesLeft {
				liesLeft -= yes
				return false
			}
			if ask(x, y) {
				yes++
			} else {
				no++
			}
		}
	}

	order := make([]int, n)
	for i := range n {
		order[i] = i + 1
	}

	var mergeSort func([]int) []int
	mergeSort = func(arr []int) []int {
		if len(arr) <= 1 {
			res := make([]int, len(arr))
			copy(res, arr)
			return res
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
	_ = q

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

	p := recoverPermutationV2(n, k, q, ask)
	fmt.Fprint(writer, "!")
	for _, v := range p {
		fmt.Fprintf(writer, " %d", v)
	}
	fmt.Fprintln(writer)
}

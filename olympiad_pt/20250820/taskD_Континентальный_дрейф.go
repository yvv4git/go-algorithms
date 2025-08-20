//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
)

const INF = 1 << 30

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var nf, mf int
	fmt.Fscan(in, &nf, &mf)
	floating := make([]string, nf)
	for i := 0; i < nf; i++ {
		fmt.Fscan(in, &floating[i])
	}

	var no, mo int
	fmt.Fscan(in, &no, &mo)
	ground := make([]string, no)
	for i := 0; i < no; i++ {
		fmt.Fscan(in, &ground[i])
	}

	botF := make([]int, mf)
	for j := 0; j < mf; j++ {
		botF[j] = -1
		for i := nf - 1; i >= 0; i-- {
			if floating[i][j] == '*' {
				botF[j] = i
				break
			}
		}
	}

	topO := make([]int, mo)
	for j := 0; j < mo; j++ {
		topO[j] = INF
		for i := 0; i < no; i++ {
			if ground[i][j] == '#' {
				topO[j] = i
				break
			}
		}
	}

	bestAns := 0
	for shift := -mf; shift <= mo; shift++ {
		d := INF
		for j := 0; j < mf; j++ {
			k := j + shift
			if k < 0 || k >= mo {
				continue
			}
			if botF[j] == -1 || topO[k] == INF {
				continue
			}
			candidate := topO[k] - botF[j] - 1
			if candidate < d {
				d = candidate
			}
		}
		if d == INF {
			continue
		}

		count := 0
		for j := 1; j < mf; j++ {
			kLeft := (j - 1) + shift
			if kLeft < 0 || kLeft >= mo {
				continue
			}
			if botF[j] == -1 || topO[kLeft] == INF {
				continue
			}
			if topO[kLeft] == botF[j]+d {
				count++
			}
		}
		if count > bestAns {
			bestAns = count
		}
	}

	fmt.Fprintln(out, bestAns)
}

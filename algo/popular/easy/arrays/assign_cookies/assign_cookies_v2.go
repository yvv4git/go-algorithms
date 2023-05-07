package assign_cookies

import "sort"

func findContentChildrenV2(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i := 0

	for j := 0; j < len(s) && i < len(g); j++ {
		if g[i] <= s[j] {
			i++
		}
	}

	return i
}

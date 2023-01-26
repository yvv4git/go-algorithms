package main

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := map[int]int{}
	for _, item := range items1 {
		m[item[0]] += item[1]
	}
	for _, item := range items2 {
		m[item[0]] += item[1]
	}

	var ret [][]int
	for key, element := range m {
		ret = append(ret, []int{key, element})
	}

	sort.Slice(ret, func(i, j int) bool { return ret[i][0] < ret[j][0] })

	return ret
}

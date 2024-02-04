package top_k_frequently_elements

import (
	"sort"
)

// Задача от Ozon.
// Дан массив чисел arr и число k.
// Необходимо вернуть k самых часто встречающихся элементов массива.

// Input:  arr=[1,1,1,1,2,3,3,3,4,4,1], k=3
// Output: res=[1,3,4]

type KV struct {
	k int
	v int
}

func topKFrequentV1(a []int, k int) []int {
	var result []int
	dict := make(map[int]int)

	// TIME COMPLEXITY: O(n)
	// Space complexity: O(n)
	for _, val := range a {
		dict[val]++
	}

	// TIME COMPLEXITY: O(n)
	// Space complexity: O(n)
	buf := make([]KV, 0, len(dict))
	for k, v := range dict {
		buf = append(buf, KV{
			k: k,
			v: v,
		})
	}

	// TIME COMPLEXITY: O(n log(n)), because QuickSort
	// Space complexity: O(1)
	sort.Slice(buf, func(i, j int) bool {
		return buf[i].v > buf[j].v
	})

	for i := 0; i < k; i++ {
		result = append(result, buf[i].k)
	}

	return result
}

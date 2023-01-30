package main

import (
	"sort"
)

type Word string
type Words []string

func Anagrams(word string, words []string) []string {
	sorted := Word(word).Sorted()

	areAnagrams := func(b Word) bool {
		return sorted == b.Sorted()
	}

	return Words(words).Filter(areAnagrams)
}

func (ctx Word) Sorted() Word {
	sorted := []rune(ctx)

	sort.SliceStable(sorted, func(i, j int) bool {
		a := sorted[i]
		b := sorted[j]

		return a < b
	})

	return Word(sorted)
}

func (ctx Words) Filter(predicate func(x Word) bool) Words {
	var filtered Words

	for _, x := range ctx {
		if predicate(Word(x)) {
			filtered = append(filtered, x)
		}
	}

	return filtered
}

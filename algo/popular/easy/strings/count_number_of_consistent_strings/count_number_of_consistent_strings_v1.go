package count_number_of_consistent_strings

import "fmt"

func countConsistentStringsV1(allowed string, words []string) int {
	/*
		METHOD: Hash
		TIME COMPLEXITY: O(n log(n))
		Space complexity: O(n log(n))
	*/
	Hashmap := make(map[byte]int)
	for i := 0; i < len(allowed); i++ {
		Hashmap[allowed[i]]++
	}

	fmt.Println(Hashmap)
	count := 0
	for i := 0; i < len(words); i++ {
		a := words[i]
		j := 0
		b := true
		for j < len(a) {
			_, ok := Hashmap[a[j]]
			if ok {
				b = true
			} else {
				b = false
				break
			}
			j++
		}
		if b == true {
			count++
		}
	}

	return count
}

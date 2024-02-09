package main

import (
	"bufio"
	"fmt"
	"os"
)

// Function to find the longest repetition in a DNA sequence
func longestRepetition(dna string) int {
	/*
		METHOD: Counter
		TIME COMPLEXITY: O(n), где n - количество символов в строке ДНК. Это обусловлено тем, что алгоритм проходит по каждому символу в строке ровно один раз.
		SPACE COMPLEXITY: O(1), так как используется фиксированное количество переменных, которые не зависят от размера входных данных.
	*/
	maxRepetition := 1
	currentRepetition := 1

	for i := 1; i < len(dna); i++ {
		if dna[i] == dna[i-1] {
			currentRepetition++
			if currentRepetition > maxRepetition {
				maxRepetition = currentRepetition
			}
		} else {
			currentRepetition = 1
		}
	}

	return maxRepetition
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dna := scanner.Text()

	result := longestRepetition(dna)
	fmt.Println(result)
}

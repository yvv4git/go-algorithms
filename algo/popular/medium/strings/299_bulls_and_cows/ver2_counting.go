package main

import "fmt"

func getHintV2(secret string, guess string) string {
	bulls, mismatches := 0, 0
	diffs := make([]int, 10)

	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretDigit := secret[i] - '0'
			guessDigit := guess[i] - '0'
			diffs[secretDigit]++
			diffs[guessDigit]--
		}
	}

	for i := 0; i < len(diffs); i++ {
		if diffs[i] > 0 {
			mismatches += diffs[i]
		}
	}

	cows := len(secret) - bulls - mismatches

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

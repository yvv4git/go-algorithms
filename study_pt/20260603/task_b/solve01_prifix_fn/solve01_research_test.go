package main

import (
	"fmt"
	"testing"
)

// Последнее значение pi — это длина хвоста, который совпадает с началом.
// Если отбросить этот хвост, останется кусок, который повторяется. Это и есть основание.

// Time complexity: O(n) — каждый символ сравнивается не более n раз в сумме
// Space complexity: O(n) — массив pi длины n

func TestResearch(t *testing.T) {
	s := "bcabcab"
	n := len(s)
	pi := prefixFunction(s)
	period := n - pi[n-1]

	fmt.Println("Period: ", period) // 7 - 4 = 3 []int len: 7, cap: 7, [0,0,0,1,2,3,4]
}

func TestResearch2(t *testing.T) {
	s := "ababab"
	n := len(s)
	pi := prefixFunction(s)
	period := n - pi[n-1] // 6 - 4 = 2 []int len: 6, cap: 6, [0,0,1,2,3,4]

	fmt.Println("Period: ", period) // 2
}

// Решение в лоб — без префикс-функции, перебором.
// Сложность: O(n * количество делителей n) — в худшем O(n²)

func brutePeriod(s string) int {
	n := len(s)
	for p := 1; p <= n; p++ {
		if n%p != 0 {
			continue
		}
		ok := true
		for i := p; i < n; i++ {
			if s[i] != s[i-p] {
				ok = false
				break
			}
		}
		if ok {
			return p
		}
	}
	return n
}

func TestBruteForce(t *testing.T) {
	fmt.Println("Brute period 'ababab':", brutePeriod("ababab"))     // 2
	fmt.Println("Brute period 'bcabcab':", brutePeriod("bcabcab"))   // 7
	fmt.Println("Brute period 'aaaaaa':", brutePeriod("aaaaaa"))     // 1
	fmt.Println("Brute period 'abcdef':", brutePeriod("abcdef"))     // 6
}

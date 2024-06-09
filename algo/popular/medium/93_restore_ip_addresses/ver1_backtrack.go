package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	/*
		METHOD: Backtrack
		TIME COMPLEXITY: O(1) - O(3^4) в худшем случае, когда мы должны проверить все возможные комбинации разбиения строки на четыре части.
		SPACE COMPLEXITY:  O(1) в лучшем и худшем случае, так как мы используем фиксированное количество переменных, не зависящих от размера входных данных.
	*/
	var result []string
	backtrack(s, 0, 0, "", &result)
	return result
}

func backtrack(s string, start, step int, ip string, result *[]string) {
	// Если мы построили 4 части IP-адреса и прошли все символы в строке,
	// то добавляем полученный IP-адрес в результат.
	if start == len(s) && step == 4 {
		*result = append(*result, ip[:len(ip)-1])
		return
	}

	// Если мы прошли все символы строки или построили 4 части IP-адреса,
	// но остались неиспользованные символы, то это не допустимый IP-адрес.
	if len(s) == start || step == 4 {
		return
	}

	// Пробуем добавить разные длины к IP-адресу.
	for i := start; i < start+3 && i < len(s); i++ {
		part := s[start : i+1]
		// Проверяем, что часть не начинается с 0 и что она <= 255.
		if isValid(part) {
			// Если часть валидна, добавляем ее к IP-адресу и продолжаем сборку.
			backtrack(s, i+1, step+1, ip+part+".", result)
		}
	}
}

func isValid(s string) bool {
	// Проверяем, что часть не начинается с 0 и что она <= 255.
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	num, _ := strconv.Atoi(s)
	return num <= 255
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s)) // Вывод: [255.255.11.135 255.255.111.35]
}

package _3_roman_to_Integer

var ALPHABET = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToIntV4(s string) int {
	result := 0
	num_start, num_end := 0, 0
	for i := 1; i <= len(s); i++ {
		num_end = i
		if _, ok := ALPHABET[s[num_start:num_end]]; !ok {
			result += ALPHABET[s[num_start:num_end-1]]
			num_start = num_end - 1
		}
	}
	result += ALPHABET[s[num_start:num_end]]

	return result
}

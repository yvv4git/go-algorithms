package extract_ints_numbers_from_str

import (
	"fmt"
	"strings"
)

func ExtractNumbers(query string) ([]string, error) {
	if query == "" {
		return nil, fmt.Errorf("extract ports: %w", ErrEmptyQuery)
	}

	result := make([]string, 0)
	var str strings.Builder
	for i := 0; i < len(query); i++ {
		if query[i] >= '0' && query[i] <= '9' {
			str.WriteRune(rune(query[i]))
		} else if str.Len() > 0 {
			num := str.String()
			if num == "0" {
				return nil, fmt.Errorf("got wrong port number: %w", ErrInvalidPortNumber)
			}
			result = append(result, num)
			str.Reset()
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("parse string on ports: %w", ErrPortsNoFound)
	}

	return result, nil
}

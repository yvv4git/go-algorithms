package _68_excel_sheet_column_title

import "strings"

func convertToTitleV3(columnNumber int) string {
	/*
		METHOD: Map / Hash
		TIME COMPLEXITY: O(columnNumber)
		SPACE COMPLEXITY: O(columnNumber)
	*/
	m := map[int]uint8{
		1:  'A',
		2:  'B',
		3:  'C',
		4:  'D',
		5:  'E',
		6:  'F',
		7:  'G',
		8:  'H',
		9:  'I',
		10: 'J',
		11: 'K',
		12: 'L',
		13: 'M',
		14: 'N',
		15: 'O',
		16: 'P',
		17: 'Q',
		18: 'R',
		19: 'S',
		20: 'T',
		21: 'U',
		22: 'V',
		23: 'W',
		24: 'X',
		25: 'Y',
		26: 'Z',
	}

	result := strings.Builder{}

	count := 0

	if val, ok := m[columnNumber]; ok {
		return string(val)
	}

	for i := columnNumber; i > 0; i -= 26 {
		if val, ok := m[i]; ok {
			result.WriteByte(val)
		}
		count++
	}

	count--

	for i := count; i > 0; i /= 26 {
		if i != 26 {
			result.WriteByte(m[i%26])
		} else {
			result.WriteByte('Z')
			if i/26 == 1 {
				break
			}
		}
	}

	sb := strings.Builder{}
	for i := result.Len() - 1; i >= 0; i-- {
		sb.WriteString(string(result.String()[i]))
	}

	return sb.String()
}

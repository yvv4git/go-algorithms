package _68_excel_sheet_column_title

import (
	"fmt"
)

func convertToTitle(columnNumber int) string {
	/*
		METHOD: Loop
		TIME COMPLEXITY: O(log(columnNumber))
		SPACE COMPLEXITY: O(1)
	*/
	result := ""

	for columnNumber > 0 {
		columnTitleIndex := (columnNumber - 1) % 26
		columnNumber = (columnNumber - 1) / 26
		result = fmt.Sprintf("%c", rune(columnTitleIndex+65)) + result
	}

	return result
}

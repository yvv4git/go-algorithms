package _68_excel_sheet_column_title

import "fmt"

func convertToTitleV2(columnNumber int) string {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(log(columnNumber))
		SPACE COMPLEXITY: O(1)
	*/
	if columnNumber == 0 {
		return ""
	}

	columnNumber--

	return convertToTitle(columnNumber/26) + fmt.Sprintf("%c", (columnNumber%26)+'A')
}

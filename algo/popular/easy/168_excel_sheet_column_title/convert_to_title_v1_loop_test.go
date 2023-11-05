package _68_excel_sheet_column_title

import (
	"fmt"
	"testing"
)

func Test_convertToTitle(t *testing.T) {
	type args struct {
		columnNumber int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				columnNumber: 1,
			},
			want: "A",
		},
		{
			name: "CASE-2",
			args: args{
				columnNumber: 28,
			},
			want: "AB",
		},
		{
			name: "CASE-3",
			args: args{
				columnNumber: 701,
			},
			want: "ZY",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToTitleResearch01(t *testing.T) {
	columnNumber := 1
	result := ""

	colNum := columnNumber
	for colNum > 0 {
		columnTitleIndex := (colNum - 1) % 26
		colNum = (colNum - 1) / 26
		result = fmt.Sprintf("%c", rune(columnTitleIndex+65)) + result
	}

	//return result
	t.Logf("Res: %v", result)
}

func TestConvertToTitleResearch02(t *testing.T) {
	columnNumber := 1
	result := ""

	for columnNumber > 0 {
		columnTitleIndex := (columnNumber - 1) % 26 // Остаток от деления, позволяет найти остаток от деления одного числа на другое
		t.Logf("columnTitleIndex=%v", columnTitleIndex)
		columnNumber = (columnNumber - 1) / 26
		result = fmt.Sprintf("%c", rune(columnTitleIndex+65)) + result
	}

	t.Logf("Res: %v", result)
}

func TestDividedReminder(t *testing.T) {
	for x := 0; x < 27; x++ {
		y := 26
		t.Logf("%d %% %d  = %d", x, y, x%y)
	}
}

func TestDivided(t *testing.T) {
	for x := 0; x < 27; x++ {
		y := 26
		t.Logf("%d / %d  = %d", x, y, x/y) // Целая часть при делении
	}
}

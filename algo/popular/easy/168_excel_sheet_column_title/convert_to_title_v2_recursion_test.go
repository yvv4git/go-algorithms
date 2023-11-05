package _68_excel_sheet_column_title

import "testing"

func Test_convertToTitleV2(t *testing.T) {
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
			if got := convertToTitleV2(tt.args.columnNumber); got != tt.want {
				t.Errorf("convertToTitleV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertToTitleV2Research01(t *testing.T) {
	columnNumber := 28
	t.Logf("Div=%v", columnNumber/26) // Целая часть от деления
	t.Logf("Div=%v", columnNumber%26) // Остаток от деления
}

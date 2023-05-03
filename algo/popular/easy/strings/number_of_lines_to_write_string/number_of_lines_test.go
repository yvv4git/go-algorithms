package number_of_lines_to_write_string

import (
	"reflect"
	"testing"
)

func Test_numberOfLines(t *testing.T) {
	type args struct {
		widths []int
		s      string
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				widths: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, // 26 - это сколько занимает пикселей какая буква в нижнем регистре английского алфавита
				s:      "abcdefghijklmnopqrstuvwxyz",                                                                                  // Это текст, который надо разбить на строки.
			},
			want: []int{3, 60}, // There are a total of 3 lines, and the last line is 60 pixels wide.
		},
		{
			name: "CASE-2",
			args: args{
				widths: []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
				s:      "bbbcccdddaaa",
			},
			want: []int{2, 4}, // 2 - это количество строк; 4 - это ширина последней строки в пикселях.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfLines(tt.args.widths, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

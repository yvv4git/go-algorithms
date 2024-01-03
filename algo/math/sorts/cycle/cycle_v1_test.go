package cycle

import (
	"testing"
)

func Test_cycleSortV1(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				arr: []int{91, 28, 73, 46, 55, 64, 37, 82, 19},
			},
			want: 4, // 4 записи
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cycleSortV1(tt.args.arr); got != tt.want {
				t.Errorf("cycleSortV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

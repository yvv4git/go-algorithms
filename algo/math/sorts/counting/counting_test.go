package count

import (
	"reflect"
	"testing"
)

func Test_countingSort(t *testing.T) {
	type args struct {
		list []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				list: []int{91, 28, 73, 46, 55, 64, 37, 82, 19},
			},
			want: []int{19, 28, 37, 46, 55, 64, 73, 82, 91},
		},
		{
			name: "CASE-2",
			args: args{
				list: []int{5, 4, 3, 2, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "CASE-3",
			args: args{
				list: []int{5, 4, 3, -2, -1},
			},
			want: []int{-2, -1, 3, 4, 5},
		},
		{
			name: "CASE-4",
			args: args{
				list: []int{1},
			},
			want: []int{1},
		},
		{
			name: "CASE-5",
			args: args{
				list: []int{},
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			countingSort(tt.args.list)
			if !reflect.DeepEqual(tt.args.list, tt.want) {
				t.Errorf("we want %v, but got %v", tt.want, tt.args.list)
			}
		})
	}
}

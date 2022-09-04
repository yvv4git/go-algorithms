package _default

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
		A []int
	}

	testCases := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				A: []int{22, 8, 3, 31, 4, 2, 42, 1, 16, 6, 11, 25, 9, 8, 10, 12, 18, 14, 7, 15},
			},
			want: []int{1, 2, 3, 4, 6, 7, 8, 8, 9, 10, 11, 12, 14, 15, 16, 18, 22, 25, 31, 42},
		},
		{
			name: "CASE-2",
			args: args{
				A: []int{},
			},
			want: []int{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if result := Sort(tt.args.A); !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Sort() = %v, want %v", result, tt.want)
			}
		})
	}
}

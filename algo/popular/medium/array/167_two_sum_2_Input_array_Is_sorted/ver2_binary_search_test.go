package _67_two_sum_2_Input_array_Is_sorted

import (
	"reflect"
	"testing"
)

func Test_twoSumV2(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			name: "Test Case 2",
			args: args{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
		{
			name: "Test Case 3",
			args: args{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumV2(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

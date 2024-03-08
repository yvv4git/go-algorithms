package main

import (
	"reflect"
	"testing"
)

func Test_grayCode(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1: n = 1",
			args: args{n: 1},
			want: []int{0, 1},
		},
		{
			name: "Test Case 2: n = 2",
			args: args{n: 2},
			want: []int{0, 1, 3, 2},
		},
		{
			name: "Test Case 3: n = 3",
			args: args{n: 3},
			want: []int{0, 1, 3, 2, 6, 7, 5, 4},
		},
		{
			name: "Test Case 4: n = 0 (edge case)",
			args: args{n: 0},
			want: []int{},
		},
		{
			name: "Test Case 5: n = -1 (edge case)",
			args: args{n: -1},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grayCode(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_lexicographicalNumbers(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{n: 1},
			want: []int{1},
		},
		{
			name: "Test Case 2",
			args: args{n: 2},
			want: []int{1, 2},
		},
		{
			name: "Test Case 3",
			args: args{n: 10},
			want: []int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Test Case 4",
			args: args{n: 13},
			want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Test Case 5",
			args: args{n: 20},
			want: []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Test Case 6",
			args: args{n: 0},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicographicalNumbers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicographicalNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

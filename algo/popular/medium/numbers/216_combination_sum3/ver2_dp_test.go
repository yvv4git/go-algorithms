package main

import (
	"reflect"
	"testing"
)

func Test_combinationSum3V2(t *testing.T) {
	type args struct {
		k int
		n int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				k: 3,
				n: 7,
			},
			want: [][]int{{1, 2, 4}},
		},
		//{
		//	name: "Test Case 2",
		//	args: args{
		//		k: 3,
		//		n: 9,
		//	},
		//	want: [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}},
		//},
		{
			name: "Test Case 5",
			args: args{
				k: 9,
				n: 45,
			},
			want: [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3V2(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3V2() = %v, want %v", got, tt.want)
			}
		})
	}
}

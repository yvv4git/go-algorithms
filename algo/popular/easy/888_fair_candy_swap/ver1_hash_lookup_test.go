package main

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	type args struct {
		aliceSizes []int
		bobSizes   []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{
				aliceSizes: []int{1, 1},
				bobSizes:   []int{2, 2},
			},
			want: []int{1, 2},
		},
		{
			name: "Test Case 2",
			args: args{
				aliceSizes: []int{1, 2},
				bobSizes:   []int{2, 3},
			},
			want: []int{1, 2},
		},
		{
			name: "Test Case 3",
			args: args{
				aliceSizes: []int{2},
				bobSizes:   []int{1, 3},
			},
			want: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap(tt.args.aliceSizes, tt.args.bobSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}

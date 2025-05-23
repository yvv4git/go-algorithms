package main

import (
	"reflect"
	"testing"
)

func Test_runningSumV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 1, 2, 10, 1},
			},
			want: []int{3, 4, 6, 16, 17},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSumV2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("runningSumV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

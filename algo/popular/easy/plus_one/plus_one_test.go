package main

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				digits: []int{1, 2, 3},
			},
			want: []int{1, 2, 4},
		},
		{
			name: "CASE-2",
			args: args{
				digits: []int{4, 3, 2, 1},
			},
			want: []int{4, 3, 2, 2},
		},
		{
			name: "CASE-3",
			args: args{
				digits: []int{9},
			},
			want: []int{1, 0},
		},
		{
			name: "CASE-4",
			args: args{
				digits: []int{},
			},
			want: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := plusOne(tt.args.digits); !reflect.DeepEqual(result, tt.want) {
				t.Errorf("plusOne() = %v, want %v", result, tt.want)
			}
		})
	}
}

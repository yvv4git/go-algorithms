package main

import (
	"reflect"
	"testing"
)

func Test_countElements(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			name: "Test Case 1: All elements are the same",
			args: args{
				arr: []int{4, 4, 4, 4},
			},
			want: map[int]int{4: 4},
		},
		{
			name: "Test Case 2: Different elements",
			args: args{
				arr: []int{1, 2, 3, 4},
			},
			want: map[int]int{1: 1, 2: 1, 3: 1, 4: 1},
		},
		{
			name: "Test Case 3: Repeated elements",
			args: args{
				arr: []int{1, 1, 2, 2, 3, 3, 4, 4},
			},
			want: map[int]int{1: 2, 2: 2, 3: 2, 4: 2},
		},
		{
			name: "Test Case 4: Empty array",
			args: args{
				arr: []int{},
			},
			want: map[int]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countElements(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

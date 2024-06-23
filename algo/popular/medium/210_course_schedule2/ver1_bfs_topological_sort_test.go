package main

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "No prerequisites",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{},
			},
			want: []int{0, 1, 2, 3}, // Any order is valid
		},
		{
			name: "Simple prerequisites",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			},
			want: []int{0, 1, 2, 3}, // One valid order
		},
		{
			name: "Cycle in prerequisites",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 1}, {0, 2}},
			},
			want: []int{}, // No valid order due to cycle
		},
		{
			name: "Single course",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: []int{0}, // Only one course
		},
		{
			name: "Multiple valid orders",
			args: args{
				numCourses:    3,
				prerequisites: [][]int{{1, 0}, {2, 1}},
			},
			want: []int{0, 1, 2}, // One valid order
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

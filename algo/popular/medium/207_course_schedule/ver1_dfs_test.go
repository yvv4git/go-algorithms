package main

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "No prerequisites",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{},
			},
			want: true,
		},
		{
			name: "Simple linear prerequisites",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}},
			},
			want: true,
		},
		{
			name: "Cycle in prerequisites",
			args: args{
				numCourses:    2,
				prerequisites: [][]int{{1, 0}, {0, 1}},
			},
			want: false,
		},
		{
			name: "Multiple courses with no cycle",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}},
			},
			want: true,
		},
		{
			name: "Multiple courses with a cycle",
			args: args{
				numCourses:    4,
				prerequisites: [][]int{{1, 0}, {2, 1}, {3, 2}, {0, 3}},
			},
			want: false,
		},
		{
			name: "Single course with no prerequisites",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{},
			},
			want: true,
		},
		{
			name: "Single course with a self-cycle",
			args: args{
				numCourses:    1,
				prerequisites: [][]int{{0, 0}},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}

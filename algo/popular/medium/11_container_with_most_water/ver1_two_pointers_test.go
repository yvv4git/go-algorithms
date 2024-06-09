package main

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			name: "Test Case 2",
			args: args{
				height: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				height: []int{4, 3, 2, 1, 4},
			},
			want: 16,
		},
		{
			name: "Test Case 4",
			args: args{
				height: []int{1, 2, 1},
			},
			want: 2,
		},
		{
			name: "Test Case 5",
			args: args{
				height: []int{2, 3, 4, 5, 18, 17, 6},
			},
			want: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

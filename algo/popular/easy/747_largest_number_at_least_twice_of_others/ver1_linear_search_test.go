package main

import "testing"

func Test_dominantIndex(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// Тестовые кейсы
		{
			name: "Example 1",
			args: args{nums: []int{3, 6, 1, 0}},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 2, 3, 4}},
			want: -1,
		},
		{
			name: "Single Element",
			args: args{nums: []int{1}},
			want: -1,
		},
		{
			name: "Two Elements, Dominant",
			args: args{nums: []int{1, 2}},
			want: 1,
		},
		{
			name: "All Elements Equal",
			args: args{nums: []int{2, 2, 2, 2}},
			want: -1,
		},
		{
			name: "Largest Element Twice as Large",
			args: args{nums: []int{1, 2, 4, 8}},
			want: 3,
		},
		{
			name: "Largest Element Not Twice as Large",
			args: args{nums: []int{1, 2, 4, 7}},
			want: -1,
		},
		{
			name: "Mixed Positive and Negative Numbers",
			args: args{nums: []int{-1, 2, -4, 8}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dominantIndex(tt.args.nums); got != tt.want {
				t.Errorf("dominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

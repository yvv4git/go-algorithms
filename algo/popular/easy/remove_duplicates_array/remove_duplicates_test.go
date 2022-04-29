package main

import "testing"

// Как я понимаю, задача этой фукнции - посчитать количество элементов, если удалить дубликаты.
func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "CASE-4",
			args: args{
				nums: []int{1, 2, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "CASE-5",
			args: args{
				nums: []int{1, 1, 1, 1, 1, 1},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

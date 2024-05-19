package main

import "testing"

func Test_pivotIndexV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: Pivot index exists",
			args: args{nums: []int{1, 7, 3, 6, 5, 6}},
			want: 3, // Ожидаемый индекс опорного элемента
		},
		{
			name: "Test Case 2: Pivot index does not exist",
			args: args{nums: []int{1, 2, 3}},
			want: -1, // Ожидаемый результат, когда опорного индекса не существует
		},
		{
			name: "Test Case 3: Array with negative numbers",
			args: args{nums: []int{-1, -1, -1, -1, -1, 0}},
			want: 2, // Ожидаемый индекс опорного элемента
		},
		{
			name: "Test Case 4: Array with zeros",
			args: args{nums: []int{0, 0, 0, 0, 0, 0}},
			want: 0, // Ожидаемый индекс опорного элемента
		},
		{
			name: "Test Case 5: Single element array",
			args: args{nums: []int{5}},
			want: 0, // Ожидаемый индекс опорного элемента
		},
		//{
		//	name: "Test Case 6: Empty array",
		//	args: args{nums: []int{}},
		//	want: -1, // Ожидаемый результат, когда массив пустой
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndexV2(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndexV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_sumOddLengthSubarraysV3(t *testing.T) {
	type args struct {
		arr []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1: Small array",
			args: args{arr: []int{1, 4, 2, 5, 3}},
			want: 58, // Подмассивы: [1], [4], [2], [5], [3], [1,4,2], [4,2,5], [2,5,3], [1,4,2,5,3]
		},
		{
			name: "Test case 2: Single element",
			args: args{arr: []int{10}},
			want: 10, // Единственный подмассив: [10]
		},
		{
			name: "Test case 4: Empty array",
			args: args{arr: []int{}},
			want: 0, // Нет подмассивов
		},
		{
			name: "Test case 5: Array with two elements",
			args: args{arr: []int{1, 2}},
			want: 3, // Подмассивы: [1], [2]
		},
		{
			name: "Test case 6: Array with three elements",
			args: args{arr: []int{1, 2, 3}},
			want: 12, // Подмассивы: [1], [2], [3], [1,2,3]
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOddLengthSubarraysV3(tt.args.arr); got != tt.want {
				t.Errorf("sumOddLengthSubarraysV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

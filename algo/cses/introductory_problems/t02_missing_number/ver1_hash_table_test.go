package main

import "testing"

func Test_findMissingNumberV2(t *testing.T) {
	type args struct {
		numbers []int
		N       int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Тест 1: Отсутствует число 5",
			args: args{numbers: []int{2, 3, 1, 5}, N: 5},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissingNumberV1(tt.args.numbers, tt.args.N); got != tt.want {
				t.Errorf("findMissingNumberV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_numberOfArithmeticSlicesV3(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// Базовые случаи
		{"Empty array", args{[]int{}}, 0},
		{"Single element", args{[]int{1}}, 0},
		{"Two elements", args{[]int{1, 2}}, 0},
		{"Three elements, arithmetic slice", args{[]int{1, 2, 3}}, 1},
		{"Three elements, no arithmetic slice", args{[]int{1, 3, 2}}, 0},

		// Случаи с несколькими арифметическими срезами
		{"Multiple arithmetic slices", args{[]int{1, 2, 3, 4, 5}}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfArithmeticSlicesV3(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlicesV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"fmt"
	"testing"
)

func Test_singleNumber(t *testing.T) {
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
				nums: []int{2, 2, 1},
			},
			want: 1,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{4, 1, 2, 1, 2},
			},
			want: 4,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{4, 1, 2, 1, 2, 6, 6, 7, 7},
			},
			want: 4,
		},
		{
			name: "CASE-4",
			args: args{
				nums: []int{4, 1, 2, 1, 2, 6, 6, 7, 7, 4},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ------------

func TestXor01(t *testing.T) {
	testCases := []struct {
		nums []int
	}{
		{
			nums: []int{1, 1},
		},
	}

	for _, tc := range testCases {
		res := 0
		for _, num := range tc.nums {
			fmt.Printf("num=%d[%08b] res=%d[%08b] res^num=%d[%08b] \n",
				num, num,
				res, res,
				res^num, res^num,
			)
			res ^= num
		}
	}
}

func TestXor02(t *testing.T) {
	/*
		Операция xor работает так:
		- если оба операнда равны, то значение становится равно 0.
		- если оба операнда не равны, то работает как сложение. 1 xor 2.
	*/
	x := 0
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%d xor x=%d", i, x)
		x ^= i
		fmt.Printf(" = x=%d \n", x)
	}
}

func TestXor03(t *testing.T) {
	/*
		XOR - Bitwise Operator.
		0 ^ 0 = 0
		0 ^ 1 = 1
		1 ^ 0 = 1
		1 ^ 1 = 0
	*/
	x := 0
	y := 0
	res := x ^ y
	t.Logf("%d(%08b) xor %d(%08b) = %d(%08b)", x, x, y, y, res, res)

	x = 1
	y = 0
	res = x ^ y
	t.Logf("%d(%08b) xor %d(%08b) = %d(%08b)", x, x, y, y, res, res)

	x = 1
	y = 1
	res = x ^ y
	t.Logf("%d(%08b) xor %d(%08b) = %d(%08b)", x, x, y, y, res, res)

	x = 1
	y = 2
	res = x ^ y // 3 - работает как сложение.
	t.Logf("%d(%08b) xor %d(%08b) = %d(%08b)", x, x, y, y, res, res)

	x = 1
	y = 3
	res = x ^ y // 2- работает как вычитание.
	t.Logf("%d(%08b) xor %d(%08b) = %d(%08b)", x, x, y, y, res, res)

	x = 1
	y = 4
	res = x ^ y // 7 - работает как сложение.
	t.Logf("%d(%08b) xor %d(%08b) = %d(%08b)", x, x, y, y, res, res)

	x = 1
	y = 5
	res = x ^ y // 4 - работает как вычитание.
	t.Logf("%d(%08b) xor %d(%08b) = %d(%08b)", x, x, y, y, res, res)

	x = 1
	y = 6
	res = x ^ y // 7 - работает как сложение.
	t.Logf("%d(%08b) xor %d(%08b) = %d(%08b)", x, x, y, y, res, res)

	x = 2
	y = 4
	res = x ^ y // 6 - работает как сложение.
	t.Logf("%d(%08b) xor %d(%08b) = %d(%08b)", x, x, y, y, res, res)
}

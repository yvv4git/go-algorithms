package minimum_bit_flip_convert_number

import (
	"fmt"
	"testing"
)

func Test_minBitFlipsV3(t *testing.T) {
	type args struct {
		start int
		goal  int
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				start: 10,
				goal:  7,
			},
			want: 3,
			desc: `The binary representation of 10 and 7 are 1010 and 0111 respectively. We can convert 10 to 7 in 3 steps:
					- Flip the first bit from the right: 1010 -> 1011.
					- Flip the third bit from the right: 1011 -> 1111.
					- Flip the fourth bit from the right: 1111 -> 0111.
					It can be shown we cannot convert 10 to 7 in less than 3 steps. Hence, we return 3.`,
		},
		{
			name: "CASE-2",
			args: args{
				start: 3,
				goal:  4,
			},
			want: 3,
			desc: `The binary representation of 3 and 4 are 011 and 100 respectively. We can convert 3 to 4 in 3 steps:
					- Flip the first bit from the right: 011 -> 010.
					- Flip the second bit from the right: 010 -> 000.
					- Flip the third bit from the right: 000 -> 100.
					It can be shown we cannot convert 3 to 4 in less than 3 steps. Hence, we return 3.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minBitFlipsV3(tt.args.start, tt.args.goal); got != tt.want {
				t.Errorf("minBitFlipsV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -----
func TestResearchV3ex01(t *testing.T) {
	start := 10
	goal := 7
	fmt.Printf("start: %d(%08b) \n", start, start) // start: 10(00001010)
	fmt.Printf("goal:   %d(%08b) \n", goal, goal)  // goal:   7(00000111)

	/*
		Операция XOR покажет 1-ами места, в которых числа отличаются. На этих местах в изначальном числе(start) надо сделать замены(flip).
	*/
	ones := start ^ goal
	fmt.Printf("ones:  %d(%08b) \n", ones, ones) // ones:  13(00001101) - после этой операции, на местах, которые надо flip(изменить) будут 1

	/*
		Вот эта формула(ones &= ones - 1) на каждом шаге убирает ровно одну 1 из ones.
		Таким образом, мы посчитаем сколько знаков надо flip(заменить) в start, чтобы получить goal.
	*/
	flips := 0
	for ; ones != 0; ones &= ones - 1 {
		fmt.Printf("ones-1: %d(%08b) => ones &= ones - 1: %d(%08b) \n", ones-1, ones-1, ones&(ones-1), ones&(ones-1))
		flips++
	}

	fmt.Printf("flips: %d(%08b)\n", flips, flips) // flips: 3(00000011) - количество бит, которые нужно поменять!
}

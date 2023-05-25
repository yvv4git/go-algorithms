package minimum_bit_flip_convert_number

import (
	"fmt"
	"testing"
)

func Test_minBitFlipsV2(t *testing.T) {
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
			if got := minBitFlipsV2(tt.args.start, tt.args.goal); got != tt.want {
				t.Errorf("minBitFlipsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ------------------------------
func TestResearchKernighan01(t *testing.T) {
	start := 10
	goal := 7
	result := 0
	fmt.Printf("start: %d(%08b)\n", start, start)
	fmt.Printf("goal: %d(%08b)\n", goal, goal)
	println()

	fmt.Printf("init x: start ^ goal = %d(%08b)", start^goal, start^goal)
	for x := start ^ goal; x != 0; x &= x - 1 {
		fmt.Printf("x: %d(%08b) \n", x, x)
		result++
	}
	fmt.Printf("result: %d(%08b) \n", result, result) // result: 3(00000011)
}

func TestResearchKernighan02(t *testing.T) {
	x := 10
	fmt.Printf("x: %d(%08b) \n", x, x)       // x: 10(00001010)
	fmt.Printf("x-1: %d(%08b) \n", x-1, x-1) // x-1: 9(00001001)

	x &= x - 1
	fmt.Printf("x: %d(%08b) \n", x, x) // x: 8(00001000)
}

func TestResearchKernighan03(t *testing.T) {
	x := 10
	fmt.Printf("x: %d(%08b) \n", x, x) // x: 10(00001010)
	for x > 0 {
		fmt.Printf("- x-1: %d(%08b) x&(x-1): %d(%08b) \n", x-1, x-1, x&(x-1), x&(x-1))
		x &= x - 1
		fmt.Printf("x: %d(%08b) \n", x, x)
	}
}

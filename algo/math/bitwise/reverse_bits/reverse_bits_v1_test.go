package reverse_bits

import (
	"fmt"
	"testing"
)

func Test_reverseBitsV1(t *testing.T) {
	type args struct {
		n int
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
				n: 11,
			},
			want: 13,
			desc: `Explanation: 11(1011) -> 13(1101)`,
		},
		{
			name: "CASE-2",
			args: args{
				n: 10,
			},
			want: 5,
			desc: `Explanation: 11(1010) -> 5(00000101)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBitsV1(tt.args.n); got != tt.want {
				t.Errorf("reverseBitsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func TestResearchMethodV1Ex01(t *testing.T) {
	n := 11
	fmt.Printf("n=%d(%08b) \n", n, n)

	result := 0

	for n > 0 {
		result <<= 1 // Left shift by 1
		fmt.Printf("result: %d(%08b) n=%d(%08b) n&1=%d(%08b) \n", result, result, n, n, n&1, n&1)
		if n&1 == 1 {
			result = result ^ 1
		}

		n >>= 1 // Right shift by 1
	}

	t.Logf("Result: %d(%08b)", result, result)
}

func TestResearchMethodV1Ex02(t *testing.T) {
	n := 11
	fmt.Printf("n=%d(%08b) \n", n, n)

	result := 0
	for n > 0 {
		/*
			Обрати внимание, что данная операция сдвигает последовательность 0 и 1 влево, а в младшем(правом) разряде образуется 0.
		*/
		result <<= 1 // Left shift - эквивалент умножения на 2

		/*
			Зачем делать бинарное И (&)?
			+ 11(00001011) & 1(00000001) = 1(00000001) - вернет 1
			+ 5(00000101) &  1(00000001) = 1(00000001) - вернет 1
			+ 0(00000010) &  1(00000001) = 0(00000000) - вернет 0

			Таким образом, мы можем узнать - в самом младшем(правом) разряде находится 1 или 0.
		*/
		if n&1 == 1 {
			fmt.Printf("+n=%d(%08b) n&1=%d(%08b) \n", n, n, n&1, n&1)
			/*
				Если в младшем(правом) разряде находится 1 у исходного числа, то в результате в младшем(правом) бите надо выставить 1. (result = result ^ 1)
			*/
			result = result ^ 1
		} else {
			// В данном случае у result в младшем(правом) разряде останется 0.
			fmt.Printf("-n=%d(%08b) n&1=%d(%08b) \n", n, n, n&1, n&1)
		}
		n >>= 1 // Right shift - эквивалент деления на 2
	}

	t.Logf("Result: %d(%08b)", result, result)
}

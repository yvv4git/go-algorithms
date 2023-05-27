package number_of_even_and_odd

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_evenOddBitV3(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				n: 17,
			},
			want: []int{2, 0},
			desc: `Explanation: The binary representation of 17 is 10001. 
					It contains 1 on the 0th and 4th indices. 
					There are 2 even and 0 odd indices.`,
		},
		{
			name: "CASE-2",
			args: args{
				n: 2,
			},
			want: []int{0, 1},
			desc: `Explanation: The binary representation of 2 is 10.
					It contains 1 on the 1st index. 
					There are 0 even and 1 odd indices.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evenOddBitV3(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evenOddBitV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func TestResearchMethodV3Ex01(t *testing.T) {
	n := 17
	arr := []int{0, 0}
	var p int = 0

	for n > 0 {
		if n%2 == 1 { // Остаток от деления n на 2 равен 1. При чем это не остаток, а модуль. Т.е. n по модулю 2.
			arr[p%2]++ // При делении по модулю 2 мы получаем либо 0, либо первый индекс в массиве
		}
		p++
		fmt.Printf("p=%d n=%d(%08b) n%%2=%d(%08b) \n", p, n, n, n%2, n%2)
		n /= 2 // Эквивалентно right shift, сдвигу право
	}

	t.Logf("result: %v", arr)
	/*
		При делении n / 2 число как бы сдвигается вправо(Right shift) и тем самым
		p=1 n=17(00010001) n%2=1(00000001) <- обрати внимание, число делится на 2 без остатка, если младший(правый) бит = 0, если там 1, то получаем остаток.
		p=2 n=8(00001000) n%2=0(00000000)
		p=3 n=4(00000100) n%2=0(00000000)
		p=4 n=2(00000010) n%2=0(00000000)
		p=5 n=1(00000001) n%2=1(00000001)
	*/
}

func TestResearchMethodV3Ex02(t *testing.T) {
	n := 17
	t.Logf("n: %d(%08b) \n", n, n)

	r := n % 2
	t.Logf("2:  %d(%08b) \n", 2, 2)
	t.Logf("r: %d(%08b) \n", r, r)
}

func TestResearchMethodV3Ex03(t *testing.T) {
	/*
		Чётность нуля — общепринятый математический факт, который, однако, иногда вызывает сомнения у людей, недостаточно знакомых с математикой.
		То, что нуль — чётное число, сразу следует из определения чётного числа. По определению, чётное число — целое число, которое делятся на 2 без остатка.
		Ноль полностью удовлетворяет этому определению. Он также обладает всеми свойствами чётных чисел — например, он с обеих сторон граничит с нечётными числами.
	*/
	for i := 0; i < 10; i++ {
		r := i % 2 // Если число по модулю два дает остаток 1, значит это нечетное число(odd)
		if r == 1 {
			fmt.Printf("i(%d[%08b]) %% 2 = %d(%08b) ---> odd(нечетное)\n", i, i, r, r)
		} else {
			fmt.Printf("i(%d[%08b]) %% 2 = %d(%08b) ---> even(четное)\n", i, i, r, r)
		}
	}
}

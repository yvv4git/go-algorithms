package count_number_of_consistent_strings

import (
	"fmt"
	"testing"
)

func Test_countConsistentStringsV2(t *testing.T) {
	type args struct {
		allowed string
		words   []string
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
				allowed: "ab",
				words:   []string{"ad", "bd", "aaab", "baa", "badab"},
			},
			want: 2,
			desc: `Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.`,
		},
		{
			name: "CASE-2",
			args: args{
				allowed: "abc",
				words:   []string{"a", "b", "c", "ab", "ac", "bc", "abc"},
			},
			want: 7,
			desc: `All strings are consistent.`,
		},
		{
			name: "CASE-3",
			args: args{
				allowed: "cad",
				words:   []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"},
			},
			want: 4,
			desc: `Explanation: Strings "cc", "acd", "ac", and "d" are consistent.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countConsistentStringsV2(tt.args.allowed, tt.args.words); got != tt.want {
				t.Errorf("countConsistentStringsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -----
func TestResearchV2Ex01(t *testing.T) {
	/*
		Основной код

		1 << (char - 'a') можно разделить на 2 части:
		Часть-1
		char - 'a' - здесь мы находим порядковый номер символа в алфавите. Это работает потому, что символ 'a' и 'b' по числовым
		кодам отличаются на 1. Тогда, если вычесть из b - a = 1.

		Часть-2
		1 << ...
		Операция сдвига влево на 1 то же самое, что умножение на 2.
	*/
	word := "ad"
	var bitMask int = 0

	fmt.Printf("%v \n", word)

	for _, char := range word {
		oldBM := bitMask
		r1 := char - 'a'
		bitMask = bitMask | (1 << (char - 'a'))
		fmt.Printf("bitMask[%d %08b] = bitMask[%d %08b] | (1 << (char[%c %d %08b] - a) = [%d %08b]) \n", bitMask, bitMask, oldBM, oldBM, char, char, char, r1, r1)
	}

	t.Logf("Result: %#v", bitMask)
}

func TestResearchV2Ex02(t *testing.T) {
	/*
		Разбираю, что делает эта часть:
		char - 'a'   <--- таким образом мы определяем позицию символа в алфавите
	*/
	a := 'a'
	fmt.Printf("a = %d[%08b]\n", a, a) // a = 97[01100001]

	char := 'a'
	r := char - 'a'
	fmt.Printf("char(%c %d %08b)-'a' = %d(%08b) \n", char, char, char, r, r) // char(a 97 01100001)-'a' = 0(00000000)
}

func TestResearchV2Ex03(t *testing.T) {
	data := []rune{'a', 'b', 'c', 'd'}
	for _, char := range data {
		fmt.Printf("char: %c %d[%08b] \n", char, char, char)
	}

	fmt.Println()
	a := 'a'
	for _, char := range data {
		r := char - a
		fmt.Printf("char[%c %d %08b] - a[%c %d %08b] = r[%c %d %08b]\n", char, char, char, a, a, a, r, r, r)
	}

	fmt.Println()
	for _, char := range data {
		r1 := char - a
		r := 1 << r1 // Операция сдвига на 1 влево эквивалентная умножению на 2. Т.е. r1 * 2
		fmt.Printf("1 << (char[%c %d %08b] - a[%c %d %08b] = r1[%c %d %08b]) = r[%c %d %08b]\n", char, char, char, a, a, a, r1, r1, r1, r, r, r)
	}
}

func TestResearchV2Ex04(t *testing.T) {
	r1 := 0
	fmt.Printf("r1 = %c %d %08b\n", r1, r1, r1)

	r := 1 << r1
	fmt.Printf("r = %d %c %08b\n", r, r, r)
}

func TestResearchV2Ex05(t *testing.T) {
	/*
		Операция сдвига на 1 влево: 1 << x
		Эквивалентная умножению x на 2^1 = x * 2.
		Т.е. можно сказать, что мы просто умножаем на 2.
	*/
	x := 0
	for i := 0; i < 10; i++ {
		x <<= 1
	}
	fmt.Printf("x=%d\n", x)
}

func TestResearchV2Ex06(t *testing.T) {
	x := 5
	a := 1
	r := a << x
	fmt.Printf("a[%d %08b] << x[%d %08b] = r[%d %08b] \n", a, a, x, x, r, r)
}

func TestResearchV2Ex07(t *testing.T) {
	a := 1 // Первый операнд для Right Shift
	for b := 0; b < 5; b++ {
		r := a << b
		fmt.Printf("r[%d %08b] = a[%b %08b] << b[%d %08b] \n", r, r, a, a, b, b)
	}
	/*
		Попробуем число a = 1 сдвигать влево(left shift) на 0, 1, 2, 3, 4, 5 разрядов.
		r[1 00000001] = a[1 00000001] << b[0 00000000] <--- с
		r[2 00000010] = a[1 00000001] << b[1 00000001]
		r[4 00000100] = a[1 00000001] << b[2 00000010]
		r[8 00001000] = a[1 00000001] << b[3 00000011]
		r[16 00010000] = a[1 00000001] << b[4 00000100]
	*/

}

// bitMask
func TestResearchV2Ex08(t *testing.T) {
	bitMask := 0
	r1 := 1 << 2
	r2 := bitMask | r1
	fmt.Printf("r2[%d %08b] = bitMask[%d %08b] | r1[%d %08b] \n", r2, r2, bitMask, bitMask, r1, r1)
	/*
		r2[4 00000100] = bitMask[0 00000000] | r1[4 00000100]
	*/

	fmt.Println()
	/*
		Как работает оператор ИЛИ?
		- если в обоих битах операндов 1, то в результирующем бите будет 1.
		- если в одном из операндов 1, то на выходе будет 1.
		- если в обоих операнда 0, то на выходе будет 0.
	*/
	a := 1
	b := 3
	r := a | b
	fmt.Printf("a[%d %08b] | b[%d %08b] = r[%d %08b] \n", a, a, b, b, r, r)
}

func TestResearchV2Ex09(t *testing.T) {
	word := "ad"
	var bitMask int = 0

	fmt.Printf("%v \n", word)

	for _, char := range word {
		r1 := char - 'a'
		fmt.Printf("char[%c %d %08c] - 'a'[%d %08b] = r[%c %d %08b]\n", char, char, char, 'a', 'a', r1, r1, r1)
		bitMask = bitMask | (1 << (char - 'a'))

	}

	t.Logf("Result: %#v", bitMask)
}

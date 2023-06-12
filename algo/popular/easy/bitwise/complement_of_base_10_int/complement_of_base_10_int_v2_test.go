package complement_of_base_10_int

import (
	"fmt"
	"math/bits"
	"testing"
)

func Test_bitwiseComplementV2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 5,
			},
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				n: 7,
			},
			want: 0,
		},
		{
			name: "CASE-3",
			args: args{
				n: 10,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseComplementV2(tt.args.n); got != tt.want {
				t.Errorf("bitwiseComplementV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -----
func TestResearchV2Ex01(t *testing.T) {
	n := 5
	//if n == 0 {
	//	fmt.Printf("result: 1 \n")
	//	return
	//}

	/*
		Определяем длину числа в битах.
	*/
	shift := bits.Len(uint(n))
	fmt.Printf("shift=%d \n", shift)

	/*
		Тут выполняется 3 действия:
		- a = 1<<shift - нужно, чтобы получить число на разряд больше. Например, n=5 занимает 3 разряда. Значит нам надо число из 4-разрядов.
		- b = a - 1	- из 4-разрядного числа вычитаем 1, чтобы получить 3 разряда, заполненные 1.
		- n ^ b - n=5(101) ^ b(111) = 2(010) <--- 2 комплементарное число.
	*/
	result := n ^ (1<<shift - 1)
	fmt.Printf("result: %d[%08b] \n", result, result)
}

func TestResearchV2Ex02(t *testing.T) {
	n := 5
	fmt.Printf("n = %d[%08b] \n", n, n)

	/*
		Определяем длину числа в битах.
	*/
	shift := bits.Len(uint(n))
	fmt.Printf("shift = %d[%08b] \n", shift, shift)

	/*
		Сдвигаем 1 на 3 разряда влево. Это эквивалентно умножению 1 * 2^3 = 1 * 8.
	*/
	a := 1 << shift
	fmt.Printf("a = %d[%08b]\n", a, a)

	/*
		XOR-им оставшиеся значения.
	*/
	r := n ^ a
	fmt.Printf("r[%d %08b] = n[%d %08b] ^ a[%d %08b]\n", r, r, n, n, a, a)
}

func TestResearchV2Ex03(t *testing.T) {
	n := 5
	shift := 3

	/*
		Порядок операций:
		- сначала выполнится операция сдвига влево: 1<<shift
		- затем выполнится операция вычитания: (1<<shift) - 1
	*/
	result := n ^ (1<<shift - 1)
	fmt.Printf("result: %d[%08b] \n", result, result)

	a0 := 1<<shift - 1
	fmt.Printf("a0 = %d[%08b] \n", a0, a0)

	a1 := 1 << shift
	fmt.Printf("a1 = %d[%08b]\n", a1, a1)

	a2 := a1 - 1
	fmt.Printf("a2 = %d[%08b]\n", a2, a2) // a2 = 7[00000111]
}

func TestResearchV2Ex04(t *testing.T) {
	n := 5     // Число, которому надо найти комплементарное число. Т.е. можно сказать противоположность по битам.
	shift := 3 // Количество разрядов, которое используется для хранения числа 5.
	fmt.Printf("n = %d[%08b] \n", n, n)
	fmt.Printf("shift = %d[%08b] \n", shift, shift)

	// Сдвиг на 1 влево эквивалентно умножению 1 * 2^shift = 1 * 2 ^ 3 = 1 * 8 = 8
	a1 := 1 << shift
	fmt.Printf("a1 = %d[%08b] \n", a1, a1)

	/*
		По сути дела, a2 = 111 - это маска из 1, чтобы можно было XOR-ить n=5(101). Для этого и нужна маска 111.
	*/
	a2 := a1 - 1
	fmt.Printf("1 = %08b \n", 1)
	fmt.Printf("a2 = %d[%08b] \n", a2, a2)
}

func TestResearchV2Ex05(t *testing.T) {
	for i := 0; i < 10; i++ {
		r := 1 << i
		r2 := r - 1
		fmt.Printf("1(%08b) << i[%d %08b] = 1 * 2^%d = r[%d %08b] r2=[%d %08b] \n", 1, i, i, i, r, r, r2, r2)
	}
}

package number_complement

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test_findComplementV2(t *testing.T) {
	type args struct {
		num int
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
				num: 5,
			},
			want: 2,
			desc: "The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.",
		},
		{
			name: "CASE-2",
			args: args{
				num: 1,
			},
			want: 0,
			desc: "The binary representation of 1 is 1 (no leading zero bits), and its complement is 0. So you need to output 0.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findComplementV2(tt.args.num); got != tt.want {
				t.Errorf("findComplementV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ------------------------------
func getBytesInt(x int) [8]byte {
	return *(*[8]byte)(unsafe.Pointer(&x))
}

func TestNegativeNum07(t *testing.T) {
	/*
		Ищем комплементарное число.
		Например, для 5(101) комплементарным будет 2(010).
	*/
	mask := ^0
	fmt.Printf("mask_bytes: %08b \n", getBytesInt(mask)) // mask_bytes: [11111111 11111111 11111111 11111111 11111111 11111111 11111111 11111111] <- все единицы

	num := 5
	fmt.Printf("num[%d]: %08b \n", num, getBytesInt(num)) // num[5]: [00000101 00000000 00000000 00000000 00000000 00000000 00000000 00000000]

	for (num & mask) > 0 {
		/*
			Смысл данного действия - сдвигать маску, пока 0 в маске не покроет все 1 в num. В результате их num & mask вернет 0.
		*/
		println()
		mask <<= 1 // Сдвигаем маску влево, т.е. на каждой итерации добавляется 0 справа(младший байт).
		fmt.Printf("mask_bytes: %08b \n", getBytesInt(mask))
		fmt.Printf("num[%d]: %08b \n", num, getBytesInt(num))
		fmt.Printf("num & mask:%d %08b \n", num&mask, getBytesInt(num&mask))
	}

	println()
	noMask := ^mask                                                // Если xor наше число, то получится обратное значение, где вместо 1 нули, а вместо 0 единицы.
	fmt.Printf("noMask[%d]: %08b \n", noMask, getBytesInt(noMask)) // noMask[7]: [00000111 00000000 00000000 00000000 00000000 00000000 00000000 00000000] <- т.е. вместо 0 в маске поставим 1, а вместо 1 поставим 0

	println()
	result := noMask ^ num // 111 ^ 101 = 010 - т.е. получим комплементарное число
	fmt.Printf("r[%d]: %08b \n", result, getBytesInt(result))
}

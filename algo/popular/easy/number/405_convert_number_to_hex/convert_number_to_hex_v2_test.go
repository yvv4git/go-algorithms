package _05_convert_number_to_hex

import (
	"fmt"
	"testing"
)

func Test_toHexV2(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				num: 26,
			},
			want: "1a",
		},
		{
			name: "CASE-2",
			args: args{
				num: -1,
			},
			want: "ffffffff",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHexV2(tt.args.num); got != tt.want {
				t.Errorf("toHexV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -----
func TestResearchMethodV2Ex01(t *testing.T) {
	num := 26
	if num == 0 {
		t.Fatal()
	}

	var i int
	bytes := []byte("0123456789abcdef        ") // 8 spaces

	for i = 23; num != 0; i-- {
		bytes[i] = bytes[num&15]
		fmt.Printf("[%d] num(%d %08b) & 15(%08b) = %d(%08b) bytes[i]=%v(%v) \n", i, num, num, 15, num&15, num&15, bytes[i], string(bytes[i]))
		fmt.Printf("  num(%d %08b)>>4 = %d(%08b) \n", num, num, uint32(num)>>4, uint32(num)>>4)
		num = int(uint32(num) >> 4)
		fmt.Printf("num=%d(%08b) \n", num, num)
	}

	result := string(bytes[i+1:])
	fmt.Printf("result: %v \n", result)
}

func TestResearchMethodV2Ex02(t *testing.T) {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	num := 5
	idx := num & 9
	fmt.Printf("idx=%d \n", idx)
	idx = len(a) - 1 - idx
	fmt.Printf("idx=%d a[idx]=%v \n", idx, a[idx])
}

func TestResearchMethodV2Ex03(t *testing.T) {
	n := 10
	r := n % 16                                                    // Остаток от деления
	fmt.Printf("r: %d(%08b) char=%v \n", r, r, string(rune(r+55))) // r: 10(00001010)

	n = 20
	r = n % 16                                             // Остаток от деления
	fmt.Printf("r: %d(%08b) char=%v \n", r, r, rune(r+55)) // r: 4(00000100) char=4

	n = 20
	r = n % 16                                             // Остаток от деления
	fmt.Printf("r: %d(%08b) char=%v \n", r, r, rune(r+55)) // r: 4(00000100) char=4

	n = n / 16
	fmt.Printf("n: %d(%08b) \n", n, n)
}

func TestResearchMethodV2Ex04(t *testing.T) {
	n := 26
	res := make([]rune, 2)
	i := 0
	for n > 0 {
		tmp := 0
		tmp = n % 16
		fmt.Printf("n: %d(%08b) \n", n, n)
		fmt.Printf("tmp: %d(%08b) \n", tmp, tmp)

		if tmp < 10 { // Цифры
			fmt.Printf("i[1] = %d \n", i)
			res[i] = rune(tmp + 48)
		} else { // Буквы
			fmt.Printf("i[2] = %d \n", i)
			res[i] = rune(tmp + 55)
		}
		i++

		n = n / 16 // Сколько раз 16 влезет в n
		fmt.Printf("- n: %d(%08b) \n", n, n)
	}
	fmt.Printf("res: %v (%#v) \n", res, string(res))
}

func TestResearchMethodV2Ex05(t *testing.T) {
	/*
		Модульная арифметика имеет дело с циклическими структурами. Потому остаток от деления и деление по модулю - это разные вещи.
	*/
	converter := func(n int) string {
		bufferLen := 8
		buffer := make([]rune, bufferLen)
		i := 0

		for n > 0 {
			tmpNum := n % 16 // Остаток от деления
			fmt.Printf("[%d] tmpNum=%d(%08b) \n", i, tmpNum, tmpNum)
			if tmpNum < 10 {
				buffer[bufferLen-1-i] = rune(tmpNum + 48) // Цифры
			} else {
				buffer[bufferLen-1-i] = rune(tmpNum + 55) // Буквы
			}
			i++
			n = n / 16 // Получаем целое от деления, т.е. как бы вычитаем 16 из n.
			fmt.Printf("[%d] n=%d(%08b) \n", i, n, n)
		}

		return string(buffer[bufferLen-i:])
	}

	result := converter(26)
	t.Logf("Result: %s", result)
}

func TestResearchMethodV2Ex06(t *testing.T) {
	x := 1
	y := 16
	r := x % y
	t.Logf("x=%d(%08b) y=%d(%08b) = %d(%08b)", x, x, y, y, r, r) // x=1(00000001) y=16(00010000) = 1(00000001)

	x = 16
	y = 16
	r = x % y
	t.Logf("x=%d(%08b) y=%d(%08b) = %d(%08b)", x, x, y, y, r, r) // x=16(00010000) y=16(00010000) = 0(00000000)

	x = 2
	y = 16
	r = x % y
	t.Logf("x=%d(%08b) y=%d(%08b) = %d(%08b)", x, x, y, y, r, r) // x=2(00000010) y=16(00010000) = 2(00000010)

	x = 10
	y = 2
	r = x % y
	t.Logf("x=%d(%08b) y=%d(%08b) = %d(%08b)", x, x, y, y, r, r) // x=10(00001010) y=2(00000010) = 0(00000000)

	x = 19
	y = -12
	r = x % y
	t.Logf("x=%d(%08b) y=%d(%08b) = %d(%08b)", x, x, y, y, r, r) // x=19(00010011) y=-12(-0001100) = 7(00000111)
}

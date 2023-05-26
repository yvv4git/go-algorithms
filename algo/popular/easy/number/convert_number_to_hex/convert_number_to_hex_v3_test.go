package convert_number_to_hex

import (
	"fmt"
	"testing"
)

func Test_toHexV3(t *testing.T) {
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
			if got := toHexV3(tt.args.num); got != tt.want {
				t.Errorf("toHexV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func TestResearchMethodV3Ex01(t *testing.T) {
	convertToHex := func(num int) string {
		arr := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

		if num == 0 {
			return "0"
		}

		//if num < 0 {
		//	num = -num
		//	num ^= 0xffffffff
		//	num += 1
		//}

		result := ""

		for num > 0 {
			digit := num % 16
			result = string(arr[digit]) + result
			num /= 16
		}

		return result
	}

	t.Logf("Result: %v", convertToHex(26))
}

func TestResearchMethodV3Ex02(t *testing.T) {
	num := -1

	if num < 0 {
		num = -num
		num ^= 0xffffffff
		num += 1
	}

	fmt.Printf("num: %d(%08b)\n", num, num)

	num = -1
	num = -num
	fmt.Printf("num: %d(%08b)\n", num, num) // num: 1(00000001)

	num = -5
	num = -num
	fmt.Printf("num: %d(%08b)\n", num, num) // num: 5(00000101)
	num ^= 0xffffffff
	fmt.Printf("num: %d(%08b)\n", num, num)
	num += 1
	fmt.Printf("num: %d(%08b)\n", num, num)
}

func TestResearchMethodV3Ex03(t *testing.T) {
	convert10ToHex := func(n int) string {
		arr := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

		if n == 0 {
			return "0"
		}

		var result string
		for n > 0 {
			tmpNum := n % 16 // Определяем последний символ, как бы отрезаем с конца
			fmt.Printf("tmpNum=%d(%08b) n=%d(%08b) \n", tmpNum, tmpNum, n, n)
			result = string(arr[tmpNum]) + result
			n /= 16 // Уменьшаем число
		}

		return result
	}

	t.Logf("Result: %s ", convert10ToHex(26))
}

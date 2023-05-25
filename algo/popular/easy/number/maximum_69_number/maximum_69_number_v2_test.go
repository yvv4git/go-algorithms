package maximum_69_number

import (
	"fmt"
	"testing"
)

func Test_maximum69NumberV2(t *testing.T) {
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
				num: 9669,
			},
			want: 9969,
			desc: `Explanation: 
					Changing the first digit results in 6669.
					Changing the second digit results in 9969.
					Changing the third digit results in 9699.
					Changing the fourth digit results in 9666.
					The maximum number is 9969.`,
		},
		{
			name: "CASE-2",
			args: args{
				num: 9996,
			},
			want: 9999,
			desc: `Explanation: Changing the last digit 6 to 9 results in the maximum number.`,
		},
		{
			name: "CASE-3",
			args: args{
				num: 9999,
			},
			want: 9999,
			desc: `Explanation: It is better not to apply any change.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69NumberV2(tt.args.num); got != tt.want {
				t.Errorf("maximum69NumberV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ------------------------------
func TestResearchMathMethod01(t *testing.T) {
	/*
		Есть 2 манипуляции с числами:
		1. Деление на 10, 100, 1000 ... etc, позволяет получить самое левый, старший разряд числа.
		Например, 9669 / 1000 = 9.
		Или 669 / 100 = 6.
		Нужно лишь, чтобы делитель(div) был 10^n, где n - количество разрядов в числе.

		2. Получение остатка от деления.
		Можно отрезать левый, старший разряд от числа, если получить остаток от деления по модулю. При чем на 10^n, где n - количество разрядов в числе.
		Например, 9669 % 1000 = 669 - отрезали старший разряд.

	*/
	num := 9669
	div := 1000 // power of 10
	n := num

	for div > 0 {
		if n/div == 6 {
			/*
				Допустим n=9669 div=1000 9669/1000 =
			*/
			num += 3 * div // replace 6 with 9
			break
		} else {
			n = n % div    // move to next digit
			div = div / 10 // lower power of 10
		}
	}

	fmt.Printf("num: %d \n", num) // num: 9969
}

func TestResearchMathMethod02(t *testing.T) {
	num := 9669
	div := 1000
	n := num

	fmt.Printf("n/div == 6 ? %v \n", n/div == 6)
}

// 3*10^n
func TestResearchMathMethod03(t *testing.T) {
	//num := 69
	r := 3*10 ^ 1000
	t.Logf("r=%d", r)
}

func TestResearchMathMethod04(t *testing.T) {
	num := 9669
	n := num
	div := 1000

	n = n % div
	fmt.Printf("n: %d \n", n)

	n = 9669
	div = 1000
	fmt.Printf("%d / %d => %v \n", n, div, n/div) // 9669 / 1000 => 9

	n = 669
	div = 100
	fmt.Printf("%d / %d => %v \n", n, div, n/div) // 669 / 100 => 6

	n = 9669
	div = 1000
	fmt.Printf("%d %% %d = %d \n", n, div, n%div) // Находим остаток от деления => 9669 % 1000 = 669
}

func TestResearchMathMethod05(t *testing.T) {
	num := 9669
	n := num
	div := 1000 // 10^4 - берем это из constraints

	fmt.Printf("n(before): %d \n", n)
	for div > 0 {
		highBit := n / div
		fmt.Printf("highBit: %d \n", highBit)
		if highBit == 6 {
			/*
				Нашли первую шестерку слева. По условию задачи, надо найти максимальное число, в котором поменяли 1 цифру 6.
				Логично, что надо менять самую левую 6 на 9. Тогда результат будет больше.
				Чтобы 9669 получить 9969, надо к числу 9669 прибавить 300, тогда 6 во втором разряде поменяется на 9.
			*/
			num += 3 * div
			break
		} else { // Если старший бит равен 9, то надо отрезать эту часть числа.
			n = n % div    // Отрезаем старший бит
			div = div / 10 // Уменьшаем делитель в 10 раз
		}
	}

	fmt.Printf("num: %d \n", num)

}

package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: Positive numbers",
			args: args{a: 1, b: 2},
			want: 3,
		},
		{
			name: "Test Case 2: Negative numbers",
			args: args{a: -1, b: -2},
			want: -3,
		},
		{
			name: "Test Case 3: Mixed numbers",
			args: args{a: 1, b: -2},
			want: -1,
		},
		{
			name: "Test Case 4: Zero",
			args: args{a: 0, b: 0},
			want: 0,
		},
		{
			name: "Test Case 5: Large numbers",
			args: args{a: 1000000, b: 2000000},
			want: 3000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResearch01(t *testing.T) {
	a := 1 // 0000000000000000000000000000000000000000000000000000000000000001
	b := 5 // 0000000000000000000000000000000000000000000000000000000000000101

	fmt.Printf("a: %d [%s] \n", a, getBinaryString(a))
	fmt.Printf("b: %d [%s] \n", b, getBinaryString(b))

	c := a ^ b // 0000000000000000000000000000000000000000000000000000000000000100
	fmt.Printf("c: %d [%s] \n", c, getBinaryString(c))

	carry := a & b // 0000000000000000000000000000000000000000000000000000000000000001
	fmt.Printf("carry: %d [%s] \n", carry, getBinaryString(carry))

	carryShifted := carry << 1 // 0000000000000000000000000000000000000000000000000000000000000010
	fmt.Printf("carryShifted: %d [%s] \n", carryShifted, getBinaryString(carryShifted))
}

// Функция для получения строки с бинарным представлением числа.
func getBinaryString(n int) string {
	// Создаем пустую строку, в которую будем добавлять биты.
	var binary string

	// Проходим по каждому биту числа.
	for i := 63; i >= 0; i-- {
		// Получаем значение i-го бита числа.
		bit := (n >> i) & 1
		// Добавляем бит в строку.
		binary += strconv.Itoa(int(bit))
	}

	// Возвращаем полученную строку.
	return binary
}

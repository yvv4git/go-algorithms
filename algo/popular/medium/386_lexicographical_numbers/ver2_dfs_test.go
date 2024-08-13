package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_lexicographicalNumbersV2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1",
			args: args{n: 1},
			want: []int{1},
		},
		{
			name: "Test Case 2",
			args: args{n: 2},
			want: []int{1, 2},
		},
		{
			name: "Test Case 3",
			args: args{n: 10},
			want: []int{1, 10, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Test Case 4",
			args: args{n: 13},
			want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "Test Case 5",
			args: args{n: 20},
			want: []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 2, 20, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lexicographicalNumbersV2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lexicographicalNumbersV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func dfsR(current, n int, result *[]int) {
	fmt.Printf("current %d, n %d, res=%v\n", current, n, result)
	if current > n {
		return
	}
	*result = append(*result, current)

	// Этот вызов позволяет нам исследовать числа, которые начинаются с текущего числа и имеют дополнительные разряды.
	// Например, если current равно 1, то следующим числом будет 10, затем 100 и так далее.
	// Это необходимо для генерации чисел в лексикографическом порядке, где числа с большим количеством разрядов должны идти перед числами с меньшим количеством разрядов.
	dfsR(current*10, n, result) // Рекурсивно генерируем числа, умножая текущее число на 10

	// Это условие проверяет, не является ли последняя цифра текущего числа равной 9.
	// Если последняя цифра равна 9, то увеличение на 1 приведет к переходу к следующему разряду (например, 19 -> 20),
	// что нарушит лексикографический порядок.
	if current%10 != 9 {
		// Этот вызов позволяет нам исследовать следующие числа в лексикографическом порядке, увеличивая последнюю цифру на 1.
		// Например, если current равно 1, то следующим числом будет 2.
		// Это необходимо для последовательного перехода к следующим числам в лексикографическом порядке.
		dfsR(current+1, n, result) // Рекурсивно генерируем следующее число, увеличивая последнюю цифру на 1
	}
}

func TestResearchDFS(t *testing.T) {
	var result []int
	dfsR(1, 10, &result)
}

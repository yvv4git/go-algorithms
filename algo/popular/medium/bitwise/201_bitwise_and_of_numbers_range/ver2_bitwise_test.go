package main

import (
	"fmt"
	"testing"
)

func Test_rangeBitwiseAndV2(t *testing.T) {
	type args struct {
		left  int
		right int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				left:  5,
				right: 7,
			},
			want: 4,
		},
		{
			name: "Test Case 2",
			args: args{
				left:  0,
				right: 0,
			},
			want: 0,
		},
		{
			name: "Test Case 3",
			args: args{
				left:  1,
				right: 2147483647,
			},
			want: 0,
		},
		{
			name: "Test Case 4",
			args: args{
				left:  10,
				right: 15,
			},
			want: 8,
		},
		{
			name: "Test Case 5",
			args: args{
				left:  3,
				right: 3,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAndV2(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeBitwiseAndV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// printBinary представляет собой функцию, которая выводит любую переменную в бинарном формате.
func printBinary(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Printf("%032b\n", uint32(v))
	case int8:
		fmt.Printf("%08b\n", uint8(v))
	case int16:
		fmt.Printf("%016b\n", uint16(v))
	case int32:
		fmt.Printf("%032b\n", uint32(v))
	case int64:
		fmt.Printf("%064b\n", uint64(v))
	case uint:
		fmt.Printf("%032b\n", v)
	case uint8:
		fmt.Printf("%08b\n", v)
	case uint16:
		fmt.Printf("%016b\n", v)
	case uint32:
		fmt.Printf("%032b\n", v)
	case uint64:
		fmt.Printf("%064b\n", v)
	case string:
		for _, c := range v {
			fmt.Printf("%08b ", c)
		}
		fmt.Println()
	default:
		fmt.Println("Unsupported type")
	}
}

func TestResearch01(t *testing.T) {
	left := 3
	right := 5
	result := left & right
	t.Logf("Res: %v", result)
	printBinary(left)
	printBinary(right)
	printBinary(result)
}

func TestResearch02(t *testing.T) {
	for i := 0; i < 10; i++ {
		printBinary(i)
	}
}

// Функция для преобразования целого числа в двоичную строку без обрезания ведущих нулей.
func intToBinaryString(n interface{}) string {
	switch v := n.(type) {
	case int:
		return fmt.Sprintf("%032b", uint32(v))
	case int8:
		return fmt.Sprintf("%08b", uint8(v))
	case int16:
		return fmt.Sprintf("%016b", uint16(v))
	case int32:
		return fmt.Sprintf("%032b", uint32(v))
	case int64:
		return fmt.Sprintf("%064b", uint64(v))
	case uint:
		return fmt.Sprintf("%032b", v)
	case uint8:
		return fmt.Sprintf("%08b", v)
	case uint16:
		return fmt.Sprintf("%016b", v)
	case uint32:
		return fmt.Sprintf("%032b", v)
	case uint64:
		return fmt.Sprintf("%064b", v)
	default:
		return "Unsupported type"
	}
}

func TestResearch03(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Logf("[%d] => %s", i, intToBinaryString(i))
	}
}

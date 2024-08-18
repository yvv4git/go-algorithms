package main

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_validUtf8(t *testing.T) {
	type args struct {
		data []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid single byte",
			args: args{data: []int{0x41}}, // 'A' in UTF-8
			want: true,
		},
		{
			name: "Valid two bytes",
			args: args{data: []int{0xC2, 0xA9}}, // '©' in UTF-8
			want: true,
		},
		{
			name: "Valid three bytes",
			args: args{data: []int{0xE2, 0x82, 0xAC}}, // '€' in UTF-8
			want: true,
		},
		{
			name: "Valid four bytes",
			args: args{data: []int{0xF0, 0x9F, 0x98, 0x81}}, // '😁' in UTF-8
			want: true,
		},
		{
			name: "Invalid continuation byte",
			args: args{data: []int{0x80}}, // Invalid continuation byte
			want: false,
		},
		{
			name: "Invalid start byte",
			args: args{data: []int{0xF8}}, // Invalid start byte (5 bytes)
			want: false,
		},
		{
			name: "Incomplete sequence",
			args: args{data: []int{0xE2, 0x82}}, // Incomplete three-byte sequence
			want: false,
		},
		{
			name: "Mixed valid and invalid",
			args: args{data: []int{0xC2, 0xA9, 0x80}}, // Valid '©' followed by invalid continuation byte
			want: false,
		},
		{
			name: "Empty array",
			args: args{data: []int{}}, // Empty array
			want: true,
		},
		{
			name: "Multiple valid characters",
			args: args{data: []int{0x41, 0xC2, 0xA9, 0xE2, 0x82, 0xAC, 0xF0, 0x9F, 0x98, 0x81}}, // 'A', '©', '€', '😁'
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.args.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}

// PrettyPrintBinary принимает целое число и возвращает его бинарное представление в виде строки
func PrettyPrintBinary(value int) string {
	// Преобразуем число в бинарную строку
	binaryString := strconv.FormatInt(int64(value), 2)

	// Дополняем строку до длины, кратной 8
	for len(binaryString)%8 != 0 {
		binaryString = "0" + binaryString
	}

	return binaryString
}

func TestResearch01(t *testing.T) {
	b := 0b00000000
	d := 0b10000000
	fmt.Printf("%s & %s = %s\n", PrettyPrintBinary(b), PrettyPrintBinary(d), PrettyPrintBinary(b&0b10000000))
}

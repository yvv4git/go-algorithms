package main

import "testing"

func Test_isPalindromeStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				str: "121",
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				str: "123454321",
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				str: "1234554321",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeStr(tt.args.str); got != tt.want {
				t.Errorf("isPalindromeStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Сложность алгоритма O(n), ибо приходится
// линейно идти по массиву/строке.
func isPalindromeStr(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}

	return true
}

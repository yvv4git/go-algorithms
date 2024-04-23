package main

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Тест 1: Обычная строка",
			args: args{s: "Hello World"},
			want: "olleH dlroW",
		},
		{
			name: "Тест 2: Строка с несколькими пробелами между словами",
			args: args{s: "The quick brown fox jumps over the lazy dog"},
			want: "ehT kciuq nworb xof spmuj revo eht yzal god",
		},
		{
			name: "Тест 3: Строка с несколькими пробелами в конце",
			args: args{s: "Let's take LeetCode contest  "},
			want: "s'teL ekat edoCteeL tsetnoc  ",
		},
		{
			name: "Тест 5: Пустая строка",
			args: args{s: ""},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

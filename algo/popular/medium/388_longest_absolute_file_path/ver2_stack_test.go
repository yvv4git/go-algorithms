package main

import "testing"

func Test_lengthLongestPathV2(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// Тестовый кейс 1: Простой случай с одним файлом
		{
			name: "Single file",
			args: args{
				input: "file.txt",
			},
			want: 8,
		},
		// Тестовый кейс 2: Одна директория с файлом
		{
			name: "Single directory with file",
			args: args{
				input: "dir\n\tfile.txt",
			},
			want: 12,
		},
		// Тестовый кейс 3: Несколько уровней вложенности
		{
			name: "Multiple levels of nesting",
			args: args{
				input: "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.txt",
			},
			want: 20,
		},
		// Тестовый кейс 4: Пустая строка
		{
			name: "Empty string",
			args: args{
				input: "",
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthLongestPathV2(tt.args.input); got != tt.want {
				t.Errorf("lengthLongestPathV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

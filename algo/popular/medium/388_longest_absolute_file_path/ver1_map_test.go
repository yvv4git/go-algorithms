package main

import "testing"

func Test_lengthLongestPath(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		// Тестовый кейс 1: Простой случай с одним файлом
		{
			name:     "Single file",
			input:    "file.txt",
			expected: 8,
		},
		// Тестовый кейс 2: Одна директория с файлом
		{
			name:     "Single directory with file",
			input:    "dir\n\tfile.txt",
			expected: 12,
		},
		// Тестовый кейс 3: Несколько уровней вложенности
		{
			name:     "Multiple levels of nesting",
			input:    "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.txt",
			expected: 20,
		},
		// Тестовый кейс 4: Пустая строка
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lengthLongestPathV2(test.input)
			if result != test.expected {
				t.Errorf("Test '%s' failed for input '%s'. Expected: %d, Got: %d", test.name, test.input, test.expected, result)
			}
		})
	}
}

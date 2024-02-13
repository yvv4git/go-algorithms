package main

import "testing"

func Test_percentageLetter(t *testing.T) {
	type args struct {
		s      string
		letter byte
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1: letter is present in the string",
			args: args{
				s:      "hello",
				letter: 'l',
			},
			want: 40,
		},
		{
			name: "Test case 2: letter is not present in the string",
			args: args{
				s:      "world",
				letter: 'z',
			},
			want: 0,
		},
		{
			name: "Test case 4: string is empty",
			args: args{
				s:      "",
				letter: 'a',
			},
			want: 0,
		},
		{
			name: "Test case 5: letter is uppercase and present in the string",
			args: args{
				s:      "Hello",
				letter: 'H',
			},
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := percentageLetter(tt.args.s, tt.args.letter); got != tt.want {
				t.Errorf("percentageLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}

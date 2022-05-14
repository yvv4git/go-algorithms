package main

import (
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				s: []byte{
					'h', 'e', 'l', 'l', 'o',
				},
			},
			want: "olleh",
		},
		{
			name: "CASE-2",
			args: args{
				s: []byte{
					'H', 'a', 'n', 'n', 'a', 'h',
				},
			},
			want: "hannaH",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)

			if string(tt.args.s) != tt.want {
				t.Errorf("We want %s %s", tt.want, tt.args.s)
			}
		})
	}
}

func TestRuString(t *testing.T) {
	s := "Владимир"
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		t.Logf("%c", r[i])
	}
}

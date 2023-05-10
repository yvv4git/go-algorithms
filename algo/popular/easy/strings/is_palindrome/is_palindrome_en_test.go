package is_palindrome

import (
	"testing"
)

func Test_isPalindromeEn(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name    string
		args    args
		wantOut bool
	}{
		{
			name: "CASE-1",
			args: args{
				input: "lolol",
			},
			wantOut: true,
		},
		{
			name: "CASE-2",
			args: args{
				input: "Anonymous",
			},
			wantOut: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := isPalindromeEn(tt.args.input); gotOut != tt.wantOut {
				t.Errorf("isPalindromeEn() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

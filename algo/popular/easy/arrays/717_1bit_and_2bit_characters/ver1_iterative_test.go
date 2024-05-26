package main

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	type args struct {
		bits []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case with valid 1-bit character at the end",
			args: args{bits: []int{1, 0, 0}},
			want: true,
		},
		{
			name: "Test case with only one bit",
			args: args{bits: []int{0}},
			want: true,
		},
		{
			name: "Test case with all 1-bit characters",
			args: args{bits: []int{1, 0, 1, 0}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneBitCharacter(tt.args.bits); got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}

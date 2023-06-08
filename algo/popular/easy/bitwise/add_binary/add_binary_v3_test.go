package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinaryV3(t *testing.T) {
	type args struct {
		a string
		b string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "CASE-2",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
		{
			name: "CASE-3",
			args: args{
				a: "",
				b: "",
			},
			want: "",
		},
		{
			name: "CASE-4",
			args: args{
				a: "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
				b: "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			},
			want: "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
		{
			name: "CASE-5",
			args: args{
				a: "1111",
				b: "1111",
			},
			want: "11110",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addBinaryV3(tt.args.a, tt.args.b), "addBinaryV3(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}

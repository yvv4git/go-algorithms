package main

import "testing"

func Test_myAtoiV1(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive number",
			args: args{str: "123"},
			want: 123,
		},
		{
			name: "negative number",
			args: args{str: "-456"},
			want: -456,
		},
		{
			name: "zero",
			args: args{str: "0"},
			want: 0,
		},
		{
			name: "empty string",
			args: args{str: ""},
			want: 0,
		},
		{
			name: "non-numeric string",
			args: args{str: "hello"},
			want: 0,
		},

		{
			name: "numeric string with leading zeros",
			args: args{str: "000123"},
			want: 123,
		},
		{
			name: "numeric string with trailing zeros",
			args: args{str: "123000"},
			want: 123000,
		},
		{
			name: "numeric string with leading and trailing zeros",
			args: args{str: "000123000"},
			want: 123000,
		},
		{
			name: "numeric string with leading and trailing non-numeric characters",
			args: args{str: "abc123def"},
			want: 0,
		},
		{
			name: "numeric string with leading and trailing non-numeric characters and leading and trailing whitespace",
			args: args{str: " abc123def "},
			want: 0,
		},
		{
			name: "numeric string with leading and trailing non-numeric characters and leading and trailing whitespace and leading and trailing zeros",
			args: args{str: " abc000123def000 "},
			want: 123,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoiV1(tt.args.str); got != tt.want {
				t.Errorf("myAtoiV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

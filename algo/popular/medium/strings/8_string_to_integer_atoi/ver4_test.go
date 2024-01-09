package main

import "testing"

func Test_myAtoiV4(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive number",
			args: args{s: "123"},
			want: 123,
		},
		{
			name: "negative number",
			args: args{s: "-456"},
			want: -456,
		},
		{
			name: "zero",
			args: args{s: "0"},
			want: 0,
		},
		{
			name: "empty string",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "non-numeric string",
			args: args{s: "hello"},
			want: 0,
		},

		{
			name: "numeric string with leading zeros",
			args: args{s: "000123"},
			want: 123,
		},
		{
			name: "numeric string with trailing zeros",
			args: args{s: "123000"},
			want: 123000,
		},
		{
			name: "numeric string with leading and trailing zeros",
			args: args{s: "000123000"},
			want: 123000,
		},
		{
			name: "numeric string with leading and trailing non-numeric characters",
			args: args{s: "abc123def"},
			want: 0,
		},
		{
			name: "numeric string with leading and trailing non-numeric characters and leading and trailing whitespace",
			args: args{s: " abc123def "},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoiV4(tt.args.s); got != tt.want {
				t.Errorf("myAtoiV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

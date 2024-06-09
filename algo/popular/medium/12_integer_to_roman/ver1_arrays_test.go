package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{num: 3},
			want: "III",
		},
		{
			name: "Test case 2",
			args: args{num: 4},
			want: "IV",
		},
		{
			name: "Test case 3",
			args: args{num: 9},
			want: "IX",
		},
		{
			name: "Test case 4",
			args: args{num: 58},
			want: "LVIII",
		},
		{
			name: "Test case 5",
			args: args{num: 1994},
			want: "MCMXCIV",
		},
		{
			name: "Test case 6",
			args: args{num: 3999},
			want: "MMMCMXCIX",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}

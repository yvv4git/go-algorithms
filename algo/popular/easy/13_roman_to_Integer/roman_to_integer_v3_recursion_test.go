package _3_roman_to_Integer

import "testing"

func Test_romanToIntV3(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				s: "III",
			},
			want: 3,
		},
		{
			name: "CASE-2",
			args: args{
				s: "LVIII",
			},
			want: 58,
			desc: "Explanation: L = 50, V= 5, III = 3.",
		},
		{
			name: "CASE-3",
			args: args{
				s: "MCMXCIV",
			},
			want: 1994,
			desc: "Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToIntV3(tt.args.s); got != tt.want {
				t.Errorf("romanToIntV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

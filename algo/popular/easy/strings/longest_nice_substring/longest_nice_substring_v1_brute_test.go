package longest_nice_substring

import "testing"

func Test_longestNiceSubstringV1(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				s: "YazaAay",
			},
			want: "aAa",
		},
		{
			name: "CASE-2",
			args: args{
				s: "Bb",
			},
			want: "Bb",
		},
		{
			name: "CASE-3",
			args: args{
				s: "c",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestNiceSubstringV1(tt.args.s); got != tt.want {
				t.Errorf("longestNiceSubstringV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

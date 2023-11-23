package _74_guess_number_higher_or_lower

import "testing"

func Test_guessNumberV2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 10,
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := guessNumberV2(tt.args.n); got != tt.want {
				t.Errorf("guessNumberV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

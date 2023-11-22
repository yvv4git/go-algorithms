package sum_of_digits_float

import "testing"

func Test_sumDigits(t *testing.T) {
	type args struct {
		n float64
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 679.25,
			},
			want: 29,
		},
		{
			name: "CASE-2",
			args: args{
				n: 12345,
			},
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumDigits(tt.args.n); got != tt.want {
				t.Errorf("sumDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

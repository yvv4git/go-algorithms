package _00_nth_digit

import "testing"

func Test_findNthDigitV2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{n: 11},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNthDigitV2(tt.args.n); got != tt.want {
				t.Errorf("findNthDigitV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

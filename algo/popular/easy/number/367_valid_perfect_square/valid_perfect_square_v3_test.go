package _67_valid_perfect_square

import "testing"

func Test_isPerfectSquareV3(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				num: 16,
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				num: 14,
			},
			want: false,
		},
		{
			name: "CASE-3",
			args: args{
				num: 9,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquareV3(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquareV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

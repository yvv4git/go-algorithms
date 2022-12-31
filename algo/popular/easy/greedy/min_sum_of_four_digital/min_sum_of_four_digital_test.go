package min_sum_of_four_digital

import "testing"

func TestMinimumSum(t *testing.T) {
	type args struct {
		num int
	}

	testCases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				num: 2932,
			},
			want: 52,
		},
		{
			name: "CASE-2",
			args: args{
				num: 4009,
			},
			want: 13,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumSum(tt.args.num); got != tt.want {
				t.Errorf("MinimumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

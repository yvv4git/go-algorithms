package maximum_69_number

import "testing"

func Test_maximum69NumberV1(t *testing.T) {
	type args struct {
		num int
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
				num: 9669,
			},
			want: 9969,
			desc: `Explanation: 
					Changing the first digit results in 6669.
					Changing the second digit results in 9969.
					Changing the third digit results in 9699.
					Changing the fourth digit results in 9666.
					The maximum number is 9969.`,
		},
		{
			name: "CASE-2",
			args: args{
				num: 9996,
			},
			want: 9999,
			desc: `Explanation: Changing the last digit 6 to 9 results in the maximum number.`,
		},
		{
			name: "CASE-3",
			args: args{
				num: 9999,
			},
			want: 9999,
			desc: `Explanation: It is better not to apply any change.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum69NumberV1(tt.args.num); got != tt.want {
				t.Errorf("maximum69NumberV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

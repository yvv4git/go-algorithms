package main

import "testing"

func Test_countOddsV2(t *testing.T) {
	type args struct {
		low  int
		high int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				low:  3,
				high: 7,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				low:  8,
				high: 10,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOddsV2(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countOddsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

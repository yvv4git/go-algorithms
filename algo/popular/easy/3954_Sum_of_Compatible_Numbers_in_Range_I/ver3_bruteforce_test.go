package main

import "testing"

func Test_sumOfCompatibleNumbersV3(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want int
	}{
		{
			name: "example 1",
			n:    2,
			k:    3,
			want: 10,
		},
		{
			name: "example 2",
			n:    5,
			k:    1,
			want: 0,
		},
		{
			name: "n=1 k=1",
			n:    1,
			k:    1,
			want: 2,
		},
		{
			name: "n=3 k=2",
			n:    3,
			k:    2,
			want: 4,
		},
		{
			name: "n=10 k=10",
			n:    10,
			k:    10,
			want: 63,
		},
		{
			name: "n=1 k=100",
			n:    1,
			k:    100,
			want: 2550,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfCompatibleNumbersV3(tt.n, tt.k); got != tt.want {
				t.Errorf("sumOfCompatibleNumbersV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_numJewelsInStonesV2(t *testing.T) {
	type args struct {
		jewels string
		stones string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// Тестовые кейсы
		{
			name: "Example 1",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				jewels: "z",
				stones: "ZZ",
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				jewels: "abc",
				stones: "aabbcc",
			},
			want: 6,
		},
		{
			name: "Example 4",
			args: args{
				jewels: "ABC",
				stones: "aabbcc",
			},
			want: 0,
		},
		{
			name: "Example 5",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "Example 6",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "Example 7",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "Example 8",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "Example 9",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
		{
			name: "Example 10",
			args: args{
				jewels: "aA",
				stones: "aAAbbbb",
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStonesV2(tt.args.jewels, tt.args.stones); got != tt.want {
				t.Errorf("numJewelsInStonesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

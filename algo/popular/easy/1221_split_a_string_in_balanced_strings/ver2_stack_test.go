package main

import "testing"

func Test_balancedStringSplitV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				s: "RLRRLLRLRL",
			},
			want: 4,
		},
		{
			name: "CASE-2",
			args: args{
				s: "RLLLLRRRLR",
			},
			want: 3,
		},
		{
			name: "CASE-3",
			args: args{
				s: "LLLLRRRR",
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balancedStringSplitV2(tt.args.s); got != tt.want {
				t.Errorf("balancedStringSplitV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

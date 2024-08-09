package main

import "testing"

func Test_canMeasureWaterV2(t *testing.T) {
	type args struct {
		x      int
		y      int
		target int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				x:      3,
				y:      5,
				target: 4,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				x:      2,
				y:      6,
				target: 5,
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				x:      1,
				y:      2,
				target: 3,
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				x:      0,
				y:      0,
				target: 0,
			},
			want: true,
		},
		{
			name: "Example 5",
			args: args{
				x:      0,
				y:      0,
				target: 1,
			},
			want: false,
		},
		{
			name: "Example 6",
			args: args{
				x:      4,
				y:      6,
				target: 8,
			},
			want: true,
		},
		{
			name: "Example 7",
			args: args{
				x:      4,
				y:      6,
				target: 10,
			},
			want: true,
		},
		{
			name: "Example 8",
			args: args{
				x:      4,
				y:      6,
				target: 11,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWaterV2(tt.args.x, tt.args.y, tt.args.target); got != tt.want {
				t.Errorf("canMeasureWaterV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

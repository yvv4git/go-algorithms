package main

import "testing"

func Test_canMeasureWaterV3(t *testing.T) {
	type args struct {
		jug1Capacity   int
		jug2Capacity   int
		targetCapacity int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				jug1Capacity:   3,
				jug2Capacity:   5,
				targetCapacity: 4,
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				jug1Capacity:   2,
				jug2Capacity:   6,
				targetCapacity: 5,
			},
			want: false,
		},
		{
			name: "Example 3",
			args: args{
				jug1Capacity:   1,
				jug2Capacity:   2,
				targetCapacity: 3,
			},
			want: true,
		},
		{
			name: "Example 4",
			args: args{
				jug1Capacity:   0,
				jug2Capacity:   0,
				targetCapacity: 0,
			},
			want: true,
		},
		{
			name: "Example 5",
			args: args{
				jug1Capacity:   0,
				jug2Capacity:   0,
				targetCapacity: 1,
			},
			want: false,
		},
		{
			name: "Example 6",
			args: args{
				jug1Capacity:   4,
				jug2Capacity:   6,
				targetCapacity: 8,
			},
			want: true,
		},
		{
			name: "Example 7",
			args: args{
				jug1Capacity:   4,
				jug2Capacity:   6,
				targetCapacity: 10,
			},
			want: true,
		},
		{
			name: "Example 8",
			args: args{
				jug1Capacity:   4,
				jug2Capacity:   6,
				targetCapacity: 11,
			},
			want: false,
		},
		{
			name: "Example 9",
			args: args{
				jug1Capacity:   3,
				jug2Capacity:   7,
				targetCapacity: 5,
			},
			want: true,
		},
		{
			name: "Example 10",
			args: args{
				jug1Capacity:   5,
				jug2Capacity:   8,
				targetCapacity: 2,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMeasureWaterV3(tt.args.jug1Capacity, tt.args.jug2Capacity, tt.args.targetCapacity); got != tt.want {
				t.Errorf("canMeasureWaterV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

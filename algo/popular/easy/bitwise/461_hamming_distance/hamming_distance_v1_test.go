package _61_hamming_distance

import "testing"

func Test_hammingDistanceV1(t *testing.T) {
	type args struct {
		x int
		y int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				x: 1,
				y: 4,
			},
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				x: 3,
				y: 1,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistanceV1(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistanceV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

package sqrt

import "testing"

func Test_mySqrtV2(t *testing.T) {
	type args struct {
		x int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				x: 4,
			},
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				x: 8,
			},
			want: 2, // The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrtV2(tt.args.x); got != tt.want {
				t.Errorf("mySqrtV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

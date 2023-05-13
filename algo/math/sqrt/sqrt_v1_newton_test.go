package sqrt

import "testing"

func TestSqrtV1(t *testing.T) {
	type args struct {
		x float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "CASE-1",
			args: args{
				x: 4.0,
			},
			want: 2.0,
		},
		{
			name: "CASE-2",
			args: args{
				x: 25.0,
			},
			want: 5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sqrtV1(tt.args.x); got != tt.want {
				t.Errorf("sqrtV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

package multition_karatsuba

import "testing"

func Test_karatsuba(t *testing.T) {
	type args struct {
		x int64
		y int64
	}

	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "CASE-1",
			args: args{
				x: 1234,
				y: 5678,
			},
			want: 7006652,
		},
		{
			name: "CASE-2",
			args: args{
				x: 1234,
				y: -5678,
			},
			want: -7006652,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := karatsuba(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("karatsuba() = %v, want %v", got, tt.want)
			}
		})
	}
}

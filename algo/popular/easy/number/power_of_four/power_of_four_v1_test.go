package power_of_four

import "testing"

func Test_isPowerOfFourV1(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				n: 5,
			},
			want: false,
		},
		{
			name: "CASE-3",
			args: args{
				n: 1,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFourV1(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfFourV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

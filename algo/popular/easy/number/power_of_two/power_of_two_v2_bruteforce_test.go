package power_of_two

import "testing"

func Test_isPowerOfTwoV2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				n: 1,
			},
			want: true,
			desc: "2^0 = 1",
		},
		{
			name: "CASE-2",
			args: args{
				n: 16,
			},
			want: true,
			desc: "2^4 = 16",
		},
		{
			name: "CASE-3",
			args: args{
				n: 3,
			},
			want: false,
			desc: "No found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwoV2(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfTwoV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

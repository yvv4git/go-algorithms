package power_of_three

import "testing"

func Test_isPowerOfThreeV3(t *testing.T) {
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
				n: 27,
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				n: 0,
			},
			want: false,
		},
		{
			name: "CASE-3",
			args: args{
				n: -1,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThreeV3(tt.args.n); got != tt.want {
				t.Errorf("isPowerOfThreeV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

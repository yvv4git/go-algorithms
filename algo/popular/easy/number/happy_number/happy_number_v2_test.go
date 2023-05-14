package happy_number

import "testing"

func Test_isHappyV2(t *testing.T) {
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
				n: 19,
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				n: 2,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappyV2(tt.args.n); got != tt.want {
				t.Errorf("isHappyV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

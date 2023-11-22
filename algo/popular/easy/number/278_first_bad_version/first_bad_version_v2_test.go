package _78_first_bad_version

import "testing"

func Test_firstBadVersionV2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n: 5,
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersionV2(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersionV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

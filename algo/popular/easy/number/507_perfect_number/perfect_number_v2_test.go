package _07_perfect_number

import "testing"

func Test_checkPerfectNumberV2(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				num: 6,
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				num: 28,
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				num: 7,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPerfectNumberV2(tt.args.num); got != tt.want {
				t.Errorf("checkPerfectNumberV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package number_of_common_factors

import "testing"

func Test_commonFactorsV1(t *testing.T) {
	type args struct {
		a int
		b int
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				a: 12,
				b: 6,
			},
			want: 4,
			desc: "The common factors of 12 and 6 are 1, 2, 3, 6.",
		},
		{
			name: "CASE-2",
			args: args{
				a: 25,
				b: 30,
			},
			want: 2,
			desc: "The common factors of 25 and 30 are 1, 5.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commonFactorsV1(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("commonFactorsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

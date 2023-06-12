package complement_of_base_10_int

import "testing"

func Test_bitwiseComplementV1(t *testing.T) {
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
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				n: 7,
			},
			want: 0,
		},
		{
			name: "CASE-3",
			args: args{
				n: 10,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitwiseComplementV1(tt.args.n); got != tt.want {
				t.Errorf("bitwiseComplementV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

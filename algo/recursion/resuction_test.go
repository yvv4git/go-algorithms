package recursion

import "testing"

func Test_rec(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "CASE-1",
			args: args{
				n: 5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			rec(tt.args.n)
		})
	}
}

package gcb

import "testing"

func Test_gcbEuclidean(t *testing.T) {
	type args struct {
		m int
		n int
	}

	tests := []struct {
		name        string
		args        args
		want        int
		description string
	}{
		{
			name: "CASE-1",
			args: args{
				m: 10,
				n: 5,
			},
			want:        5,
			description: "НОД 10 и 5 = 5",
		},
		{
			name: "CASE-2",
			args: args{
				m: 100,
				n: 5,
			},
			want:        5,
			description: "НОД 100 и 5 = 5",
		},
		{
			name: "CASE-3",
			args: args{
				m: 20,
				n: 6,
			},
			want:        2,
			description: "НОД 20 и 6 = 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcbEuclidean(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("gcb() = %v, want %v", got, tt.want)
			}
		})
	}
}

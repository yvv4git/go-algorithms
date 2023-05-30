package placements

import "testing"

func TestPlacementsCnt(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case-1",
			args: args{
				n: 3,
				m: 2,
			},
			want: 6,
		},
		{
			name: "case-2",
			args: args{
				n: 3,
				m: 3,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Placements(tt.args.n, tt.args.m); got != tt.want {
				t.Errorf("Placements() = %v, want %v", got, tt.want)
			}
		})
	}
}

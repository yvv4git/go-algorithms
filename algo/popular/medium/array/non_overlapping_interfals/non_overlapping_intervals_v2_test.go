package non_overlapping_interfals

import "testing"

func Test_eraseOverlapIntervalsV2(t *testing.T) {
	type args struct {
		intervals [][]int
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
				intervals: [][]int{
					{1, 2}, {2, 3}, {3, 4}, {1, 3},
				},
			},
			want: 1,
			desc: `[1,3] can be removed and the rest of the intervals are non-overlapping.`,
		},
		{
			name: "CASE-2",
			args: args{
				intervals: [][]int{
					{1, 2}, {1, 2}, {1, 2},
				},
			},
			want: 2,
			desc: `You need to remove two [1,2] to make the rest of the intervals non-overlapping.`,
		},
		{
			name: "CASE-3",
			args: args{
				intervals: [][]int{
					{1, 2}, {2, 3},
				},
			},
			want: 0,
			desc: `You don't need to remove any of the intervals since they're already non-overlapping.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervalsV2(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervalsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

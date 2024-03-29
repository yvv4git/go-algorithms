package selection

import "testing"

func TestSelectSort(t *testing.T) {
	type args struct {
		ar []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				ar: []int{2, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-2",
			args: args{
				ar: []int{2, 1, 4, 3, 6, 5},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(_ *testing.T) {
			SelectSort(tt.args.ar)
		})
	}
}

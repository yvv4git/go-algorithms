package sorts

import "testing"

func TestShellSort(t *testing.T) {
	type args struct {
		ar []int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "CASE-1",
			args: args{
				ar: []int{2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ShellSort(tt.args.ar)
		})
	}
}

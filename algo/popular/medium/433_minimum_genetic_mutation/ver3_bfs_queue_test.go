package main

import "testing"

func Test_minMutationV3(t *testing.T) {
	type args struct {
		startGene string
		endGene   string
		bank      []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case with valid mutation",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AACCGGTA",
				bank:      []string{"AACCGGTA"},
			},
			want: 1,
		},
		{
			name: "Test case with multiple valid mutations",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AAACGGTA",
				bank:      []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			},
			want: 2,
		},
		{
			name: "Test case with no valid mutation path",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AACCGGTA",
				bank:      []string{"AACCGCTA"},
			},
			want: -1,
		},
		//{
		//	name: "Test case with end sequence being the start sequence",
		//	args: args{
		//		startGene: "AACCGGTA",
		//		endGene:   "AACCGGTA",
		//		bank:      []string{"AACCGGTA"},
		//	},
		//	want: 0,
		//},
		{
			name: "Test case with no mutation bank",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AACCGGTA",
				bank:      []string{},
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutationV3(tt.args.startGene, tt.args.endGene, tt.args.bank); got != tt.want {
				t.Errorf("minMutationV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

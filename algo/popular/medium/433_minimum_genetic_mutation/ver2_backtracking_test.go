package main

import "testing"

func Test_minMutationV2(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case with valid mutation",
			args: args{
				start: "AACCGGTT",
				end:   "AACCGGTA",
				bank:  []string{"AACCGGTA"},
			},
			want: 1,
		},
		{
			name: "Test case with multiple valid mutations",
			args: args{
				start: "AACCGGTT",
				end:   "AAACGGTA",
				bank:  []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			},
			want: 2,
		},
		{
			name: "Test case with no valid mutation path",
			args: args{
				start: "AACCGGTT",
				end:   "AACCGGTA",
				bank:  []string{"AACCGCTA"},
			},
			want: -1,
		},
		//{
		//	name: "Test case with end sequence being the start sequence",
		//	args: args{
		//		start: "AACCGGTA",
		//		end:   "AACCGGTA",
		//		bank:  []string{"AACCGGTA"},
		//	},
		//	want: 0,
		//},
		{
			name: "Test case with no mutation bank",
			args: args{
				start: "AACCGGTT",
				end:   "AACCGGTA",
				bank:  []string{},
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutationV2(tt.args.start, tt.args.end, tt.args.bank); got != tt.want {
				t.Errorf("minMutationV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_longestRepetitionV3(t *testing.T) {
	type args struct {
		dna string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case with no repetition",
			args: args{dna: "ATCG"},
			want: 1,
		},
		{
			name: "Test case with repetition at the beginning",
			args: args{dna: "AAATCG"},
			want: 3,
		},
		{
			name: "Test case with repetition in the middle",
			args: args{dna: "ATTCCGGGA"},
			want: 3,
		},
		{
			name: "Test case with repetition at the end",
			args: args{dna: "ATTCGGGAA"},
			want: 3,
		},
		{
			name: "Test case with all characters the same",
			args: args{dna: "AAAAAAAAA"},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestRepetitionV3(tt.args.dna); got != tt.want {
				t.Errorf("longestRepetitionV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

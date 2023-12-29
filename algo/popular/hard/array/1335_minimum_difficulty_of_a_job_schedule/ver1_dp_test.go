package _335_minimum_difficulty_of_a_job_schedule

import "testing"

func Test_minDifficultyV1(t *testing.T) {
	type args struct {
		jobDifficulty []int
		d             int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Тест 1",
			args: args{
				jobDifficulty: []int{6, 5, 4, 3, 2, 1},
				d:             2,
			},
			want: 7,
		},
		{
			name: "Тест 2",
			args: args{
				jobDifficulty: []int{9, 9, 9},
				d:             4,
			},
			want: -1,
		},
		{
			name: "Тест 3",
			args: args{
				jobDifficulty: []int{1, 1, 1},
				d:             3,
			},
			want: 3,
		},
		{
			name: "Тест 4",
			args: args{
				jobDifficulty: []int{7, 1, 7, 1, 7, 1},
				d:             3,
			},
			want: 15,
		},
		{
			name: "Тест 5",
			args: args{
				jobDifficulty: []int{11, 111, 22, 222, 33, 333, 44, 444},
				d:             6,
			},
			want: 843,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifficultyV1(tt.args.jobDifficulty, tt.args.d); got != tt.want {
				t.Errorf("minDifficultyV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

package _700_number_of_students_unable_to_eat_lunch

import "testing"

func Test_countStudentsV3(t *testing.T) {
	type args struct {
		students   []int
		sandwiches []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				students:   []int{1, 1, 0, 0},
				sandwiches: []int{0, 1, 0, 1},
			},
			want: 0,
		},
		{
			name: "Test case 2",
			args: args{
				students:   []int{1, 1, 1, 0, 0, 1},
				sandwiches: []int{1, 0, 0, 0, 1, 1},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countStudentsV3(tt.args.students, tt.args.sandwiches); got != tt.want {
				t.Errorf("countStudentsV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

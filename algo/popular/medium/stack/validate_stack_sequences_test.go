package stack

import "testing"

func Test_validateStackSequences(t *testing.T) {
	/*
		Дано:
		- На вход подаются 2 массива целых чисел.

		Надо:
		- Проверить, что оба массива являются результатом работы стека. Т.е. один массив это положили значения, а другой - достали из стека.
	*/
	type args struct {
		pushed []int
		popped []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				pushed: []int{1, 2, 3, 4, 5},
				popped: []int{4, 5, 3, 2, 1},
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				pushed: []int{1, 2, 3, 4, 5},
				popped: []int{4, 3, 5, 1, 2},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Simple addition",
			args: args{tokens: []string{"2", "1", "+"}},
			want: 3,
		},
		{
			name: "Simple subtraction",
			args: args{tokens: []string{"5", "3", "-"}},
			want: 2,
		},
		{
			name: "Simple multiplication",
			args: args{tokens: []string{"4", "3", "*"}},
			want: 12,
		},
		{
			name: "Simple division",
			args: args{tokens: []string{"10", "2", "/"}},
			want: 5,
		},
		{
			name: "Complex expression",
			args: args{tokens: []string{"2", "1", "+", "3", "*"}},
			want: 9,
		},
		{
			name: "Another complex expression",
			args: args{tokens: []string{"4", "13", "5", "/", "+"}},
			want: 6,
		},
		{
			name: "Expression with negative result",
			args: args{tokens: []string{"3", "4", "-"}},
			want: -1,
		},
		{
			name: "Expression with division resulting in zero",
			args: args{tokens: []string{"1", "2", "/"}},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

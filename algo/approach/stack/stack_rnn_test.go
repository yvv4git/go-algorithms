package main

import "testing"

func Test_evaluateRPN(t *testing.T) {
	type args struct {
		tokens []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1: Addition and Multiplication",
			args: args{tokens: []string{"2", "1", "+", "3", "*"}},
			want: 9,
		},
		{
			name: "Test case 2: Single Number",
			args: args{tokens: []string{"5"}},
			want: 5,
		},
		{
			name: "Test case 3: Large Numbers",
			args: args{tokens: []string{"1000", "2000", "3000", "+", "*"}},
			want: 5000000,
		},
		{
			name: "Test case 4: Negative Numbers",
			args: args{tokens: []string{"-2", "-3", "*", "5", "+"}},
			want: 11,
		},
		{
			name: "Test case 5: Complex Expression",
			args: args{tokens: []string{"15", "7", "1", "1", "+", "-", "/", "3", "*", "2", "1", "1", "+", "+", "-"}},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evaluateRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

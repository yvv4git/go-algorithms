package main

import "testing"

func Test_parseBoolExprV2(t *testing.T) {
	type args struct {
		expression string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1: Simple NOT",
			args: args{expression: "!(f)"},
			want: true,
		},
		{
			name: "Test Case 2: Simple AND",
			args: args{expression: "&(t,f)"},
			want: false,
		},
		{
			name: "Test Case 3: Simple OR",
			args: args{expression: "|(f,t)"},
			want: true,
		},
		{
			name: "Test Case 5: Multiple Nested Expressions",
			args: args{expression: "!(&(|(f,t),f))"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBoolExprV2(tt.args.expression); got != tt.want {
				t.Errorf("parseBoolExprV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

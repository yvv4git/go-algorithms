package main

import (
	"reflect"
	"testing"
)

func Test_generateParenthesisV2(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test Case 1: n = 1",
			args: args{n: 1},
			want: []string{"()"},
		},
		{
			name: "Test Case 2: n = 2",
			args: args{n: 2},
			want: []string{"()()", "(())"},
		},
		{
			name: "Test Case 3: n = 3",
			args: args{n: 3},
			want: []string{"()()()", "()(())", "(())()", "(()())", "((()))"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesisV2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Logf("%#v", got)
				t.Errorf("generateParenthesisV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

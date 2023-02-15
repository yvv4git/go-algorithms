package remove_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeOuterParentheses(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				s: "(()())(())",
			},
			want: "()()()",
		},
		{
			name: "CASE-2",
			args: args{
				s: "(()())(())(()(()))",
			},
			want: "()()()()(())",
		},
		{
			name: "CASE-2",
			args: args{
				s: "()()",
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeOuterParentheses(tt.args.s))
		})
	}
}

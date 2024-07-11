package main

import (
	"reflect"
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test Case 1",
			args: args{s: ")("},
			want: []string{""},
		},
		{
			name: "Test Case 2",
			args: args{s: "((()"},
			want: []string{"()"},
		},
		{
			name: "Test Case 3",
			args: args{s: "()))"},
			want: []string{"()"},
		},
		{
			name: "Test Case 4",
			args: args{s: "((())))"},
			want: []string{"((()))"},
		},
		{
			name: "Test Case 5",
			args: args{s: ")(f"},
			want: []string{"f"},
		},
		{
			name: "Test Case 6",
			args: args{s: ""},
			want: []string{},
		},
		{
			name: "Test Case 7",
			args: args{s: "()"},
			want: []string{"()"},
		},
		{
			name: "Test Case 8",
			args: args{s: "(()))("},
			want: []string{"(())"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeInvalidParentheses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeInvalidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}

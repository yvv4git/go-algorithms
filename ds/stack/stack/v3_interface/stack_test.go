package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type args struct {
		val interface{}
	}

	testCases := []struct {
		name string
		args args
		want []interface{}
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				val: nil,
			},
			want: []interface{}{nil},
			desc: "В стек можем поместить даже nil",
		},
		{
			name: "CASE-2",
			args: args{
				val: 1,
			},
			want: []interface{}{1},
			desc: "В стек можем поместить int",
		},
		{
			name: "CASE-3",
			args: args{
				val: 1,
			},
			want: []interface{}{"Vladimir"},
			desc: "В стек можем поместить string",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var s Stack
			s.Push(tt.args.val)

			//t.Log(s.values)
			assert.Equal(t, tt.want, s.values)
		})
	}
}

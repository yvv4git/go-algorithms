package rune_brackets

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "CASE-1",
			str:  "[](){}",
			want: true,
		},
		{
			name: "CASE-2",
			str:  "[({})]",
			want: true,
		},
		{
			name: "CASE-3",
			str:  "[(])",
			want: false,
		},
		{
			name: "CASE-4",
			str:  "}{",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Validate(tt.str)
			assert.Equal(t, tt.want, result)
		})
	}
}

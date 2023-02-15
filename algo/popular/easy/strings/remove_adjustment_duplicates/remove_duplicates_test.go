package remove_adjustment_duplicates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
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
				s: "abbaca",
			},
			want: "ca",
		},
		{
			name: "CASE-2",
			args: args{
				s: "azxxzy",
			},
			want: "ay",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeDuplicates(tt.args.s))
		})
	}
}

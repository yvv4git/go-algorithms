package _7_letter_combinations_of_a_phone_number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_letterCombinationsV2(t *testing.T) {
	type args struct {
		digits string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				digits: "23",
			},
			want: []string{"ad", "bd", "cd", "ae", "be", "ce", "af", "bf", "cf"},
		},
		{
			name: "Test case 2",
			args: args{
				digits: "2",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test case 3",
			args: args{
				digits: "",
			},
			want: []string{},
		},
		{
			name: "Test case 4",
			args: args{
				digits: "234",
			},
			want: []string{"adg", "bdg", "cdg", "aeg", "beg", "ceg", "afg", "bfg", "cfg", "adh", "bdh", "cdh", "aeh", "beh", "ceh", "afh", "bfh", "cfh", "adi", "bdi", "cdi", "aei", "bei", "cei", "afi", "bfi", "cfi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, letterCombinationsV2(tt.args.digits))
		})
	}
}

package _7_letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

func Test_letterCombinationsV1(t *testing.T) {
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
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
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
			want: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinationsV1(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinationsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

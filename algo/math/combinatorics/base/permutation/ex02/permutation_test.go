package ex02

import "testing"

func TestPermutation(t *testing.T) {
	type args struct {
		data   []int
		i      int
		length int
	}

	testCases := []struct {
		name string
		args args
	}{
		{
			name: "CASE-1",
			args: args{
				data:   []int{1, 2, 3, 4, 5},
				i:      3,
				length: 5,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			Permutation(tc.args.data, tc.args.i, tc.args.length)
		})
	}
}

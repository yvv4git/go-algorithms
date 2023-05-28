package decode_xored_array_v1

import (
	"reflect"
	"testing"
)

func Test_decodeV2(t *testing.T) {
	type args struct {
		encoded []int
		first   int
	}

	tests := []struct {
		name string
		args args
		want []int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				encoded: []int{1, 2, 3},
				first:   1,
			},
			want: []int{1, 0, 2, 1},
			desc: `Explanation: If arr = [1,0,2,1], then first = 1 and encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]`,
		},
		{
			name: "CASE-2",
			args: args{
				encoded: []int{6, 2, 7, 3},
				first:   4,
			},
			want: []int{4, 2, 0, 7, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeV2(tt.args.encoded, tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

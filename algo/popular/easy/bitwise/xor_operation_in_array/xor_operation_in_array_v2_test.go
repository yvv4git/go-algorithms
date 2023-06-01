package xor_operation_in_array

import "testing"

func Test_xorOperationV2(t *testing.T) {
	type args struct {
		n     int
		start int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				n:     5,
				start: 0,
			},
			want: 8,
		},
		{
			name: "CASE-2",
			args: args{
				n:     4,
				start: 3,
			},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperationV2(tt.args.n, tt.args.start); got != tt.want {
				t.Errorf("xorOperationV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package __zigzag_conversion

import "testing"

func Test_convertV1(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "CASE-2",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "CASE-3",
			args: args{
				s:       "A",
				numRows: 1,
			},
			want: "A",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertV1(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convertV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

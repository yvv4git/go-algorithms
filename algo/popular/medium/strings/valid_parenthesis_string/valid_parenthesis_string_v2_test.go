package valid_parenthesis_string

import "testing"

func Test_checkValidStringV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				s: "()",
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				s: "(*)",
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				s: "(*))",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidStringV2(tt.args.s); got != tt.want {
				t.Errorf("checkValidStringV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package is_palindrome

import "testing"

func Test_isPalindromeRu(t *testing.T) {
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
				s: "lool",
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				s: "довод",
			},
			want: true,
		},
		{
			name: "CASE-3",
			args: args{
				s: "доход",
			},
			want: true,
		},
		{
			name: "CASE-4",
			args: args{
				s: "шабаш",
			},
			want: true,
		},
		{
			name: "CASE-5",
			args: args{
				s: "Шабаш",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeRu(tt.args.s); got != tt.want {
				t.Errorf("isPalindromeRu() = %v, want %v", got, tt.want)
			}
		})
	}
}

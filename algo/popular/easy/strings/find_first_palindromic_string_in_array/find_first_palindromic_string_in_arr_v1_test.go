package find_first_palindromic_string_in_array

import "testing"

func Test_firstPalindromeV1(t *testing.T) {
	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				words: []string{"abc", "car", "ada", "racecar", "cool"},
			},
			want: "ada",
		},
		{
			name: "CASE-2",
			args: args{
				words: []string{"notapalindrome", "racecar"},
			},
			want: "racecar",
		},
		{
			name: "CASE-3",
			args: args{
				words: []string{"def", "ghi"},
			},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstPalindromeV1(tt.args.words); got != tt.want {
				t.Errorf("firstPalindromeV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

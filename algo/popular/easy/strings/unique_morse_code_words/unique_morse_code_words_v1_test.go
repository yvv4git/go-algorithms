package unique_morse_code_words

import "testing"

func Test_uniqueMorseRepresentationsV1(t *testing.T) {
	type args struct {
		words []string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				words: []string{"gin", "zen", "gig", "msg"},
			},
			want: 2,
		},
		{
			name: "CASE-2",
			args: args{
				words: []string{"a"},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniqueMorseRepresentationsV1(tt.args.words); got != tt.want {
				t.Errorf("uniqueMorseRepresentationsV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_shortestCompletingWordV2(t *testing.T) {
	type args struct {
		licensePlate string
		words        []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				licensePlate: "1s3 PSt",
				words:        []string{"step", "steps", "stripe", "stepple"},
			},
			want: "steps",
		},
		{
			name: "Example 2",
			args: args{
				licensePlate: "1s3 456",
				words:        []string{"looks", "pest", "stew", "show"},
			},
			want: "pest",
		},
		{
			name: "Example 3",
			args: args{
				licensePlate: "Ah71752",
				words:        []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"},
			},
			want: "husband",
		},
		{
			name: "Example 4",
			args: args{
				licensePlate: "iMSlpe4",
				words:        []string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"},
			},
			want: "simple",
		},
		{
			name: "Example 5",
			args: args{
				licensePlate: "Ar16259",
				words:        []string{"nature", "though", "party", "food", "any", "democratic", "building", "eat", "structure", "our"},
			},
			want: "nature",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCompletingWordV2(tt.args.licensePlate, tt.args.words); got != tt.want {
				t.Errorf("shortestCompletingWordV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

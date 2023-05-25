package maximum_number_balloons

import "testing"

func Test_maxNumberOfBalloonsV2(t *testing.T) {
	type args struct {
		text string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				text: "nlaebolko",
			},
			want: 1,
		},
		{
			name: "CASE-2",
			args: args{
				text: "loonbalxballpoon",
			},
			want: 2,
		},
		{
			name: "CASE-3",
			args: args{
				text: "leetcode",
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfBalloonsV2(tt.args.text); got != tt.want {
				t.Errorf("maxNumberOfBalloonsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

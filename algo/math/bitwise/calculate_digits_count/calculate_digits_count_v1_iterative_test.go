package calculate_digits_count

import "testing"

func TestCalculateDigitsNumberV1(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{num: 1},
			want: 1,
		},
		{
			name: "CASE-2",
			args: args{num: 5},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateDigitsNumberV1(tt.args.num); got != tt.want {
				t.Errorf("CalculateDigitsNumberV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

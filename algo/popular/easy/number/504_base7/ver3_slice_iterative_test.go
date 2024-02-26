package main

import "testing"

func Test_convertToBase7V4(t *testing.T) {
	type args struct {
		num int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Zero",
			args: args{num: 0},
			want: "0",
		},
		{
			name: "Positive number",
			args: args{num: 100},
			want: "202",
		},
		{
			name: "Negative number",
			args: args{num: -7},
			want: "-10",
		},
		{
			name: "Small negative number",
			args: args{num: -3},
			want: "-3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToBase7V4(tt.args.num); got != tt.want {
				t.Errorf("convertToBase7V4() = %v, want %v", got, tt.want)
			}
		})
	}
}

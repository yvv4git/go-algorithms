package main

import "testing"

func Test_getHintV2(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1: No matches",
			args: args{
				secret: "1234",
				guess:  "5678",
			},
			want: "0A0B",
		},
		{
			name: "Test Case 2: All digits match but in wrong position",
			args: args{
				secret: "1234",
				guess:  "4321",
			},
			want: "0A4B",
		},
		{
			name: "Test Case 3: All digits match in correct position",
			args: args{
				secret: "1234",
				guess:  "1234",
			},
			want: "4A0B",
		},
		{
			name: "Test Case 4: Repeated digits in guess",
			args: args{
				secret: "1123",
				guess:  "0111",
			},
			want: "1A1B",
		},
		{
			name: "Test Case 5",
			args: args{
				secret: "1807",
				guess:  "7810",
			},
			want: "1A3B",
		},
		{
			name: "Test Case 6",
			args: args{
				secret: "1123",
				guess:  "0111",
			},
			want: "1A1B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHintV2(tt.args.secret, tt.args.guess); got != tt.want {
				t.Errorf("getHintV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

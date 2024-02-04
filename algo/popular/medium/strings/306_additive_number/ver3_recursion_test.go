package main

import "testing"

func Test_isAdditiveNumberV3(t *testing.T) {
	type args struct {
		num string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1: Positive test case with additive number",
			args: args{num: "112358"},
			want: true,
		},
		{
			name: "Test Case 2: Positive test case with additive number",
			args: args{num: "199100199"},
			want: true,
		},
		{
			name: "Test Case 3: Positive test case with additive number",
			args: args{num: "123"},
			want: true,
		},
		{
			name: "Test Case 4: Negative test case with non-additive number",
			args: args{num: "1023"},
			want: false,
		},
		{
			name: "Test Case 5: Positive test case with additive number starting with two zeros",
			args: args{num: "000"},
			want: true,
		},
		{
			name: "Test Case 7: Positive test case with additive number with multiple zeros",
			args: args{num: "101020305080130210"},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAdditiveNumberV3(tt.args.num); got != tt.want {
				t.Errorf("isAdditiveNumberV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

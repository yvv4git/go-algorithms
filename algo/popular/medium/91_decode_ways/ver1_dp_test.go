package main

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{s: "12"},
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{s: "226"},
			want: 3,
		},
		{
			name: "Test Case 3",
			args: args{s: "06"},
			want: 0,
		},
		{
			name: "Test Case 4",
			args: args{s: "11106"},
			want: 2,
		},
		{
			name: "Test Case 5",
			args: args{s: "12345"},
			want: 3,
		},
		{
			name: "Test Case 6",
			args: args{s: "27"},
			want: 1,
		},
		{
			name: "Test Case 7",
			args: args{s: "10"},
			want: 1,
		},
		{
			name: "Test Case 8",
			args: args{s: "100"},
			want: 0,
		},
		{
			name: "Test Case 9",
			args: args{s: "101"},
			want: 1,
		},
		{
			name: "Test Case 10",
			args: args{s: "110"},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}

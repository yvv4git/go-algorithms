package main

import "testing"

func Test_compressV4(t *testing.T) {
	type args struct {
		chars []byte
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case with empty array",
			args: args{chars: []byte{}},
			want: 0,
		},
		{
			name: "Test case with no repeating characters",
			args: args{chars: []byte("abc")},
			want: 3,
		},
		{
			name: "Test case with repeating characters",
			args: args{chars: []byte("aabbcc")},
			want: 6,
		},
		{
			name: "Test case with repeating characters and single characters",
			args: args{chars: []byte("aabbc")},
			want: 5,
		},
		{
			name: "Test case with repeating characters and numbers",
			args: args{chars: []byte("aabbcc111")},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressV4(tt.args.chars); got != tt.want {
				t.Errorf("compressV4() = %v, want %v", got, tt.want)
			}
		})
	}
}

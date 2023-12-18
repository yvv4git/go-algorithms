package _83_ransom_note

import "testing"

func Test_canConstructV1(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{
				ransomNote: "aa",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
		{
			name: "Test case 4",
			args: args{
				ransomNote: "abc",
				magazine:   "aabbcc",
			},
			want: true,
		},
		{
			name: "Test case 5",
			args: args{
				ransomNote: "abc",
				magazine:   "abc",
			},
			want: true,
		},
		{
			name: "Test case 6",
			args: args{
				ransomNote: "abc",
				magazine:   "abcd",
			},
			want: true,
		},
		{
			name: "Test case 7",
			args: args{
				ransomNote: "abc",
				magazine:   "ab",
			},
			want: false,
		},
		{
			name: "Test case 8",
			args: args{
				ransomNote: "abc",
				magazine:   "",
			},
			want: false,
		},
		{
			name: "Test case 9",
			args: args{
				ransomNote: "",
				magazine:   "abc",
			},
			want: true,
		},
		{
			name: "Test case 10",
			args: args{
				ransomNote: "",
				magazine:   "",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstructV1(tt.args.ransomNote, tt.args.magazine); got != tt.want {
				t.Errorf("canConstructV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

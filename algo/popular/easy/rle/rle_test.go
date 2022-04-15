package main

import "testing"

func Test_rleString(t *testing.T) {
	type args struct {
		in string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "CASE-1",
			args: args{
				in: "",
			},
			want: "",
		},
		{
			name: "CASE-2",
			args: args{
				in: "aaabbavv",
			},
			want: "a3b2a1v2",
		},
		{
			name: "CASE-3",
			args: args{
				in: "Владимир",
			},
			want: "В1л1а1д1и1м1и1р1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rleString(tt.args.in)
			//t.Log(tt.args.in, "--->", result)
			if result != tt.want {
				t.Errorf("rleString() = %v, want %v", result, tt.want)
			}
		})
	}
}

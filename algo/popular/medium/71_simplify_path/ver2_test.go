package main

import "testing"

func Test_simplifyPathV2(t *testing.T) {
	type args struct {
		path string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1: Simple path",
			args: args{
				path: "/home/",
			},
			want: "/home",
		},
		{
			name: "Test Case 2: Go back to root",
			args: args{
				path: "/../",
			},
			want: "/",
		},
		{
			name: "Test Case 3: Multiple slashes",
			args: args{
				path: "/home//foo/",
			},
			want: "/home/foo",
		},
		{
			name: "Test Case 4: Go back multiple directories",
			args: args{
				path: "/a/./b/../../c/",
			},
			want: "/c",
		},
		{
			name: "Test Case 5: Empty path",
			args: args{
				path: "",
			},
			want: "/",
		},
		{
			name: "Test Case 6: Single directory",
			args: args{
				path: "dir",
			},
			want: "/dir",
		},
		{
			name: "Test Case 7: Multiple directories",
			args: args{
				path: "/dir1/dir2/dir3/",
			},
			want: "/dir1/dir2/dir3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := simplifyPathV2(tt.args.path); got != tt.want {
				t.Errorf("simplifyPathV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_toLowerCase(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1: Mixed Case String",
			args: args{str: "Hello, World!"},
			want: "hello, world!",
		},
		{
			name: "Test Case 2: All Uppercase String",
			args: args{str: "HELLO, WORLD!"},
			want: "hello, world!",
		},
		{
			name: "Test Case 3: All Lowercase String",
			args: args{str: "hello, world!"},
			want: "hello, world!",
		},
		{
			name: "Test Case 4: String with Non-Alphabetic Characters",
			args: args{str: "Hello, World!123"},
			want: "hello, world!123",
		},
		{
			name: "Test Case 5: Empty String",
			args: args{str: ""},
			want: "",
		},
		{
			name: "Test Case 6: String with Only Non-Alphabetic Characters",
			args: args{str: "1234567890!@#$%^&*()-_=+[]{}|;:',.<>?/"},
			want: "1234567890!@#$%^&*()-_=+[]{}|;:',.<>?/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.str); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

func Test_addStringsV2(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1: Add two single-digit numbers",
			args: args{
				num1: "1",
				num2: "2",
			},
			want: "3",
		},
		{
			name: "Test Case 2: Add two multi-digit numbers",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "579",
		},
		{
			name: "Test Case 3: Add two numbers with different lengths",
			args: args{
				num1: "123456789",
				num2: "987654321",
			},
			want: "1111111110",
		},
		{
			name: "Test Case 4: Add two numbers with carry",
			args: args{
				num1: "999",
				num2: "1",
			},
			want: "1000",
		},
		{
			name: "Test Case 5: Add two zeros",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStringsV2(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStringsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddressesV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test Case 1: Valid IP Address",
			args: args{
				s: "25525511135",
			},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "Test Case 2: Leading Zero",
			args: args{
				s: "010010",
			},
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			name: "Test Case 3: All Zeros",
			args: args{
				s: "0000",
			},
			want: []string{"0.0.0.0"},
		},
		{
			name: "Test Case 4: Large Valid IP",
			args: args{
				s: "255255255255",
			},
			want: []string{"255.255.255.255"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddressesV2(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddressesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

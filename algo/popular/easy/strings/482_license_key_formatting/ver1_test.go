package main

import "testing"

func Test_licenseKeyFormatting(t *testing.T) {
	type args struct {
		s string
		k int
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}

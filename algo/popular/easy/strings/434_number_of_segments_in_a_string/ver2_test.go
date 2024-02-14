package main

import "testing"

func Test_countSegmentsV2(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "No spaces",
			args: args{s: "HelloWorld"},
			want: 1,
		},
		{
			name: "One space at the beginning",
			args: args{s: " HelloWorld"},
			want: 1,
		},
		{
			name: "One space at the end",
			args: args{s: "HelloWorld "},
			want: 1,
		},
		{
			name: "Multiple spaces at the beginning",
			args: args{s: "   HelloWorld"},
			want: 1,
		},
		{
			name: "Multiple spaces at the end",
			args: args{s: "HelloWorld   "},
			want: 1,
		},
		{
			name: "Multiple words with one space",
			args: args{s: "Hello World"},
			want: 2,
		},
		{
			name: "Multiple words with multiple spaces",
			args: args{s: "  Hello   World  "},
			want: 2,
		},
		{
			name: "Empty string",
			args: args{s: ""},
			want: 0,
		},
		{
			name: "String with only spaces",
			args: args{s: "   "},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotInc := countSegmentsV2(tt.args.s); gotInc != tt.want {
				t.Errorf("countSegmentsV2() = %v, want %v", gotInc, tt.want)
			}
		})
	}
}

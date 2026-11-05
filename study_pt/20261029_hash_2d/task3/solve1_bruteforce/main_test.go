package main

import (
	"reflect"
	"testing"
)

func Test_countPatternOccurrences(t *testing.T) {
	examplePatterns := [][]string{
		{"010", "111", "010"},
		{"100", "011", "001"},
	}
	exampleImage := []string{
		"0010010010",
		"0100001111",
		"0000000100",
		"0000100000",
		"0001111000",
		"0000101000",
		"0100000010",
		"0001000110",
		"0011111000",
		"0101000000",
	}

	tests := []struct {
		name     string
		n        int
		m        int
		l        int
		patterns [][]string
		image    []string
		want     []int
	}{
		{
			name:     "example 1",
			n:        10,
			m:        10,
			l:        3,
			patterns: examplePatterns,
			image:    exampleImage,
			want:     []int{2, 1},
		},
		{
			name:     "single cell patterns",
			n:        2,
			m:        2,
			l:        1,
			patterns: [][]string{{"0"}, {"1"}},
			image:    []string{"01", "10"},
			want:     []int{2, 2},
		},
		{
			name:     "pattern equals whole image",
			n:        2,
			m:        2,
			l:        2,
			patterns: [][]string{{"10", "01"}},
			image:    []string{"10", "01"},
			want:     []int{1},
		},
		{
			name:     "pattern larger than image",
			n:        2,
			m:        2,
			l:        3,
			patterns: [][]string{{"111", "111", "111"}},
			image:    []string{"11", "11"},
			want:     []int{0},
		},
		{
			name:     "overlapping occurrences",
			n:        3,
			m:        3,
			l:        2,
			patterns: [][]string{{"11", "11"}},
			image:    []string{"111", "111", "111"},
			want:     []int{4},
		},
		{
			name:     "two patterns one absent",
			n:        2,
			m:        3,
			l:        2,
			patterns: [][]string{{"00", "00"}, {"11", "11"}},
			image:    []string{"000", "000"},
			want:     []int{2, 0},
		},
		{
			name:     "rotated pattern is not counted",
			n:        2,
			m:        2,
			l:        2,
			patterns: [][]string{{"01", "10"}},
			image:    []string{"10", "01"},
			want:     []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPatternOccurrences(tt.n, tt.m, tt.l, tt.patterns, tt.image)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("countPatternOccurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}

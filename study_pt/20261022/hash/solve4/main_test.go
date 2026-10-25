package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func naiveEqual(s string, a, b, length int) bool {
	for i := 0; i < length; i++ {
		if s[a+i] != s[b+i] {
			return false
		}
	}
	return true
}

func naiveSolve(s string, queries [][3]int) []string {
	answers := make([]string, len(queries))
	for i, query := range queries {
		if naiveEqual(s, query[1], query[2], query[0]) {
			answers[i] = "yes"
		} else {
			answers[i] = "no"
		}
	}
	return answers
}

func TestSolve(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		queries [][3]int
		want    []string
	}{
		{
			name:    "example 1",
			s:       "acabaca",
			queries: [][3]int{{4, 3, 2}, {3, 4, 0}, {2, 0, 1}},
			want:    []string{"no", "yes", "no"},
		},
		{
			name: "example 2",
			s:    "caeabaeadedcbdcdccec",
			queries: [][3]int{
				{13, 4, 3},
				{2, 12, 15},
				{10, 1, 3},
				{3, 8, 15},
				{13, 5, 6},
				{7, 2, 6},
				{9, 8, 8},
				{19, 0, 0},
				{19, 0, 0},
				{6, 7, 13},
			},
			want: []string{"no", "no", "no", "no", "no", "no", "yes", "yes", "yes", "no"},
		},
		{
			name:    "single character",
			s:       "z",
			queries: [][3]int{{1, 0, 0}},
			want:    []string{"yes"},
		},
		{
			name:    "no queries",
			s:       "abc",
			queries: [][3]int{},
			want:    []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.s, tt.queries)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve(%q, %v) = %v, want %v", tt.s, tt.queries, got, tt.want)
			}
		})
	}
}

func TestSolveRandom(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for iteration := 0; iteration < 300; iteration++ {
		n := 1 + rng.Intn(40)
		data := make([]byte, n)
		for i := range data {
			data[i] = byte('a' + rng.Intn(4))
		}
		s := string(data)

		q := rng.Intn(40)
		queries := make([][3]int, q)
		for i := range queries {
			length := 1 + rng.Intn(n)
			a := rng.Intn(n - length + 1)
			b := rng.Intn(n - length + 1)
			queries[i] = [3]int{length, a, b}
		}

		got := solve(s, queries)
		want := naiveSolve(s, queries)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("solve(%q, %v) = %v, want %v", s, queries, got, want)
		}
	}
}

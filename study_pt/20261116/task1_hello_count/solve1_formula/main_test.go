package main

import "testing"

func Test_countGreetings(t *testing.T) {
	tests := []struct {
		name string
		a    int64
		b    int64
		want int64
	}{
		{name: "example 1", a: 1, b: 2, want: 6},
		{name: "example 2", a: 5, b: 0, want: 0},
		{name: "empty open space", a: 0, b: 0, want: 0},
		{name: "only one extrovert", a: 0, b: 1, want: 0},
		{name: "two extroverts", a: 0, b: 2, want: 2},
		{name: "one introvert one extrovert", a: 1, b: 1, want: 2},
		{name: "only introverts", a: 3, b: 0, want: 0},
		{name: "two introverts one extrovert", a: 2, b: 1, want: 4},
		{name: "three extroverts", a: 0, b: 3, want: 6},
		{name: "two of each", a: 2, b: 2, want: 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGreetings(tt.a, tt.b); got != tt.want {
				t.Fatalf("countGreetings(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func Test_countGreetings_maxSize(t *testing.T) {
	const max = int64(1000000000)
	// 2*10^9*10^9 + 10^9*(10^9-1) = 2*10^18 + 10^18 - 10^9 = 3*10^18 - 10^9.
	want := int64(2999999999000000000)
	if got := countGreetings(max, max); got != want {
		t.Fatalf("countGreetings(%d, %d) = %d, want %d", max, max, got, want)
	}
}

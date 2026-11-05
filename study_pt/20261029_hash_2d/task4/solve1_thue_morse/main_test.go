package main

import "testing"

func hashString(s string, x int) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result = (result*x + int(s[i])) % 128
	}
	return result
}

func verifyHackedStrings(t *testing.T, n int, got []string) {
	t.Helper()

	if len(got) != n {
		t.Fatalf("len(buildHackedStrings(%d)) = %d, want %d", n, len(got), n)
	}

	seen := make(map[string]struct{}, n)
	for _, s := range got {
		if len(s) == 0 || len(s) > 5000 {
			t.Fatalf("string length = %d, want from 1 to 5000", len(s))
		}
		for i := 0; i < len(s); i++ {
			if s[i] < 'a' || s[i] > 'z' {
				t.Fatalf("string contains non-lowercase char %q", s[i])
			}
		}
		if _, ok := seen[s]; ok {
			t.Fatalf("duplicate string found")
		}
		seen[s] = struct{}{}
	}

	for x := 2; x <= 127; x++ {
		want := hashString(got[0], x)
		for i := 1; i < len(got); i++ {
			if gotHash := hashString(got[i], x); gotHash != want {
				t.Fatalf("hash mismatch for x = %d and string %d: got %d, want %d", x, i, gotHash, want)
			}
		}
	}
}

func Test_buildHackedStrings(t *testing.T) {
	tests := []struct {
		name string
		n    int
	}{
		{name: "one string", n: 1},
		{name: "two strings", n: 2},
		{name: "three strings", n: 3},
		{name: "seventeen strings", n: 17},
		{name: "one less than maximum", n: 127},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			verifyHackedStrings(t, tt.n, buildHackedStrings(tt.n))
		})
	}
}

func Test_buildHackedStrings_maxSize(t *testing.T) {
	const n = 128
	verifyHackedStrings(t, n, buildHackedStrings(n))
}

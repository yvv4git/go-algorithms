package main

import "testing"

func totalPresses(freq []int64, groups []int) int64 {
	idx := 0
	var total int64
	for _, g := range groups {
		for i := 0; i < g; i++ {
			total += int64(i+1) * freq[idx]
			idx++
		}
	}
	return total
}

func sum(groups []int) int {
	s := 0
	for _, g := range groups {
		s += g
	}
	return s
}

func Test_distributeEmojis(t *testing.T) {
	tests := []struct {
		name     string
		freq     []int64
		buttons  int
		wantCost int64
	}{
		{
			name:     "example 1",
			freq:     []int64{1, 3, 5, 1, 1},
			buttons:  3,
			wantCost: 14,
		},
		{
			name:     "single button",
			freq:     []int64{2, 1, 3},
			buttons:  1,
			wantCost: 13,
		},
		{
			name:     "more buttons than emojis",
			freq:     []int64{5, 1},
			buttons:  4,
			wantCost: 6,
		},
		{
			name:     "equal buttons and emojis",
			freq:     []int64{4, 2, 3},
			buttons:  3,
			wantCost: 9,
		},
		{
			name:     "empty group helpful",
			freq:     []int64{1, 10},
			buttons:  3,
			wantCost: 11,
		},
		{
			name:     "all zero frequencies",
			freq:     []int64{0, 0, 0, 0},
			buttons:  2,
			wantCost: 0,
		},
		{
			name:     "single emoji",
			freq:     []int64{7},
			buttons:  1,
			wantCost: 7,
		},
		{
			name:     "no emojis",
			freq:     []int64{},
			buttons:  3,
			wantCost: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := distributeEmojis(tt.freq, tt.buttons)
			if len(got) != tt.buttons {
				t.Fatalf("len = %d, want %d", len(got), tt.buttons)
			}
			if gotSum := sum(got); gotSum != len(tt.freq) {
				t.Fatalf("sum = %d, want %d", gotSum, len(tt.freq))
			}
			if gotCost := totalPresses(tt.freq, got); gotCost != tt.wantCost {
				t.Errorf("cost = %d, want %d, groups=%v", gotCost, tt.wantCost, got)
			}
		})
	}
}

func Test_distributeEmojis_maxSize(t *testing.T) {
	freq := make([]int64, 30)
	for i := range freq {
		freq[i] = 1
	}
	got := distributeEmojis(freq, 10)
	if len(got) != 10 {
		t.Fatalf("len = %d, want 10", len(got))
	}
	if gotSum := sum(got); gotSum != 30 {
		t.Fatalf("sum = %d, want 30", gotSum)
	}
	if gotCost := totalPresses(freq, got); gotCost != 60 {
		t.Errorf("cost = %d, want 60, groups=%v", gotCost, got)
	}
}

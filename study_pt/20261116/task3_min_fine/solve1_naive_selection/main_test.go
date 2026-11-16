package main

import (
	"slices"
	"testing"
)

func totalFine(deadlines, days []int) int {
	sum := 0
	for i, day := range days {
		if day > deadlines[i] {
			sum += day - deadlines[i]
		}
	}
	return sum
}

func optimalFine(deadlines []int) int {
	sorted := append([]int(nil), deadlines...)
	slices.Sort(sorted)
	sum := 0
	for i, deadline := range sorted {
		day := i + 1
		if day > deadline {
			sum += day - deadline
		}
	}
	return sum
}

func validSchedule(deadlines, days []int) bool {
	n := len(deadlines)
	if len(days) != n {
		return false
	}
	used := make([]bool, n+1)
	for _, day := range days {
		if day < 1 || day > n || used[day] {
			return false
		}
		used[day] = true
	}
	return totalFine(deadlines, days) == optimalFine(deadlines)
}

func Test_solveDeadlines(t *testing.T) {
	tests := []struct {
		name      string
		deadlines []int
	}{
		{name: "example 1", deadlines: []int{2, 3, 1}},
		{name: "example 2", deadlines: []int{2, 1, 2, 1}},
		{name: "single problem", deadlines: []int{5}},
		{name: "all equal deadlines", deadlines: []int{1, 1, 1, 1}},
		{name: "already sorted", deadlines: []int{1, 2, 3, 4}},
		{name: "reverse sorted", deadlines: []int{4, 3, 2, 1}},
		{name: "duplicate deadlines", deadlines: []int{3, 1, 3, 1, 2}},
		{name: "deadlines larger than n", deadlines: []int{1000000, 1000000, 1000000}},
		{name: "strict deadlines", deadlines: []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveDeadlines(tt.deadlines)
			if !validSchedule(tt.deadlines, got) {
				t.Fatalf("solveDeadlines(%v) = %v, fine %d, want fine %d",
					tt.deadlines, got, totalFine(tt.deadlines, got), optimalFine(tt.deadlines))
			}
		})
	}
}

func Test_solveDeadlines_maxSize(t *testing.T) {
	deadlines := make([]int, 100)
	for i := range deadlines {
		deadlines[i] = 1
	}

	got := solveDeadlines(deadlines)
	if !validSchedule(deadlines, got) {
		t.Fatalf("solveDeadlines(maxSize) = %v, fine %d, want fine %d",
			got, totalFine(deadlines, got), optimalFine(deadlines))
	}
	if want := 4950; totalFine(deadlines, got) != want {
		t.Fatalf("solveDeadlines(maxSize) fine = %d, want %d", totalFine(deadlines, got), want)
	}
}

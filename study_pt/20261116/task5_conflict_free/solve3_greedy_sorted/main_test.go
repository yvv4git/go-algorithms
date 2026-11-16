package main

import "testing"

func hasConflictSorted(transactions []Transaction, order []int) bool {
	byID := make(map[int]Transaction, len(transactions))
	for _, tx := range transactions {
		byID[tx.ID] = tx
	}
	written := make(map[int]bool)
	for _, id := range order {
		tx, ok := byID[id]
		if !ok {
			return true
		}
		for _, loc := range tx.RS {
			if written[loc] {
				return true
			}
		}
		for _, loc := range tx.WS {
			written[loc] = true
		}
	}
	return false
}

func Test_selectGreedySorted(t *testing.T) {
	tests := []struct {
		name         string
		m            int
		transactions []Transaction
		minLen       int
	}{
		{
			name: "example 1",
			m:    3,
			transactions: []Transaction{
				{ID: 1, RS: []int{1}, WS: []int{2}},
				{ID: 2, RS: []int{2}, WS: []int{3}},
				{ID: 3, RS: []int{3}, WS: []int{1}},
			},
			minLen: 1,
		},
		{
			name: "single transaction",
			m:    10,
			transactions: []Transaction{
				{ID: 1, RS: []int{5}, WS: []int{7}},
			},
			minLen: 1,
		},
		{
			name:         "no transactions",
			m:            5,
			transactions: []Transaction{},
			minLen:       0,
		},
		{
			name: "all independent",
			m:    4,
			transactions: []Transaction{
				{ID: 1, RS: []int{1}, WS: []int{2}},
				{ID: 2, RS: []int{3}, WS: []int{4}},
			},
			minLen: 2,
		},
		{
			name: "heuristic can reorder",
			m:    3,
			transactions: []Transaction{
				{ID: 1, RS: []int{1, 2}, WS: []int{3}},
				{ID: 2, RS: []int{3}, WS: []int{1}},
			},
			minLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectGreedySorted(tt.transactions, tt.m)
			if len(got) < tt.minLen {
				t.Fatalf("len = %d, want >= %d, order=%v", len(got), tt.minLen, got)
			}
			if hasConflictSorted(tt.transactions, got) {
				t.Fatalf("order has conflict: %v", got)
			}
		})
	}
}

func Test_selectGreedySorted_maxSize(t *testing.T) {
	n := 100000
	transactions := make([]Transaction, n)
	for i := 0; i < n; i++ {
		transactions[i] = Transaction{ID: i + 1, RS: []int{i + 1}, WS: []int{i + 1}}
	}
	got := selectGreedySorted(transactions, n)
	if len(got) != n {
		t.Fatalf("len = %d, want %d", len(got), n)
	}
	if hasConflictSorted(transactions, got) {
		t.Fatalf("order has conflict")
	}
}

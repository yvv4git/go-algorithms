package main

import "testing"

func hasConflictGreedy(transactions []Transaction, order []int) bool {
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

func Test_selectGreedy(t *testing.T) {
	tests := []struct {
		name         string
		m            int
		transactions []Transaction
		wantLen      int
	}{
		{
			name: "example 1",
			m:    3,
			transactions: []Transaction{
				{ID: 1, RS: []int{1}, WS: []int{2}},
				{ID: 2, RS: []int{2}, WS: []int{3}},
				{ID: 3, RS: []int{3}, WS: []int{1}},
			},
			wantLen: 2,
		},
		{
			name: "single transaction",
			m:    10,
			transactions: []Transaction{
				{ID: 1, RS: []int{5}, WS: []int{7}},
			},
			wantLen: 1,
		},
		{
			name:         "no transactions",
			m:            5,
			transactions: []Transaction{},
			wantLen:      0,
		},
		{
			name: "all independent",
			m:    4,
			transactions: []Transaction{
				{ID: 1, RS: []int{1}, WS: []int{2}},
				{ID: 2, RS: []int{3}, WS: []int{4}},
			},
			wantLen: 2,
		},
		{
			name: "write after read is allowed",
			m:    1,
			transactions: []Transaction{
				{ID: 1, RS: []int{1}, WS: []int{1}},
				{ID: 2, RS: []int{1}, WS: []int{1}},
			},
			wantLen: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectGreedy(tt.transactions, tt.m)
			if len(got) != tt.wantLen {
				t.Fatalf("len = %d, want %d, order=%v", len(got), tt.wantLen, got)
			}
			if hasConflictGreedy(tt.transactions, got) {
				t.Fatalf("order has conflict: %v", got)
			}
		})
	}
}

func Test_selectGreedy_maxSize(t *testing.T) {
	n := 100000
	transactions := make([]Transaction, n)
	for i := 0; i < n; i++ {
		transactions[i] = Transaction{ID: i + 1, RS: []int{i + 1}, WS: []int{i + 1}}
	}
	got := selectGreedy(transactions, n)
	if len(got) != n {
		t.Fatalf("len = %d, want %d", len(got), n)
	}
	if hasConflictGreedy(transactions, got) {
		t.Fatalf("order has conflict")
	}
}

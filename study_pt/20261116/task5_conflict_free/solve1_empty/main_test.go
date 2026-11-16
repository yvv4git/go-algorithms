package main

import "testing"

func hasConflict(transactions []Transaction, order []int) bool {
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

func Test_selectEmpty(t *testing.T) {
	tests := []struct {
		name         string
		transactions []Transaction
	}{
		{
			name: "example 1",
			transactions: []Transaction{
				{ID: 1, RS: []int{1}, WS: []int{2}},
				{ID: 2, RS: []int{2}, WS: []int{3}},
				{ID: 3, RS: []int{3}, WS: []int{1}},
			},
		},
		{
			name: "single transaction",
			transactions: []Transaction{
				{ID: 1, RS: []int{5}, WS: []int{7}},
			},
		},
		{
			name:         "no transactions",
			transactions: []Transaction{},
		},
		{
			name: "all independent",
			transactions: []Transaction{
				{ID: 1, RS: []int{1}, WS: []int{2}},
				{ID: 2, RS: []int{3}, WS: []int{4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectEmpty(tt.transactions)
			if len(got) != 0 {
				t.Fatalf("len = %d, want 0", len(got))
			}
			if hasConflict(tt.transactions, got) {
				t.Fatalf("empty sequence must not conflict")
			}
		})
	}
}

func Test_selectEmpty_maxSize(t *testing.T) {
	transactions := make([]Transaction, 100000)
	for i := 0; i < len(transactions); i++ {
		transactions[i] = Transaction{ID: i + 1, RS: []int{i + 1}, WS: []int{i + 1}}
	}
	got := selectEmpty(transactions)
	if len(got) != 0 {
		t.Fatalf("len = %d, want 0", len(got))
	}
}

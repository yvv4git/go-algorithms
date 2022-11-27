package main

import "testing"

func TestHeapEx01(t *testing.T) {
	graph := NewGraph()
	graph.AddEdge("BTC", "USD", 5)
	graph.AddEdge("USD", "ETH", 6)
}

package math

import "testing"

// Остаток от деления.
func TestDevisionReminder(t *testing.T) {
	t.Logf("10%%5 = %d", 10%5) // 0
	t.Logf("10%%3 = %d", 10%3) // 1
}

// Обычное деление.
func TestDevision(t *testing.T) {
	t.Logf("10/5 = %d", 10/5)                   // 2
	t.Logf("10/3 = %d", 10/3)                   // 3
	t.Logf("10/3 = %f", float32(10)/float32(3)) // 3.333333
}

package _02_happy_number

import "testing"

func TestResearch01(t *testing.T) {
	x := 15
	t.Logf("x %% 10 = %v", x%10) // 5 - остаток от деления
	t.Logf("x / 10 = %v", x/10)  // 1 - целое от деления
}

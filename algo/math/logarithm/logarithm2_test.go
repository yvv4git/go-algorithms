package logarithm

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Двоичный логарифм (логарифм по основанию 2).
func TestLogarithm2(t *testing.T) {
	t.Log("Test - logarithm 2")

	// Двоичный логарифм - log2(2) = 1
	// В какую степень надо возвести 2, чтобы получить значение в скобках.
	// log2(2) = 1
	// log2(4) = 2
	// log2(8) = 3
	// log2(16) = 4
	// log2(32) = 5
	testCases := []struct {
		InValue  float64
		Expected float64
	}{
		{1, 0},
		{2, 1},
		{4, 2},
		{8, 3},
		{16, 4},
		{32, 5},
	}

	for _, tc := range testCases {
		actualValue := math.Log2(tc.InValue)
		assert.Equal(t, tc.Expected, actualValue)
	}
}

// Логарифм по основанию 2 от n.
// Допустим мы у нас есть массив элементов n.
// Если я его буду делить по полам, то:
// step-0: n
// step-1: n/2
// step-2: n/4
// step-3: n/8
// ... и т.д.
// последняя итерация: 1, т.е. на последней итерации длина n=1(единице).

// Т.е. i-ая итерация: n/2^i
// Этот процесс идет до тех пор, пока не получим: 1 = n/2^i.
// Чтобы определиться i - т.е. итерацию, на которой n = 1 надо решить уравнение: 1 = n/2^i.
// Т.е. n = 2^i. Т.е. при возведении 2 в некоторую степень(i) мы получим n. ОГО! Где-то я уже видел такую зависимость.
// О, кажется я видел это в логарифмах.
func foo(n int, i int, t *testing.T) {
	nNew := n / 2
	i++ // i - это шаг(step)
	t.Logf("[step-%d] %d / 2 =: %d", i, n, nNew)
	if nNew <= 1 {
		return
	}
	foo(nNew, i, t)
	foo(nNew, i, t)
}

func TestLogarithm2_Division(t *testing.T) {
	n := 10
	i := 0
	t.Logf("[step-%d] n = %d", i, n)
	foo(n, i, t)

	// Количество уровней = log_2(n)
	// n = 10
	t.Log("Levels(steps) count(log_2(10)):", math.Log2(10)) // 3.321928094887362
}

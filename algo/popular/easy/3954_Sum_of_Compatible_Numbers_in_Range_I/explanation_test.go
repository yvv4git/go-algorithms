package main

import (
	"testing"
)

func TestExample1(t *testing.T) {
	/*
		ЗАДАЧА: Sum of Compatible Numbers in Range I

		Даны два числа n и k. Нужно найти сумму всех положительных x,
		которые одновременно удовлетворяют двум условиям:

		1. abs(n - x) <= k   — x лежит в диапазоне [n-k, n+k]
		2. (n & x) == 0      — побитовое AND между n и x равно нулю

		Побитовое AND (&) даёт 1 только в тех позициях, где оба числа имеют 1.
		Условие (n & x) == 0 означает, что у n и x не должно быть единиц
		на одних и тех же битовых позициях.
	*/

	/*
		ПРИМЕР 1: n = 2, k = 3

		Шаг 1. Диапазон:
		  L = max(1, 2-3) = 1
		  R = 2+3 = 5
		  Ищем x в [1, 5].

		Шаг 2. n = 2 -> 010 (bin, в трёх битах)

		Шаг 3. (n & x) == 0:
		  n:   0 1 0
		  x:   ? 0 ?    <- на позиции 1 можно ТОЛЬКО 0

		Шаг 4. Комбинации разрешённых битов:
		  0 0 0  -> 0  (нет, x < 1)
		  0 0 1  -> 1  (да)
		  1 0 0  -> 4  (да)
		  1 0 1  -> 5  (да)

		Шаг 5. Сумма = 1 + 4 + 5 = 10
	*/
	n, k := 2, 3
	got := sumOfCompatibleNumbers(n, k)
	want := 10

	t.Logf("n=%d k=%d", n, k)
	t.Logf("n=%b, диапазон [%d,%d]", n, max(1, n-k), n+k)
	t.Logf("Совместимые x: 1 (001), 4 (100), 5 (101)")
	t.Logf("Сумма = %d", got)

	if got != want {
		t.Errorf("sumOfCompatibleNumbers(%d,%d) = %d, want %d", n, k, got, want)
	}
}

func TestExample2(t *testing.T) {
	/*
		ПРИМЕР 2: n = 5, k = 1

		Шаг 1. Диапазон:
		  L = max(1, 5-1) = 4
		  R = 5+1 = 6
		  Ищем x в [4, 6].

		Шаг 2. n = 5 -> 101 (bin)

		Шаг 3. (n & x) == 0:
		  n:   1 0 1
		  x:   0 ? 0    <- на позициях 0 и 2 можно ТОЛЬКО 0

		Шаг 4. Кандидаты:
		  x=4 (100): 101 & 100 = 100 != 0 -> нет
		  x=5 (101): 101 & 101 = 101 != 0 -> нет
		  x=6 (110): 101 & 110 = 100 != 0 -> нет

		Шаг 5. Сумма = 0
	*/
	n, k := 5, 1
	got := sumOfCompatibleNumbers(n, k)
	want := 0

	t.Logf("n=%d k=%d", n, k)
	t.Logf("n=%b, диапазон [%d,%d]", n, max(1, n-k), n+k)
	t.Logf("Нет совместимых x в диапазоне")
	t.Logf("Сумма = %d", got)

	if got != want {
		t.Errorf("sumOfCompatibleNumbers(%d,%d) = %d, want %d", n, k, got, want)
	}
}

func TestExample3(t *testing.T) {
	/*
		ПРИМЕР 3: n = 1, k = 1

		Шаг 1. Диапазон: [max(1,0), 2] = [1, 2]

		Шаг 2. n = 1 -> 001 (bin)

		Шаг 3. (n & x) == 0:
		  n:   0 0 1
		  x:   ? ? 0    <- в позиции 0 можно ТОЛЬКО 0

		Шаг 4. Кандидаты:
		  x=1 (001): 001 & 001 = 001 != 0 -> нет
		  x=2 (010): 001 & 010 = 000 = 0, |1-2|=1 <= 1 -> да

		Шаг 5. Сумма = 2
	*/
	n, k := 1, 1
	got := sumOfCompatibleNumbers(n, k)
	want := 2

	t.Logf("n=%d k=%d", n, k)
	t.Logf("n=%b, диапазон [%d,%d]", n, max(1, n-k), n+k)
	t.Logf("x=2 (010) - единственный совместимый")
	t.Logf("Сумма = %d", got)

	if got != want {
		t.Errorf("sumOfCompatibleNumbers(%d,%d) = %d, want %d", n, k, got, want)
	}
}

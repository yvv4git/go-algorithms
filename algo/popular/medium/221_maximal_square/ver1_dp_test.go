package main

import "testing"

// Тест из примера 1: квадрат 2x2, площадь 4.
func TestMaximalSquare_Example1(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	expected := 4
	result := maximalSquare(matrix)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// Тест из примера 2: квадрат 1x1, площадь 1.
func TestMaximalSquare_Example2(t *testing.T) {
	matrix := [][]byte{
		{'0', '1'},
		{'1', '0'},
	}
	expected := 1
	result := maximalSquare(matrix)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// Тест из примера 3: только '0', площадь 0.
func TestMaximalSquare_Example3(t *testing.T) {
	matrix := [][]byte{
		{'0'},
	}
	expected := 0
	result := maximalSquare(matrix)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// Тест: пустая матрица.
func TestMaximalSquare_EmptyMatrix(t *testing.T) {
	matrix := [][]byte{}
	expected := 0
	result := maximalSquare(matrix)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// Тест: матрица с одним '1'.
func TestMaximalSquare_SingleOne(t *testing.T) {
	matrix := [][]byte{
		{'1'},
	}
	expected := 1
	result := maximalSquare(matrix)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

// Тест: матрица только с '0'.
func TestMaximalSquare_AllZeros(t *testing.T) {
	matrix := [][]byte{
		{'0', '0', '0'},
		{'0', '0', '0'},
		{'0', '0', '0'},
	}
	expected := 0
	result := maximalSquare(matrix)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

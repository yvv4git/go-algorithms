package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	xCoords := make([]int, n)
	yCoords := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &xCoords[i], &yCoords[i])
	}

	// Находим медианные координаты
	xMedian := findMedian(xCoords)
	yMedian := findMedian(yCoords)

	// Проверяем, не занята ли медианная точка существующим городом
	if !isCityExists(xCoords, yCoords, xMedian, yMedian) {
		fmt.Fprintln(writer, xMedian, yMedian)
		return
	}

	// Если медианная точка занята, ищем ближайшую свободную точку
	minAvg := math.MaxFloat64
	bestX, bestY := 0, 0

	// Проверяем точки в окрестности медианной точки
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

	for _, dir := range directions {
		for dist := 1; dist <= 10; dist++ {
			newX := xMedian + dir[0]*dist
			newY := yMedian + dir[1]*dist

			if !isCityExists(xCoords, yCoords, newX, newY) {
				avg := calculateAverageDistance(xCoords, yCoords, newX, newY)
				if avg < minAvg {
					minAvg = avg
					bestX, bestY = newX, newY
				}
			}
		}
	}

	fmt.Fprintln(writer, bestX, bestY)
}

func findMedian(coords []int) int {
	// Создаем копию для сортировки
	sorted := make([]int, len(coords))
	copy(sorted, coords)

	// Сортируем координаты
	for i := 0; i < len(sorted)-1; i++ {
		for j := i + 1; j < len(sorted); j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}

	// Возвращаем медиану
	return sorted[len(sorted)/2]
}

func isCityExists(xCoords, yCoords []int, x, y int) bool {
	for i := 0; i < len(xCoords); i++ {
		if xCoords[i] == x && yCoords[i] == y {
			return true
		}
	}
	return false
}

func calculateAverageDistance(xCoords, yCoords []int, x, y int) float64 {
	total := 0
	for i := 0; i < len(xCoords); i++ {
		total += abs(xCoords[i]-x) + abs(yCoords[i]-y)
	}
	return float64(total) / float64(len(xCoords))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

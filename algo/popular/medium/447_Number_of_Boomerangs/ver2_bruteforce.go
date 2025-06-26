//go:build ignore

package main

import "fmt"

/*
Approach: Brute Force
- Перебираем все возможные тройки точек (i, j, k)
- Для каждой тройки вычисляем расстояния от i до j и от i до k
- Если расстояния равны, увеличиваем счетчик бумерангов
- Используем квадрат расстояния для избежания операций с плавающей точкой

Time Complexity: O(n³) - три вложенных цикла для перебора всех троек
Space Complexity: O(1) - используем только константное количество дополнительной памяти
*/

// numberOfBoomerangs решает задачу с использованием брутфорс подхода
func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	if n < 3 {
		return 0
	}

	count := 0

	// Перебираем все возможные тройки точек (i, j, k)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			for k := 0; k < n; k++ {
				if i == k || j == k {
					continue
				}

				// Вычисляем квадраты расстояний от точки i до точек j и k
				distIJ := getDistanceSquared(points[i], points[j])
				distIK := getDistanceSquared(points[i], points[k])

				// Если расстояния равны, то найден бумеранг
				if distIJ == distIK {
					count++
				}
			}
		}
	}

	return count
}

// getDistanceSquared вычисляет квадрат расстояния между двумя точками
// Используем квадрат расстояния для избежания операций с плавающей точкой
func getDistanceSquared(point1, point2 []int) int {
	dx := point1[0] - point2[0]
	dy := point1[1] - point2[1]
	return dx*dx + dy*dy
}

func main() {
	// Тестовые случаи
	testCases := [][][]int{
		{{0, 0}, {1, 0}, {2, 0}}, // Expected: 2
		{{1, 1}, {2, 2}, {3, 3}}, // Expected: 2
		{{1, 1}},                 // Expected: 0
		{{0, 0}, {1, 0}, {-1, 0}, {0, 1}, {0, -1}}, // Expected: 20
	}

	expected := []int{2, 2, 0, 20}

	fmt.Println("=== Brute Force Approach ===")
	fmt.Println("Time Complexity: O(n³)")
	fmt.Println("Space Complexity: O(1)")
	fmt.Println()

	for i, points := range testCases {
		result := numberOfBoomerangs(points)
		fmt.Printf("Test %d: points = %v\n", i+1, points)
		fmt.Printf("Result: %d, Expected: %d\n", result, expected[i])
		if result == expected[i] {
			fmt.Println("✅ PASS")
		} else {
			fmt.Println("❌ FAIL")
		}
		fmt.Println()
	}
}

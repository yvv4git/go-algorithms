//go:build ignore

package main

import "fmt"

/*
APPROACH: "Алгоритм Бойера-Мура для поиска элемента большинства (majority element)"
- Идея: поддерживать кандидата и счетчик.
- Если счетчик равен 0, выбираем текущий элемент как кандидата.
- Если текущий элемент равен кандидату, увеличиваем счетчик, иначе уменьшаем.
- В конце кандидат будет искомым элементом, так как гарантируется его существование.

TIME COMPLEXITY: O(n)
SPACE COMPLEXITY: O(1)
*/
func majorityElement(nums []int) int {
	candidate := 0
	count := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))             // Output: 3
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // Output: 2
}

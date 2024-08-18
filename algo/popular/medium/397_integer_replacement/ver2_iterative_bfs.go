package main

import (
	"container/list"
)

// Функция integerReplacementV2 решает задачу "Integer Replacement" с использованием BFS
func integerReplacementV2(n int) int {
	/*
		METHOD: Breadth-First Search (BFS)

		DESCRIPTION:
		Этот метод использует BFS для поиска минимального количества операций.
		Мы используем очередь для хранения текущих значений и количества шагов, необходимых для достижения этих значений.
		Для каждого значения мы проверяем, можно ли его уменьшить или увеличить, и результат добавляем в очередь, если он еще не был посещен.

		TIME COMPLEXITY:
		Временная сложность этого метода составляет O(n), так как в худшем случае мы можем посетить все значения от 1 до n.

		SPACE COMPLEXITY:
		Пространственная сложность этого метода составляет O(n), так как в худшем случае мы можем хранить все значения от 1 до n в очереди и множестве visited.
	*/
	if n == 1 {
		return 0
	}

	queue := list.New()
	queue.PushBack(n)
	visited := make(map[int]bool)
	visited[n] = true
	steps := 0

	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			current := queue.Remove(queue.Front()).(int)

			if current == 1 {
				return steps
			}

			if current%2 == 0 {
				next := current / 2
				if !visited[next] {
					queue.PushBack(next)
					visited[next] = true
				}
			} else {
				next1 := current + 1
				next2 := current - 1
				if !visited[next1] {
					queue.PushBack(next1)
					visited[next1] = true
				}
				if !visited[next2] {
					queue.PushBack(next2)
					visited[next2] = true
				}
			}
		}
		steps++
	}

	return -1 // Этот случай никогда не должен произойти, если n > 0
}

//func main() {
//	// Пример использования функции integerReplacement
//	n := 8
//	fmt.Printf("Минимальное количество операций для преобразования %d в 1: %d\n", n, integerReplacement(n))
//}

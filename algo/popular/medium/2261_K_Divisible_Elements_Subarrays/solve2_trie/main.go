package main

import (
	"bufio"
	"fmt"
	"os"
)

type trieNode struct {
	children map[int]*trieNode
}

func countDistinctV2(nums []int, k int, p int) int {
	/*
		TASK: Посчитать количество различных подмассивов массива nums, в которых не более k элементов делятся на p.
		Подмассив - это непрерывная часть массива, заданная границами left и right.
		Два подмассива различны, если у них разная длина или хотя бы один элемент на одинаковой позиции отличается.
		Элемент x делится на p, если x % p == 0.
		Пример: nums = [2, 3, 3, 2, 2], k = 2, p = 2.
		Элементы 2, 2 и 2 делятся на 2, поэтому подмассив [2, 3, 3, 2] подходит, а [2, 3, 3, 2, 2] нет: в нём три таких элемента.
		Различных подходящих подмассивов здесь 11.
		Граничные случаи:
		k = 0: подходит только подмассив без элементов, делящихся на p, например [3] или [3, 3];
		k = n: подходит любой подмассив, и ответ равен числу различных подмассивов.
		Отрицательный пример: nums = [2, 2], k = 1, p = 2, подмассив [2, 2] содержит два делящихся элемента, поэтому он не учитывается, ответ равен 1.
		Функция возвращает количество различных подходящих подмассивов.

		METHOD: Trie (префиксное дерево)
		Trie - это дерево, в котором путь от корня задаёт последовательность значений.
		Каждая вершина хранит map children, где ключ - значение элемента, а значение - дочерняя вершина.
		Для каждой левой границы left начинаем с корня root и двигаем right вправо, поддерживая счётчик divisible.
		Если nums[right] % p == 0, увеличиваем divisible.
		Когда divisible > k, прекращаем перебор для текущей left: дальше счётчик только вырастет.
		Значение nums[right] вставляется как ребёнок текущей вершины.
		Если такого ребёнка не было, создаём новую вершину и увеличиваем distinct на 1.
		Каждая новая вершина соответствует новому различному подмассиву, поэтому ответ - число созданных вершин.

		TIME COMPLEXITY: O(n^2).
		Для каждой из n левых границ просматривается не более n правых границ, вставка в map занимает O(1) в среднем.

		SPACE COMPLEXITY: O(n^2).
		В дереве не более O(n^2) вершин, по одной на каждый различный подмассив.
	*/
	root := &trieNode{children: make(map[int]*trieNode)}
	distinct := 0

	for left := range nums {
		node := root
		divisible := 0

		for right := left; right < len(nums); right++ {
			if nums[right]%p == 0 {
				divisible++
			}
			if divisible > k {
				break
			}

			next, created := node.ensureChild(nums[right])
			if created {
				distinct++
			}
			node = next
		}
	}

	return distinct
}

func (node *trieNode) ensureChild(value int) (*trieNode, bool) {
	if next, exists := node.children[value]; exists {
		return next, false
	}

	next := &trieNode{children: make(map[int]*trieNode)}
	node.children[value] = next

	return next, true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, k, p int
	if _, err := fmt.Fscan(reader, &n, &k, &p); err != nil {
		fmt.Fprintln(os.Stderr, err)

		return
	}

	nums := make([]int, n)
	for index := range nums {
		if _, err := fmt.Fscan(reader, &nums[index]); err != nil {
			fmt.Fprintln(os.Stderr, err)

			return
		}
	}

	fmt.Fprintln(writer, countDistinctV2(nums, k, p))
}

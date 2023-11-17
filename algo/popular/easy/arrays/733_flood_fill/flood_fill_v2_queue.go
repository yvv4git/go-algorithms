package _33_flood_fill

func floodFillV2(image [][]int, sr int, sc int, color int) [][]int {
	/*
		Method: Queue
		Time complexity: O(n)
		Space complexity: O(1)

		Данный алгоритм представляет собой алгоритм покраски (flood fill) для изображений.
		Он использует очередь (queue) для хранения координат пикселей, которые нужно покрасить.

		Работа алгоритма:
		1. Проверяется, не является ли цвет пикселя, который нужно покрасить, уже целевым цветом.
		Если цвет пикселя уже целевым, то алгоритм завершается, так как ничего покрасить не нужно.
		2. Затем создается очередь (queue) и добавляется в нее координата пикселя, который нужно покрасить.
		3. Затем алгоритм входит в цикл, который продолжается, пока очередь не станет пустой.
		4. В каждой итерации алгоритм берет первый элемент из очереди, проверяет, является ли цвет этого пикселя начальным цветом.
		Если цвет пикселя не совпадает с начальным цветом, то алгоритм пропускает этот пиксель и переходит к следующей итерации.
		5. Если цвет пикселя совпадает с начальным цветом, то алгоритм меняет цвет этого пикселя на целевой цвет и добавляет в очередь всех его соседей.
		Соседними считаются пиксели, которые находятся сверху, снизу, слева и справа от текущего пикселя.
		6. После того, как все пиксели, которые можно покрасить, будут покрашены, алгоритм завершается и возвращает измененное изображение.
	*/
	if image[sr][sc] == color {
		return image
	}

	startColor := image[sr][sc]
	q := [][2]int{{sr, sc}}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		if image[v[0]][v[1]] != startColor {
			continue
		}
		image[v[0]][v[1]] = color
		if v[0] > 0 {
			q = append(q, [2]int{v[0] - 1, v[1]})
		}
		if v[0] < len(image)-1 {
			q = append(q, [2]int{v[0] + 1, v[1]})
		}
		if v[1] > 0 {
			q = append(q, [2]int{v[0], v[1] - 1})
		}
		if v[1] < len(image[0])-1 {
			q = append(q, [2]int{v[0], v[1] + 1})
		}
	}

	return image
}

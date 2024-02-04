package _33_flood_fill

func floodFillV5(image [][]int, sr int, sc int, color int) [][]int {
	/*
		METHOD: Union-Find.
		TIME COMPLEXITY: O(m*n)
		Space complexity: O(1)

		Сложность алгоритма зависит от размера входного изображения. Предположим, что изображение имеет размеры M x N.
		1. Создание Union-Find структуры: O(MN)
		2. Проход по всем пикселям и объединение пикселей с соседними пикселями, имеющими тот же цвет: O(MN)
		3. Проход по всем пикселям и замена цвета: O(MN)

		Суммарная асимптотическая сложность будет O(MN).
		Обратите внимание, что в реальном мире, где M и N могут быть очень большими числами,
		этот алгоритм может быть неэффективным, поскольку он может потребовать много памяти.
		В реальных приложениях, где требуется эффективное решение, могут использоваться другие алгоритмы,
		такие как алгоритм сканирующей строки для заливки.

		В этом коде мы используем Union-Find для объединения всех пикселей, которые имеют тот же цвет,
		что и начальный пиксель. Затем мы проходим по всем пикселям и заменяем цвет каждого пикселя на новый цвет,
		если он принадлежит той же группе, что и начальный пиксель.
	*/
	rows := len(image)
	cols := len(image[0])
	uf := NewUnionFind(rows * cols)
	startColor := image[sr][sc]
	image[sr][sc] = color
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if image[i][j] == startColor {
				if i > 0 && image[i-1][j] == startColor {
					uf.Union(i*cols+j, (i-1)*cols+j)
				}
				if i < rows-1 && image[i+1][j] == startColor {
					uf.Union(i*cols+j, (i+1)*cols+j)
				}
				if j > 0 && image[i][j-1] == startColor {
					uf.Union(i*cols+j, i*cols+j-1)
				}
				if j < cols-1 && image[i][j+1] == startColor {
					uf.Union(i*cols+j, i*cols+j+1)
				}
			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if uf.Find(i*cols+j) == uf.Find(sr*cols+sc) {
				image[i][j] = color
			}
		}
	}
	return image
}

type UnionFind struct {
	parent []int
	rank   []int
}

func NewUnionFind(size int) *UnionFind {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &UnionFind{parent, rank}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x int, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.parent[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.parent[rootX] = rootY
		} else {
			uf.parent[rootY] = rootX
			uf.rank[rootX] += 1
		}
	}
}

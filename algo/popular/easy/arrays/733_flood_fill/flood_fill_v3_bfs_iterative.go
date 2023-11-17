package _33_flood_fill

func floodFillV3(image [][]int, sr int, sc int, color int) [][]int {
	/*
		Method: BFS iterative.
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	type Point struct {
		x, y int
	}

	if image[sr][sc] == color {
		return image
	}

	startColor := image[sr][sc]
	q := []Point{{sr, sc}}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]

		if image[v.x][v.y] != startColor {
			continue
		}
		image[v.x][v.y] = color
		if v.x > 0 {
			q = append(q, Point{v.x - 1, v.y})
		}
		if v.x < len(image)-1 {
			q = append(q, Point{v.x + 1, v.y})
		}
		if v.y > 0 {
			q = append(q, Point{v.x, v.y - 1})
		}
		if v.y < len(image[0])-1 {
			q = append(q, Point{v.x, v.y + 1})
		}
	}

	return image
}

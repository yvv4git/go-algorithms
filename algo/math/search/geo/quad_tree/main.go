package main

import "fmt"

// Point represents a point in 2D space.
type Point struct {
	X, Y float64
}

// BoundingBox represents a bounding box in 2D space.
type BoundingBox struct {
	Center Point
	Width  float64
	Height float64
}

// QuadTree represents a QuadTree.
type QuadTree struct {
	BoundingBox BoundingBox
	Points      []Point
	Children    [4]*QuadTree
}

// NewQuadTree creates a new QuadTree with the given bounding box.
// Time complexity: O(1)
// Space complexity: O(1)
func NewQuadTree(bb BoundingBox) *QuadTree {
	return &QuadTree{
		BoundingBox: bb,
		Points:      make([]Point, 0),
	}
}

// Insert inserts a point into the QuadTree.
// Time complexity: O(log n) in the average case, O(n) in the worst case
// Space complexity: O(1)
func (qt *QuadTree) Insert(p Point) {
	if !qt.BoundingBox.Contains(p) {
		return
	}

	if len(qt.Points) < 4 || qt.Children[0] == nil {
		qt.Points = append(qt.Points, p)
	} else {
		if qt.Children[0] == nil {
			qt.subdivide()
		}
		for _, child := range qt.Children {
			child.Insert(p)
		}
	}
}

// subdivide divides the QuadTree into four sub-quadtrees.
// Time complexity: O(1)
// Space complexity: O(1)
func (qt *QuadTree) subdivide() {
	halfWidth := qt.BoundingBox.Width / 2
	halfHeight := qt.BoundingBox.Height / 2

	qt.Children[0] = NewQuadTree(BoundingBox{
		Center: Point{qt.BoundingBox.Center.X - halfWidth/2, qt.BoundingBox.Center.Y - halfHeight/2},
		Width:  halfWidth,
		Height: halfHeight,
	})
	qt.Children[1] = NewQuadTree(BoundingBox{
		Center: Point{qt.BoundingBox.Center.X + halfWidth/2, qt.BoundingBox.Center.Y - halfHeight/2},
		Width:  halfWidth,
		Height: halfHeight,
	})
	qt.Children[2] = NewQuadTree(BoundingBox{
		Center: Point{qt.BoundingBox.Center.X - halfWidth/2, qt.BoundingBox.Center.Y + halfHeight/2},
		Width:  halfWidth,
		Height: halfHeight,
	})
	qt.Children[3] = NewQuadTree(BoundingBox{
		Center: Point{qt.BoundingBox.Center.X + halfWidth/2, qt.BoundingBox.Center.Y + halfHeight/2},
		Width:  halfWidth,
		Height: halfHeight,
	})
}

// Contains checks if the bounding box contains the point.
// Time complexity: O(1)
// Space complexity: O(1)
func (bb BoundingBox) Contains(p Point) bool {
	return p.X >= bb.Center.X-bb.Width/2 &&
		p.X <= bb.Center.X+bb.Width/2 &&
		p.Y >= bb.Center.Y-bb.Height/2 &&
		p.Y <= bb.Center.Y+bb.Height/2
}

func main() {
	bb := BoundingBox{
		Center: Point{0, 0},
		Width:  200,
		Height: 200,
	}
	qt := NewQuadTree(bb)

	qt.Insert(Point{10, 10})
	qt.Insert(Point{100, 100})
	qt.Insert(Point{150, 75})

	fmt.Println(qt)
}

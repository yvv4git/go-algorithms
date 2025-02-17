package main

import (
	"fmt"
	"strings"
)

// Rectangle определяет границы квадранта
type Rectangle struct {
	X, Y, Width, Height float64
}

// Point представляет точку в пространстве
type Point struct {
	X, Y float64
	Data interface{}
}

// QuadTree узел
type QuadTree struct {
	Bounds         Rectangle // Границы текущего узла (координаты и размеры области)
	Capacity       int       // Максимальное количество точек в узле до разбиения
	Points         []Point   // Список точек, находящихся в данном узле
	Divided        bool      // Флаг, показывающий, был ли узел уже разбит на подузлы
	NE, NW, SE, SW *QuadTree // Ссылки на четыре дочерних квадранта (северо-восточный, северо-западный, юго-восточный и юго-западный)
}

// NewQuadTree создаёт новый QuadTree
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func NewQuadTree(bounds Rectangle, capacity int) *QuadTree {
	return &QuadTree{
		Bounds:   bounds,
		Capacity: capacity,
		Points:   []Point{},
		Divided:  false,
	}
}

// Insert добавляет точку в QuadTree
// Временная сложность: O(log n) в среднем, O(n) в худшем случае
// Пространственная сложность: O(n)
func (qt *QuadTree) Insert(p Point) bool {
	if !qt.contains(p) {
		return false
	}

	if len(qt.Points) < qt.Capacity {
		qt.Points = append(qt.Points, p)
		return true
	}

	if !qt.Divided {
		qt.subdivide()
	}

	return qt.NE.Insert(p) || qt.NW.Insert(p) || qt.SE.Insert(p) || qt.SW.Insert(p)
}

// Проверка, содержится ли точка в границах
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (qt *QuadTree) contains(p Point) bool {
	return p.X >= qt.Bounds.X && p.X <= qt.Bounds.X+qt.Bounds.Width &&
		p.Y >= qt.Bounds.Y && p.Y <= qt.Bounds.Y+qt.Bounds.Height
}

// Разбиение узла на 4 квадранта
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (qt *QuadTree) subdivide() {
	hw := qt.Bounds.Width / 2
	hh := qt.Bounds.Height / 2
	x := qt.Bounds.X
	y := qt.Bounds.Y

	qt.NE = NewQuadTree(Rectangle{x + hw, y, hw, hh}, qt.Capacity)
	qt.NW = NewQuadTree(Rectangle{x, y, hw, hh}, qt.Capacity)
	qt.SE = NewQuadTree(Rectangle{x + hw, y + hh, hw, hh}, qt.Capacity)
	qt.SW = NewQuadTree(Rectangle{x, y + hh, hw, hh}, qt.Capacity)
	qt.Divided = true
}

// Вывод QuadTree в консоль
// Временная сложность: O(n)
// Пространственная сложность: O(1)
func (qt *QuadTree) Print(depth int) {
	fmt.Printf("%sBounds: %+v, Points: %d\n", strings.Repeat(" ", depth*2), qt.Bounds, len(qt.Points))
	if qt.Divided {
		qt.NE.Print(depth + 1)
		qt.NW.Print(depth + 1)
		qt.SE.Print(depth + 1)
		qt.SW.Print(depth + 1)
	}
}

func main() {
	qt := NewQuadTree(Rectangle{0, 0, 10, 10}, 4)
	points := []Point{{2, 3, "A"}, {5, 5, "B"}, {8, 8, "C"}, {9, 9, "D"}, {4, 4, "E"}}

	for _, p := range points {
		qt.Insert(p)
	}

	qt.Print(0)
}

# 836. Rectangle Overlap


## Level - easy


## Task
An axis-aligned rectangle is represented as a list [x1, y1, x2, y2], where (x1, y1) is the coordinate of its bottom-left corner, and (x2, y2) is the coordinate of its top-right corner. 
Its top and bottom edges are parallel to the X-axis, and its left and right edges are parallel to the Y-axis.

Two rectangles overlap if the area of their intersection is positive. To be clear, two rectangles that only touch at the corner or edges do not overlap.

Given two axis-aligned rectangles rec1 and rec2, return true if they overlap, otherwise return false.


## Объяснение
Даны два прямоугольника, представленные в виде списков или массивов. 
Каждый прямоугольник определяется четырьмя числами: [x1, y1, x2, y2], где (x1, y1) — координаты левого нижнего угла, а (x2, y2) — координаты правого верхнего угла. 
Прямоугольники расположены на двумерной плоскости, и их стороны параллельны осям координат.

Необходимо определить, пересекаются ли два прямоугольника. Если они пересекаются, то вернуть true, иначе — false.

```
Вход: rec1 = [0, 0, 2, 2], rec2 = [1, 1, 3, 3]
Выход: true
Объяснение: Прямоугольники пересекаются, так как есть общая область.
```


## Example 1:
```
Input: rec1 = [0,0,2,2], rec2 = [1,1,3,3]
Output: true
```


## Example 2:
```
Input: rec1 = [0,0,1,1], rec2 = [1,0,2,1]
Output: false
```

## Example 3:
```
Input: rec1 = [0,0,1,1], rec2 = [2,2,3,3]
Output: false
```


## Constraints:
- rec1.length == 4
- rec2.length == 4
- -10^9 <= rec1[i], rec2[i] <= 10^9
- rec1 and rec2 represent a valid rectangle with a non-zero area.
# 1037. Valid Boomerang


## Level - easy


## Task
Given an array points where points[i] = [xi, yi] represents a point on the X-Y plane, 
return true if these points are a boomerang.

A boomerang is a set of three points that are all distinct and not in a straight line.

## Объяснение
Задача состоит в том, чтобы определить, является ли заданный треугольник бумерангом или нет. 
Бумеранг - это треугольник, у которого все углы не равны 180 градусов.

Бумеранг может быть определен по трем точкам, которые даны в виде координат. 
Задача требует проверить, не лежат ли эти точки на одной прямой, то есть является ли треугольник, 
образованный этими точками, бумерангом.

Функция, которая решает эту задачу, должна принимать в качестве аргумента массив из трех элементов, 
каждый из которых является массивом из двух целых чисел, представляющих собой координаты точки. 
Функция должна возвращать булево значение true, если треугольник является бумерангом, и false, если нет.

Примеры:
- Для входных данных [[1,1],[2,3],[3,2]] функция должна вернуть true, потому что эти точки не лежат на одной прямой, 
и образуют бумеранг.

- Для входных данных [[1,1],[2,2],[3,3]] функция должна вернуть false, потому что эти точки лежат на одной прямой, 
и не образуют бумеранг.

## Example 1:
````
Input: points = [[1,1],[2,3],[3,2]]
Output: true
````


## Example 2:
````
Input: points = [[1,1],[2,2],[3,3]]
Output: false
````


## Constraints:
- points.length == 3
- points[i].length == 2
- 0 <= xi, yi <= 100
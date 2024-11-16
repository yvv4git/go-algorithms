# 976. Largest Perimeter Triangle


## Level - easy


## Task
Given an integer array nums, return the largest perimeter of a triangle with a non-zero area, 
formed from three of these lengths. If it is impossible to form any triangle of a non-zero area, return 0.


## Объяснение
Дан массив целых чисел nums, где каждое число представляет длину отрезка. 
Необходимо найти наибольший периметр треугольника, который можно сформировать из трех различных отрезков, выбранных из массива. 
Если такой треугольник невозможно сформировать, то вернуть 0.

Пример:
```
Вход: nums = [2,1,2]
Выход: 5
Объяснение: Можно сформировать треугольник со сторонами 2, 1, 2. Периметр равен 2 + 1 + 2 = 5
```

Пример-2:
```
Вход: nums = [1,2,1]
Выход: 0
Объяснение: Невозможно сформировать треугольник из отрезков 1, 2, 1, 
так как не выполняется неравенство треугольника.
```

Неравенство треугольника:
- a + b > c
- a + c > b
- b + c > a

где a, b, и c — длины сторон треугольника.

Чтобы найти наибольший периметр, можно отсортировать массив по убыванию. 
Затем проверить тройки соседних элементов, начиная с самого большого, и проверить, могут ли они образовать треугольник. 
Первая тройка, которая удовлетворяет условию, даст наибольший периметр.


## Example 1:
```
Input: nums = [2,1,2]
Output: 5
Explanation: You can form a triangle with three side lengths: 1, 2, and 2.
```

## Example 2:
```
Input: nums = [1,2,1,10]
Output: 0
Explanation: 
You cannot use the side lengths 1, 1, and 2 to form a triangle.
You cannot use the side lengths 1, 1, and 10 to form a triangle.
You cannot use the side lengths 1, 2, and 10 to form a triangle.
As we cannot use any three side lengths to form a triangle of non-zero area, we return 0.
```


## Constraints:
- 3 <= nums.length <= 10^4
- 1 <= nums[i] <= 10^6

# 832. Flipping an Image


## Level - easy


## Task
Given an n x n binary matrix image, flip the image horizontally, then invert it, and return the resulting image.

To flip an image horizontally means that each row of the image is reversed.
- For example, flipping [1,1,0] horizontally results in [0,1,1].

To invert an image means that each 0 is replaced by 1, and each 1 is replaced by 0.
- For example, inverting [0,1,1] results in [1,0,0].


## Объяснение
Задача требует от вас выполнить две операции над бинарной матрицей (изображением), представленной в виде двумерного массива:
- Отражение (Flipping).
Каждая строка матрицы должна быть перевернута. Это означает, что первый элемент строки становится последним, второй — предпоследним и так далее.

- Инверсия (Inverting). 
После отражения каждый элемент в матрице должен быть инвертирован. 
Для бинарного изображения это означает, что все 0 становятся 1, а все 1 становятся 0.

Например, вход:
```
[[1,1,0],
 [1,0,1],
 [0,0,0]]
```

Отражение каждой строки:
```
[[0,1,1],
 [1,0,1],
 [0,0,0]]
```

Инверсия каждого элемента:
```
[[1,0,0],
 [0,1,0],
 [1,1,1]]
```

Выход:
```
[[1,0,0],
 [0,1,0],
 [1,1,1]]
```

Для решения задачи надо:
Для решения этой задачи можно использовать простой подход:
1. Пройтись по каждой строке матрицы.
2. Перевернуть каждую строку.
3. Инвертировать каждый элемент в перевернутой строке.


## Example 1:
```
Input: image = [[1,1,0],[1,0,1],[0,0,0]]
Output: [[1,0,0],[0,1,0],[1,1,1]]
Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
```


## Example 2:
```
Input: image = [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
Explanation: First reverse each row: [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
```


## Constraints:
- n == image.length
- n == image[i].length
- 1 <= n <= 20
- images[i][j] is either 0 or 1.
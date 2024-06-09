# 11. Container With Most Water


## Level - medium


## Task
You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.


## Объяснение
Задача заключается в нахождении максимального объема воды, которую можно уложить между двумя линиями, 
представленными в виде массива высот. Каждая линия имеет определенную высоту, и ваша задача - найти две линии, которые, 
когда они уложены между собой, образуют контейнер с максимально возможным объемом воды.

Для решения этой задачи можно использовать алгоритм с двумя указателями. 
Один указатель начинается с начала массива, а второй - с конца. 
Затем вычисляется объем воды, который можно уложить между двумя линиями, и обновляется максимальный объем. 
Если высота левой линии меньше высоты правой, то левый указатель сдвигается вправо, иначе правый указатель сдвигается влево. 
Процесс продолжается до тех пор, пока указатели не встретятся.


## Example 1:
![img.png](img.png)
````
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
````


## Example 2:
````
Input: height = [1,1]
Output: 1
````


## Constraints:
- n == height.length
- 2 <= n <= 10^5
- 0 <= height[i] <= 10^4
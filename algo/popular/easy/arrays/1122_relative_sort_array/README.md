# 1122. Relative Sort Array


## Level - easy


## Task
Given two arrays arr1 and arr2, the elements of arr2 are distinct, and all elements in arr2 are also in arr1.

Sort the elements of arr1 such that the relative ordering of items in arr1 are the same as in arr2. 
Elements that do not appear in arr2 should be placed at the end of arr1 in ascending order.


## Объяснение
Задача предлагает нам отсортировать массив arr1 в соответствии с порядком элементов в массиве arr2, 
при этом оставляя все элементы, которые не встречаются в arr2, в исходном порядке.

Таким образом, задача состоит в том, чтобы создать новый отсортированный массив, основываясь на порядке элементов в arr2, 
а также сохранить порядок оставшихся элементов, которые не указаны в arr2.

Например, если у нас есть arr1 = [2,3,1,3,2,4,6,7,9,2,19] и arr2 = [2,1,4,3,9], 
то отсортированный массив должен быть [2,2,2,1,4,3,3,9,6,7,19].

Обратите внимание, что элементы, которые не указаны в arr2, должны быть отсортированы в порядке возрастания, 
но после всех элементов, которые указаны в arr2.


## Example 1:
````
Input: arr1 = [2,3,1,3,2,4,6,7,9,2,19], arr2 = [2,1,4,3,9,6]
Output: [2,2,2,1,4,3,3,9,6,7,19]
````


## Example 2:
````
Input: arr1 = [28,6,22,8,44,17], arr2 = [22,28,8,6]
Output: [22,28,8,6,17,44]
````


## Constraints:
- 1 <= arr1.length, arr2.length <= 1000
- 0 <= arr1[i], arr2[i] <= 1000
- All the elements of arr2 are distinct.
- Each arr2[i] is in arr1.
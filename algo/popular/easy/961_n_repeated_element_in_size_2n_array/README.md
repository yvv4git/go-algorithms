# 961. N-Repeated Element in Size 2N Array


## Level - easy


## Task
You are given an integer array nums with the following properties:
- nums.length == 2 * n.
- nums contains n + 1 unique elements.
- Exactly one element of nums is repeated n times.

Return the element that is repeated n times.


## Объяснение
Дан массив nums, содержащий 2N элементов. 
В этом массиве ровно один элемент повторяется N раз, а остальные элементы уникальны. 
Необходимо найти и вернуть этот повторяющийся элемент.

Пример-1:
- Вход: nums = [1,2,3,3]
- Выход: 3
- Объяснение: Элемент 3 повторяется 2 раза, что соответствует условию, так как массив содержит 4 элемента (2N = 4).

Пример-2:
- Вход: nums = [2,1,2,5,3,2]
- Выход: 2
- Объяснение: Элемент 2 повторяется 3 раза, что соответствует условию, так как массив содержит 6 элементов (2N = 6).

Пример-3:
- Вход: nums = [5,1,5,2,5,3,5,4]
- Выход: 5
- Объяснение: Элемент 5 повторяется 4 раза, что соответствует условию, так как массив содержит 8 элементов (2N = 8).


## Example 1:
```
Input: nums = [1,2,3,3]
Output: 3
```


## Example 2:
```
Input: nums = [2,1,2,5,3,2]
Output: 2
```


## Example 3:
```
Input: nums = [5,1,5,2,5,3,5,4]
Output: 5
```


## Constraints:
- 2 <= n <= 5000
- nums.length == 2 * n
- 0 <= nums[i] <= 10^4
- nums contains n + 1 unique elements and one of them is repeated exactly n times.
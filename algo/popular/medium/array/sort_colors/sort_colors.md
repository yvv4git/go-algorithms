# 75. Sort Colors


## Task
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, 
with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
You must solve this problem without using the library's sort function.


## Суть задачи
Надо отсортировать массив самостоятельно, без использования пакета sort.
Тут можно пойти разными способами:
- использовать алгоритм сортировки Selection sort, это займет по времени O(n^2).
- использовать алгоритм сортировки Counting sort, это займет по времени O(n + k), где k - это количество символов в словаре. Но по памяти будут затраты порядка (n + k).
- использовать какую-то хитрую хитрость типа как в примере v1.

## Constraints:
- n == nums.length
- 1 <= n <= 300
- nums[i] is either 0, 1, or 2.


## Example 1:
``
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
``


## Example 2:
``
Input: nums = [2,0,1]
Output: [0,1,2]
``
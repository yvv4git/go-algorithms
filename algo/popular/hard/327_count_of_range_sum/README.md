# 327. Count of Range Sum


## Level - medium


## Task
Given an integer array nums and two integers lower and upper, return the number of range sums that lie in [lower, upper] inclusive.

Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j inclusive, where i <= j.


## Объяснение
Дан массив целых чисел nums и два целых числа lower и upper. 
Требуется найти количество подмассивов, сумма элементов которых находится в диапазоне [lower, upper].

Формально, нужно найти количество пар индексов (i, j) таких, что lower ≤ sum(nums[i:j+1]) ≤ upper.

Пример:
- Вход: nums = [-2,5,-1], lower = -2, upper = 2
- Выход: 3
- Объяснение: Подмассивы с суммой в диапазоне [-2, 2] это: [-2], [5, -1], [-2, 5, -1].


Для решения этой задачи можно использовать метод префиксных сумм 
и модифицированное бинарное дерево поиска (например, дерево Фенвика или дерево отрезков).


## Example 1:
````
Input: nums = [-2,5,-1], lower = -2, upper = 2
Output: 3
Explanation: The three ranges are: [0,0], [2,2], and [0,2] and their respective sums are: -2, -1, 2.
````


## Example 2:
````
Input: nums = [0], lower = 0, upper = 0
Output: 1
````


## Constraints:
- 1 <= nums.length <= 10^5
- -2^31 <= nums[i] <= 2^31 - 1
- -10^5 <= lower <= upper <= 10^5
- The answer is guaranteed to fit in a 32-bit integer.
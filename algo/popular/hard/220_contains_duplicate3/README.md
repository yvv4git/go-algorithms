# 220. Contains Duplicate III


## Level - hard


## Task
You are given an integer array nums and two integers indexDiff and valueDiff.

Find a pair of indices (i, j) such that:
- i != j,
- abs(i - j) <= indexDiff.
- abs(nums[i] - nums[j]) <= valueDiff, and

Return true if such pair exists or false otherwise.


## Объяснение
Задача требует от нас определить, есть ли в массиве чисел nums два различных индекса i и j, 
таких что разность между nums[i] и nums[j] не превышает заданного числа k, 
а разность между индексами i и j не превышает заданного числа t.

Формально, задача формулируется следующим образом:
Дан массив целых чисел nums и два целых числа k и t. Необходимо определить, существуют ли индексы i и j такие, что:
- |i - j| <= k
- |nums[i] - nums[j]| <= t

Если такие индексы существуют, то нужно вернуть true, в противном случае — false.

````
Input: nums = [1,2,3,1], k = 3, t = 0
Output: true
Объяснение: nums[0] = 1 и nums[3] = 1, разность между индексами 3 - 0 = 3, что не превышает k = 3, и разность между числами |1 - 1| = 0, что не превышает t = 0.
````

````
Input: nums = [1,5,9,1,5,9], k = 2, t = 3
Output: false
Объяснение: Нет таких индексов i и j, чтобы выполнялись оба условия.
````

Для решения этой задачи можно использовать структуру данных, которая позволяет эффективно находить элементы в заданном диапазоне, 
например, TreeSet в Java или SortedSet в Python. 
Эти структуры данных позволяют быстро находить ближайшие элементы к заданному значению и проверять условие |nums[i] - nums[j]| <= t.


## Example 1:
````
Input: nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
Output: true
Explanation: We can choose (i, j) = (0, 3).
We satisfy the three conditions:
i != j --> 0 != 3
abs(i - j) <= indexDiff --> abs(0 - 3) <= 3
abs(nums[i] - nums[j]) <= valueDiff --> abs(1 - 1) <= 0
````


## Example 2:
````
Input: nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3
Output: false
Explanation: After trying all the possible pairs (i, j), we cannot satisfy the three conditions, so we return false.
````


## Constraints:
- 2 <= nums.length <= 10^5
- -10^9 <= nums[i] <= 10^9
- 1 <= indexDiff <= nums.length
- 0 <= valueDiff <= 10^9
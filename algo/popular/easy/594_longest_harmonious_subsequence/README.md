# 594. Longest Harmonious Subsequence


## Level - easy


## Task
We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.

Given an integer array nums, return the length of its longest harmonious subsequence among all its possible subsequences.


## Объяснение
Дан массив целых чисел nums. Подпоследовательность называется гармонической, 
если разница между любыми двумя элементами в подпоследовательности равна 1. 
Найдите длину самой длинной гармонической подпоследовательности в массиве nums.

Примеры:
- nums = [1, 3, 2, 2, 5, 2, 3, 7]: Самая длинная гармоническая подпоследовательность - [2, 2, 2, 3, 3], поэтому ответ - 5.
- nums = [1, 2, 3, 4]: Самая длинная гармоническая подпоследовательность - [1, 2], поэтому ответ - 2.
- nums = [1, 1, 1, 1]: Самая длинная гармоническая подпоследовательность - [1, 1, 1, 1], поэтому ответ - 4.


## Example 1:
````
Input: nums = [1,3,2,2,5,2,3,7]
Output: 5

Explanation:
The longest harmonious subsequence is [3,2,2,2,3].
````


## Example 2:
````
Input: nums = [1,2,3,4]
Output: 2

Explanation:
The longest harmonious subsequences are [1,2], [2,3], and [3,4], all of which have a length of 2.
````


## Example 3:
````
Input: nums = [1,1,1,1]
Output: 0

Explanation:
No harmonic subsequence exists.
````


## Constraints:
- 1 <= nums.length <= 2 * 10^4
- -10^9 <= nums[i] <= 10^9
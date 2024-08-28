# 377. Combination Sum IV


## Level - medium


## Task
Given an array of distinct integers nums and a target integer target, return the number of possible combinations that add up to target.

The test cases are generated so that the answer can fit in a 32-bit integer.


## Объяснение
Задача является классической задачей динамического программирования и комбинаторики.
Дан массив уникальных целых чисел nums и целевое целое число target. 
Требуется найти количество различных комбинаций из чисел массива nums, которые в сумме дают target. 
Каждое число из массива nums может быть использовано несколько раз в одной комбинации.


Пример:
Вход: nums = [1, 2, 3], target = 4

Выход: 7

Объяснение:
Все возможные комбинации, которые в сумме дают 4:
- (1, 1, 1, 1)
- (1, 1, 2)
- (1, 2, 1)
- (1, 3)
- (2, 1, 1)
- (2, 2)
- (3, 1)

Для решения этой задачи можно использовать динамическое программирование. 
Создадим массив dp, где dp[i] будет представлять количество комбинаций, которые в сумме дают i.

## Example 1:
````
Input: nums = [1,2,3], target = 4
Output: 7
Explanation:
The possible combination ways are:
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
Note that different sequences are counted as different combinations.
````


## Example 2:
````
Input: nums = [9], target = 3
Output: 0
````


## Constraints:
- 1 <= nums.length <= 200
- 1 <= nums[i] <= 1000
- All the elements of nums are unique.
- 1 <= target <= 1000


## Follow up
What if negative numbers are allowed in the given array? How does it change the problem? 
What limitation we need to add to the question to allow negative numbers?
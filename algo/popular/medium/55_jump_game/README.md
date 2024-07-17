# 55. Jump Game


## Level - medium


## Task
You are given an integer array nums. You are initially positioned at the array's first index, 
and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.


## Объяснение
Вам дан массив целых чисел nums. Начиная с первого индекса, каждый элемент в массиве представляет максимальное количество шагов, 
которые вы можете сделать вперед от этого индекса. Ваша задача — определить, можете ли вы достичь последнего индекса массива.

Например:
- Для nums = [2,3,1,1,4] ответ true, потому что вы можете начать с индекса 0, сделать 1 шаг к индексу 1, затем 3 шага к индексу 4.

- Для nums = [3,2,1,0,4] ответ false, потому что независимо от того, сколько шагов вы делаете, 
вы никогда не сможете перепрыгнуть индекс 3, где значение равно 0.

## Example 1:
````
Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
````


## Example 2:
````
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
````


## Constraints:
- 1 <= nums.length <= 10^4
- 0 <= nums[i] <= 10^5
# 137. Single Number II


## Level - medium


## Task
Given an integer array nums where every element appears three times except for one, which appears exactly once. 
Find the single element and return it.

You must implement a solution with a linear runtime complexity and use only constant extra space.


## Объяснение
Задача предлагает найти элемент, который встречается только один раз в массиве, 
при условии, что каждый элемент встречается три раза, кроме одного.

Для решения этой задачи можно использовать хеш-таблицу для подсчета количества вхождений каждого элемента, 
но это неэффективно, так как требует дополнительной памяти.

Более эффективный подход - использовать битовые операции. 
Вы можете представить каждый элемент в виде 32-битного числа, 
где каждый бит представляет собой сумму соответствующих битов во всех числах. 
Таким образом, если элемент встречается три раза, его биты будут сброшены (становятся нулями), 
а если встречается один раз, они остаются в единице.


## Example 1:
````
Input: nums = [2,2,3,2]
Output: 3
````


## Example 2:
````
Input: nums = [0,1,0,1,0,1,99]
Output: 99
````


## Constraints:
- 1 <= nums.length <= 3 * 10^4
- -2^31 <= nums[i] <= 231 - 1
- Each element in nums appears exactly three times except for one element which appears once.
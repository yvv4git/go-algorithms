# 645. Set Mismatch


## Level - easy


## Task
You have a set of integers s, which originally contains all the numbers from 1 to n. 
Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, 
which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.


## Суть задачи
Дано: массив чисел от 1 до n, числа не повторяются, идут по порядку: 1,2,3,4,5...n.
Одно из чисел повторяется, вроде такого 1,2,2,4, т.е. 3 пропущена, а двойка повторилась.
Надо вернуть результат, который будет содержать 2 числа:
- которое повторилось
- которое пропущено




## Constraints:
- 2 <= nums.length <= 10^4
- 1 <= nums[i] <= 10^4


## Example 1:
````
Input: nums = [1,2,2,4]
Output: [2,3]
````


## Example 2:
````
Input: nums = [1,1]
Output: [1,2]
````
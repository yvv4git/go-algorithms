# 45. Jump Game II


## Level - medium


## Task
You are given a 0-indexed array of integers nums of length n. You are initially positioned at nums[0].

Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at nums[i], you can jump to any nums[i + j] where:
- 0 <= j <= nums[i] and
- i + j < n

Return the minimum number of jumps to reach nums[n - 1]. The test cases are generated such that you can reach nums[n - 1].


## Объяснение
Нам дан массив целых чисел nums, где каждый элемент nums[i] представляет собой максимальную длину прыжка из позиции i. 
Наша задача — найти минимальное количество прыжков, необходимых для достижения последнего элемента массива, 
начиная с первого элемента.

Например, если у вас есть массив nums = [2,3,1,1,4], то:
- Из позиции 0 вы можете прыгнуть на 1 или 2 позиции вперед.
- Если вы прыгнете на 1 позицию, вы окажетесь на позиции 1, где значение равно 3, что позволяет вам прыгнуть на 3 позиции вперед и достичь последнего элемента.
- Если вы прыгнете на 2 позиции, вы сразу окажетесь на позиции 2, где значение равно 1, что позволяет вам прыгнуть на 1 позицию вперед и достичь последнего элемента.

В этом примере минимальное количество прыжков, необходимых для достижения последнего элемента, равно 2.

Формально задача формулируется так:
- Напишите функцию, которая принимает массив nums и возвращает минимальное количество прыжков, необходимых для достижения последнего элемента.
- Если массив состоит из одного элемента, то количество прыжков равно 0, так как вы уже находитесь на последнем элементе.

Эта задача требует от вас найти оптимальное решение, используя алгоритмы динамического программирования, 
жадные алгоритмы или другие подходы для минимизации количества прыжков.


## Example 1:
````
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
````


## Example 2:
````
Input: nums = [2,3,0,1,4]
Output: 2
````


## Constraints:
- 1 <= nums.length <= 10^4
- 0 <= nums[i] <= 1000
- It's guaranteed that you can reach nums[n - 1].
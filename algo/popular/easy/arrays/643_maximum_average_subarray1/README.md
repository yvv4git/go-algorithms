# 643. Maximum Average Subarray I


## Level - easy


## Task
You are given an integer array nums consisting of n elements, and an integer k.

Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.


## Объяснение
Это задача на поиск максимального среднего значения подмассива заданной длины в заданном массиве.

Условие задачи:  
Дан массив целых чисел nums и целое число k, где k - длина подмассива, который мы будем исследовать. 
Наша задача - найти подмассив длины k, у которого среднее значение максимально.

Например, если у нас есть массив nums = [1, 12, -5, -6, 50, 3] и k = 4, 
то подмассив [12, -5, -6, 50] имеет максимальное среднее значение 12.75.

Задача решается с помощью одного прохода по массиву, используя сложность O(n), 
где n - количество элементов в массиве.


## Example 1:
````
Input: nums = [1,12,-5,-6,50,3], k = 4
Output: 12.75000
Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
````


## Example 2:
````
Input: nums = [5], k = 1
Output: 5.00000
````

Constraints:
- n == nums.length
- 1 <= k <= n <= 10^5
- -10^4 <= nums[i] <= 10^4

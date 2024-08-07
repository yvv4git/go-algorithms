# 628. Maximum Product of Three Numbers


## Level = easy


## Task
Given an integer array nums, find three numbers whose product is maximum and return the maximum product.


## Объяснение
Задача предлагает найти максимальное произведение трех чисел из массива.

Массив содержит целые числа, и ваша задача - найти три числа, произведение которых будет максимальным.

Одно из решений заключается в сортировке массива и выборе трех наибольших чисел. 
Однако, есть более эффективное решение, которое не требует полного сортирования массива. 
Вместо этого можно просто найти три наибольших числа и два наименьших числа в массиве, 
и сравнить произведения двух вариантов: произведение трех наибольших чисел 
и произведение одного наибольшего числа и двух наименьших.

Таким образом, задача сводится к поиску трех наибольших и двух наименьших чисел в массиве.


## Example 1:
````
Input: nums = [1,2,3]
Output: 6
````


## Example 2:
````
Input: nums = [1,2,3,4]
Output: 24
````


## Example 3:
````
Input: nums = [-1,-2,-3]
Output: -6
````


## Constraints:
- 3 <= nums.length <= 10^4
- -1000 <= nums[i] <= 1000
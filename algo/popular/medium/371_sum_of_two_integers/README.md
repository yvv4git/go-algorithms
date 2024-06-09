# 371. Sum of Two Integers


## Level - medium


## Task
Given two integers a and b, return the sum of the two integers without using the operators + and -.


## Объяснение
Задача предлагает решить задачу сложения двух целых чисел без использования оператора "+" или "-". 
Это можно сделать с помощью побитовых операций.

Основная идея заключается в использовании побитовых операций XOR (^) и AND (&) для выполнения сложения. 
XOR выполняет сложение без учета переноса, а AND выполняет сдвиг влево для учета переноса.


## Example 1:
````
Input: a = 1, b = 2
Output: 3
````


## Example 2:
````
Input: a = 2, b = 3
Output: 5
````


## Constraints:
- -1000 <= a, b <= 1000
# 1295. Find Numbers with Even Number of Digits


## Level - easy


## Task
Given an array nums of integers, return how many of them contain an even number of digits.


## Объяснение
Дан массив целых чисел nums. Необходимо определить, сколько чисел в этом массиве содержат четное количество цифр.

Пример:
- Вход: nums = [12, 345, 2, 6, 7896]
- Выход: 2

Объяснение:
- 12 содержит 2 цифры (четное количество).
- 345 содержит 3 цифры (нечетное количество).
- 2 содержит 1 цифру (нечетное количество).
- 6 содержит 1 цифру (нечетное количество).
- 7896 содержит 4 цифры (четное количество).

Таким образом, только 12 и 7896 содержат четное количество цифр.


## Example 1:
```
Input: nums = [12,345,2,6,7896]
Output: 2
Explanation: 
12 contains 2 digits (even number of digits). 
345 contains 3 digits (odd number of digits). 
2 contains 1 digit (odd number of digits). 
6 contains 1 digit (odd number of digits). 
7896 contains 4 digits (even number of digits). 
Therefore only 12 and 7896 contain an even number of digits.
```


## Example 2:
```
Input: nums = [555,901,482,1771]
Output: 1 
Explanation: 
Only 1771 contains an even number of digits.
```


## Constraints:
- 1 <= nums.length <= 500
- 1 <= nums[i] <= 10^5
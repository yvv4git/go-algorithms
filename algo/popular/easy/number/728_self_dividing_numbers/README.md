# 728. Self Dividing Numbers


## Level - easy


## Task
A self-dividing number is a number that is divisible by every digit it contains.

- For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.

A self-dividing number is not allowed to contain the digit zero.

Given two integers left and right, return a list of all the self-dividing numbers in the range [left, right].


## Объяснение 
Задача заключается в нахождении всех чисел в заданном диапазоне, которые могут быть разделены на свои цифры и не имеют остатка.

Число называется "саморазделившимся", если оно может быть разделено на каждую из своих цифр. Например, 128 делится на 1, 2, и 8.

Алгоритм должен принимать два целых числа left и right и возвращать список всех саморазделившихся чисел в этом диапазоне, включая границы.

Например, если left = 1 и right = 22, то функция должна вернуть [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22], так как эти числа могут быть разделены на свои цифры и не имеют остатка.


## Example 1:
````
Input: left = 1, right = 22
Output: [1,2,3,4,5,6,7,8,9,11,12,15,22]
````


## Example 2:
````
Input: left = 47, right = 85
Output: [48,55,66,77]
````


## Constraints:
- 1 <= left <= right <= 10^4
# 201. Bitwise AND of Numbers Range


## Level - medium


## Task
Given two integers left and right that represent the range [left, right], return the bitwise AND of all numbers in this range, inclusive.


## Объяснение
Задача предлагает найти побитовое И для всех чисел в заданном диапазоне.

Побитовое И (AND) - это бинарная операция, которая сравнивает каждый бит первого операнда с соответствующим битом второго операнда. 
Если оба бита равны 1, то результатом сравнения будет 1. В противном случае результатом будет 0.

Задача требует найти побитовое И для всех чисел в диапазоне от m до n. Это означает, что нужно найти число, которое, 
когда сравнивается с каждым числом в диапазоне, дает наибольшее число единиц в результате.

Например, если m=5 (в двоичном виде 101) и n=7 (в двоичном виде 111), то побитовое И для всех чисел в этом диапазоне будет 100, 
которое в десятичном виде равно 4.

Решение можно реализовать с помощью побитового сдвига. 
Если m и n различаются, то наибольший общий префикс для m и n будет содержать единицы в старших разрядах, 
а младшие разряды будут различаться. Поэтому нужно сдвинуть m и n вправо до тех пор, пока они не станут равными, 
а затем сдвинуть обратно на то количество разрядов, на которое были сдвинуты.


## Example 1:
````
Input: left = 5, right = 7
Output: 4
````


## Example 2:
````
Input: left = 0, right = 0
Output: 0
````


## Example 3:
````
Input: left = 1, right = 2147483647
Output: 0
````


## Constraints:
- 0 <= left <= right <= 2^31 - 1
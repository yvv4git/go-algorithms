# 189. Rotate Array


## Level - medium


## Task
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.


## Объяснение
Задача предлагает вам повернуть массив на определенное количество шагов. 
Это означает, что последние k элементов массива должны быть перемещены в начало массива, 
а остальные элементы должны быть сдвинуты вправо.

Например, если у вас есть массив [1, 2, 3, 4, 5, 6, 7] и k = 3, то после поворота массив должен стать [5, 6, 7, 1, 2, 3, 4].

Ваша задача - реализовать функцию, которая будет принимать массив и число k, и возвращать повернутый массив.
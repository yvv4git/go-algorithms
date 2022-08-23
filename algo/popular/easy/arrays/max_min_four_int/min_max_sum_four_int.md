# Найти min и max в массиве

## О задаче
https://www.hackerrank.com/challenges/mini-max-sum/problem?isFullScreen=true
Но к сожалению, на этом сайте не дают посмотреть все проверочные кейсы, просят денег.
Так что не известно, верно ли я ее решил.

## Описание
Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. 
Then print the respective minimum and maximum values as a single line of two space-separated long integers.


## Example
Дан массив arr = [1, 3, 5, 7, 9].
Надо найти 4 минимальных элемента и 4 максимальных.


## Решение
Изначально, массив не отсортированный. 
1. Отсортировать массив по возрастанию.
2. Min = сумма первых 4-х элементов.
3. Max = сумма последних 4-х элементов.
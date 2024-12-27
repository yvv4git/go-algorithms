# 1399. Count Largest Group


## Level - easy


## Task
You are given an integer n.

Each number from 1 to n is grouped according to the sum of its digits.

Return the number of groups that have the largest size.


## Объяснение
Дано целое число n. Необходимо сгруппировать все числа от 1 до n по сумме их цифр. 
Затем нужно определить, сколько групп имеют максимальный размер.

Группировка по сумме цифр:
- Для каждого числа от 1 до n вычисляется сумма его цифр.
- Например, для числа 13 сумма цифр равна 1 + 3 = 4.
- Все числа с одинаковой суммой цифр попадают в одну группу.

Определение размера групп:
После группировки нужно найти размер каждой группы.
Например, если n = 13, то группы будут выглядеть так:
- Сумма 1: [1, 10] (размер 2)
- Сумма 2: [2, 11] (размер 2)
- Сумма 3: [3, 12] (размер 2)
- Сумма 4: [4, 13] (размер 2)
- Сумма 5: [5] (размер 1)
и так далее.

Подсчёт групп с максимальным размером:
- Найдите максимальный размер среди всех групп.
- Посчитайте, сколько групп имеют этот максимальный размер.


## Example 1:
```
Input: n = 13
Output: 4
Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9].
There are 4 groups with largest size.
```


## Example 2:
```
Input: n = 2
Output: 2
Explanation: There are 2 groups [1], [2] of size 1.
```


## Constraints:
- 1 <= n <= 10^4
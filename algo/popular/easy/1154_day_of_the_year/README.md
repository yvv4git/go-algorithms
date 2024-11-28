# 1154. Day of the Year


## Level - easy


## Task
Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.


## Объяснение
Дана дата в формате "день.месяц.год". Необходимо определить, какой по счету день года это соответствует.

Пример:  
Если дата — 15.03.2023, то нужно определить, что это 74-й день года.


## Example 1:
```
Input: date = "2019-01-09"
Output: 9
Explanation: Given date is the 9th day of the year in 2019.
```


## Example 2:
```
Input: date = "2019-02-10"
Output: 41
```


## Constraints:
- date.length == 10
- date[4] == date[7] == '-', and all other date[i]'s are digits
- date represents a calendar date between Jan 1st, 1900 and Dec 31th, 2019.
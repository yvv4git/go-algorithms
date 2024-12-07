# 1185. Day of the Week


## Level - easy


## Task
Given a date, return the corresponding day of the week for that date.

The input is given as three integers representing the day, month and year respectively.

Return the answer as one of the following values {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.


## Объяснение
Вам даны три целых числа: d (день), m (месяц) и y (год). Ваша задача — определить, 
какой день недели приходится на эту дату.

Входные данные:
- d — день (1 ≤ d ≤ 31)
- m — месяц (1 ≤ m ≤ 12)
- y — год (1900 ≤ y ≤ 2100)

Выходные данные:
- Название дня недели на английском языке (например, "Monday", "Tuesday", и т.д.).


## Example 1:
```
Input: day = 31, month = 8, year = 2019
Output: "Saturday"
```


## Example 2:
```
Input: day = 18, month = 7, year = 1999
Output: "Sunday"
```


## Example 3:
```
Input: day = 15, month = 8, year = 1993
Output: "Sunday"
```


## Constraints:
- The given dates are valid dates between the years 1971 and 2100.
# 1507. Reformat Date


## Level - easy


## Task
Given a date string in the form Day Month Year, where:
- Day is in the set {"1st", "2nd", "3rd", "4th", ..., "30th", "31st"}.
- Month is in the set {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}.
- Year is in the range [1900, 2100].

Convert the date string to the format YYYY-MM-DD, where:
- YYYY denotes the 4 digit year.
- MM denotes the 2 digit month.
- DD denotes the 2 digit day.


## Объяснение
Дана строка date, представляющая дату в формате Day Month Year, где:
- Day — это день месяца, записанный в виде числа с суффиксом (1st, 2nd, 3rd, 4th, ..., 31st).
- Month — это название месяца на английском языке (Jan, Feb, Mar, ..., Dec).
- Year — это год, записанный четырьмя цифрами (например, 2023).

Необходимо преобразовать эту строку в формат YYYY-MM-DD, где:
- YYYY — год, записанный четырьмя цифрами.
- MM — месяц, записанный двумя цифрами (например, 01 для января, 02 для февраля и т.д.).
- DD — день, записанный двумя цифрами (например, 01, 02, ..., 31).


## Example 1:
```
Input: date = "20th Oct 2052"
Output: "2052-10-20"
```


## Example 2:
```
Input: date = "6th Jun 1933"
Output: "1933-06-06"
```


## Example 3:
```
Input: date = "26th May 1960"
Output: "1960-05-26"
```


## Constraints:
- The given dates are guaranteed to be valid, so no error handling is necessary.
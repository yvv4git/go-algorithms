# 258. Add Digits

## Task
Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

## Объяснение
Дается число, которое состоит из двух знаков и более. Надо разложить его на сумму чисел, затем каждое из них аналогично.
До тех пор, пока не получится одна цифра.ы

## Constraints:
- 0 <= num <= 231 - 1

## Example 1:
````
Input: num = 38
Output: 2
Explanation: The process is
38 --> 3 + 8 --> 11
11 --> 1 + 1 --> 2
Since 2 has only one digit, return it.
````

## Example 2:
````
Input: num = 0
Output: 0
````
# 1556. Thousand Separator


## Level - easy


## Task
Given an integer n, add a dot (".") as the thousands separator and return it in string format.


## Объяснение
Дано целое число n, нужно вернуть его строковое представление с точками в качестве разделителей тысяч. 
Разделитель ставится каждые три цифры справа налево, начиная с конца числа.

Уточнение:
- Число n неотрицательное.
- Разделители добавляются только в целой части, так как n — целое.


## Example 1:
````
Input: n = 987
Output: "987"
````


## Example 2:
````
Input: n = 1234
Output: "1.234"
````

## Constraints:
- 0 <= n <= 231 - 1
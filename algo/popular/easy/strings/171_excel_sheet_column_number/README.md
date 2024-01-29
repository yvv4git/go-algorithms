# 171. Excel Sheet Column Number


## Level - easy


## Task
Given a string columnTitle that represents the column title as appears in an Excel sheet, 
return its corresponding column number.

For example:
````
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28
...
````

## Объяснение
Задача заключается в том, чтобы преобразовать название столбца в Excel таблицы в его порядковый номер. 
Например, столбец A в Excel имеет порядковый номер 1, столбец B - 2, и так далее до Z (26), затем столбцы AA (27), 
AB (28) и так далее.

Таким образом, задача сводится к преобразованию системы счисления с основанием 26 в десятичную систему счисления. 
Каждый символ столбца представляет собой позицию в алфавите, где A - 1, B - 2, ..., Z - 26.

Например, для столбца "ZY", его порядковый номер будет равен 26 * 26 + 25 = 701.


## Example 1:
````
Input: columnTitle = "A"
Output: 1
````


## Example 2:
````
Input: columnTitle = "AB"
Output: 28
````


## Example 3:
````
Input: columnTitle = "ZY"
Output: 701
````


## Constraints:
- 1 <= columnTitle.length <= 7
- columnTitle consists only of uppercase English letters.
- columnTitle is in the range ["A", "FXSHRXW"].
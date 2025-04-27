# 423. Reconstruct Original Digits from English


## Level - medium


## Task
Given a string s containing an out-of-order English representation of digits 0-9, return the digits in ascending order.


## Объяснение
Задача требует восстановить исходный порядок цифр от 0 до 9 на основе строки s, которая содержит неправильный порядок цифр.
Нужно вернуть цифры в порядке возрастания.
Пример:
- Input: s = "owoztneoer"
- Output: "012"


## Example 1:
````
Input: s = "owoztneoer"
Output: "012"
````

## Example 2:
````
Input: s = "fviefuro"
Output: "45"
````

## Constraints:
- 1 <= s.length <= 105
- s[i] is one of the characters ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"].
- s is guaranteed to be valid.
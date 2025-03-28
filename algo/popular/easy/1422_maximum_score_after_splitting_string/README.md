# 1422. Maximum Score After Splitting a String


## Level - easy


## Task
Given a string s of zeros and ones, return the maximum score after splitting the string into two non-empty substrings (i.e. left substring and right substring).

The score after splitting a string is the number of zeros in the left substring plus the number of ones in the right substring.


## Объяснение
Дана строка s, состоящая только из символов '0' и '1'. 
Необходимо разделить строку на две непустые подстроки (левую и правую) таким образом, 
чтобы максимизировать сумму количества '0' в левой части и количества '1' в правой части.

Нужно найти такой индекс i (где 1 <= i < len(s)), чтобы:
- Левая подстрока s[0..i-1] содержала как можно больше '0'.
- Правая подстрока s[i..n-1] содержала как можно больше '1'.
- Сумма количества '0' слева и '1' справа была максимальной.

Пример:

Входная строка: "011101".

- Возможные разделения:
- - Разделение после 1-го символа: "0" | "11101" → '0' слева: 1, '1' справа: 4 → сумма = 1 + 4 = 5.
- - Разделение после 2-го символа: "01" | "1101" → '0' слева: 1, '1' справа: 3 → сумма = 1 + 3 = 4.
- - Разделение после 3-го символа: "011" | "101" → '0' слева: 1, '1' справа: 2 → сумма = 1 + 2 = 3.
- - Разделение после 4-го символа: "0111" | "01" → '0' слева: 1, '1' справа: 1 → сумма = 1 + 1 = 2.
- - Разделение после 5-го символа: "01110" | "1" → '0' слева: 1, '1' справа: 1 → сумма = 1 + 1 = 2.
- Максимальная сумма: 5 (при разделении после первого символа).


## Example 1:
```
Input: s = "011101"
Output: 5 
Explanation: 
All possible ways of splitting s into two non-empty substrings are:
left = "0" and right = "11101", score = 1 + 4 = 5 
left = "01" and right = "1101", score = 1 + 3 = 4 
left = "011" and right = "101", score = 1 + 2 = 3 
left = "0111" and right = "01", score = 1 + 1 = 2 
left = "01110" and right = "1", score = 2 + 1 = 3
```

## Example 2:
```
Input: s = "00111"
Output: 5
Explanation: When left = "00" and right = "111", we get the maximum score = 2 + 3 = 5
```


## Example 3:
```
Input: s = "1111"
Output: 3
```

## Constraints:
- 2 <= s.length <= 500
- The string s consists of characters '0' and '1' only.
# 1071. Greatest Common Divisor of Strings


## Level - easy


## Task
For two strings s and t, we say "t divides s" if and only if s = t + t + t + ... + t + t (i.e., t is concatenated with itself one or more times).

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.


## Объяснение
Даны две строки str1 и str2. Нужно найти наибольшую строку gcd, которая может быть повторена для формирования обеих строк. 
Если такой строки не существует, то возвращается пустая строка.

Пример-1:
- Ввод: str1 = "ABCABC", str2 = "ABC"
- Вывод: "ABC"
- Объяснение: Строка "ABC" повторяется два раза, чтобы получить "ABCABC", и она также является частью "ABC".

Пример-2:
- Ввод: str1 = "ABABAB", str2 = "ABAB"
- Вывод: "AB"
- Объяснение: Строка "AB" повторяется три раза, чтобы получить "ABABAB", и два раза для "ABAB".

Пример-3:
- Ввод: str1 = "LEET", str2 = "CODE"
- Вывод: ""
- Объяснение: Нет общей строки, которая может быть повторена для формирования обеих строк.


## Example 1:
```
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
```

## Example 2:
```
Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
```

## Example 3:
```
Input: str1 = "LEET", str2 = "CODE"
Output: ""
```

## Constraints:
- 1 <= str1.length, str2.length <= 1000
- str1 and str2 consist of English uppercase letters.`
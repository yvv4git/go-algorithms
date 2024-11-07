# 771. Jewels and Stones


## Level - easy


## Task
You're given strings jewels representing the types of stones that are jewels, and stones representing the stones you have. Each character in stones is a type of stone you have. 
You want to know how many of the stones you have are also jewels.

Letters are case sensitive, so "a" is considered a different type of stone from "A".


## Объяснение
У вас есть две строки: jewels (драгоценности) и stones (камни). Строка jewels представляет собой набор символов, каждый из которых является драгоценностью. 
Строка stones представляет собой набор символов, каждый из которых является камнем. Вам нужно определить, сколько камней из строки stones являются драгоценностями.

Пример:
```
Ввод: jewels = "aA", stones = "aAAbbbb"
Вывод: 3
```
Надо написать функцию, которая принимает две строки jewels и stones и возвращает количество символов в строке stones, которые являются драгоценностями.

Один из простых способов решения этой задачи - использовать цикл для перебора каждого символа в строке stones и проверки, 
содержится ли этот символ в строке jewels. Если символ содержится в jewels, то увеличиваем счетчик.


## Example 1:
```
Input: jewels = "aA", stones = "aAAbbbb"
Output: 3
```

## Example 2:
```
Input: jewels = "z", stones = "ZZ"
Output: 0
```


## Constraints:
- 1 <= jewels.length, stones.length <= 50
- jewels and stones consist of only English letters.
- All the characters of jewels are unique.
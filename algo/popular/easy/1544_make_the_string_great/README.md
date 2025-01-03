# 1544. Make The String Great


## Level - easy


## Task
Given a string s of lower and upper case English letters.

A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:
- 0 <= i <= s.length - 2
- s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.

To make the string good, you can choose two adjacent characters that make the string bad and remove them. You can keep doing this until the string becomes good.

Return the string after making it good. The answer is guaranteed to be unique under the given constraints.

Notice that an empty string is also good.


## Объяснение
Задача заключается в том, чтобы преобразовать строку, удаляя соседние символы, 
которые являются одной и той же буквой, но в разных регистрах (например, aA или Bb). 
Цель — сделать строку "хорошей", то есть такой, в которой нет таких "плохих" пар символов.

Дана строка s, состоящая из букв английского алфавита в верхнем и нижнем регистрах. 
Необходимо преобразовать строку так, чтобы в ней не было соседних символов, которые являются одной и той же буквой, 
но в разных регистрах (например, aA, Bb, cC и т.д.). Процесс нужно повторять до тех пор, пока строка не станет "хорошей".

Для решения задачи можно использовать стек (stack). Алгоритм следующий:  
1. Проходим по каждому символу строки.
2. Если стек не пуст и текущий символ образует "плохую" пару с верхним элементом стека (то есть это одна и та же буква, но в разных регистрах), удаляем верхний элемент стека.
3. Иначе добавляем текущий символ в стек.
4. В конце собираем строку из символов, оставшихся в стеке.


## Example 1:
```
Input: s = "leEeetcode"
Output: "leetcode"
Explanation: In the first step, either you choose i = 1 or i = 2, both will result "leEeetcode" to be reduced to "leetcode".
```


## Example 2:
```
Input: s = "abBAcC"
Output: ""
Explanation: We have many possible scenarios, and all lead to the same answer. For example:
"abBAcC" --> "aAcC" --> "cC" --> ""
"abBAcC" --> "abBA" --> "aA" --> ""
```


## Example 3:
```
Input: s = "s"
Output: "s"
```


## Constraints:
- 1 <= s.length <= 100
- s contains only lower and upper case English letters.
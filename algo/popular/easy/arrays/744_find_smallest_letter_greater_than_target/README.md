# 744. Find Smallest Letter Greater Than Target


## Level - easy


## Task
You are given an array of characters letters that is sorted in non-decreasing order, and a character target. 
There are at least two different characters in letters. Return the smallest character in letters that is lexicographically greater than target. 
If such a character does not exist, return the first character in letters.


## Объяснение
Задача предполагает нахождение наименьшей буквы в отсортированном по возрастанию массиве букв, которая строго больше заданной цели. 
Если такой буквы не существует, то возвращается первая буква в массиве.

Например, если у нас есть массив ['c', 'f', 'j'] и наша цель - 'a', то наименьшей буквой, которая строго больше 'a', будет 'c'. 
Если наша цель - 'c', наименьшей буквай, которая строго больше 'c', будет 'f'. 
Если наша цель - 'j' или 'd', такой буквы не существует, поэтому будет возвращена первая буква в массиве - 'c'.

Решение этой задачи может быть реализовано различными способами, но одно из простых решений - это использование бинарного поиска. 
Бинарный поиск позволяет находить элемент в отсортированном массиве за логарифмическое время, что делает его эффективным для этой задачи.


## Example 1:
````
Input: letters = ["c","f","j"], target = "a"
Output: "c"
Explanation: The smallest character that is lexicographically greater than 'a' in letters is 'c'.
````


## Example 2:
````
Input: letters = ["c","f","j"], target = "c"
Output: "f"
Explanation: The smallest character that is lexicographically greater than 'c' in letters is 'f'.
````


## Example 3:
````
Input: letters = ["x","x","y","y"], target = "z"
Output: "x"
Explanation: There are no characters in letters that is lexicographically greater than 'z' so we return letters[0].
````

## Constraints:
- 2 <= letters.length <= 10^4
- letters[i] is a lowercase English letter.
- letters is sorted in non-decreasing order.
- letters contains at least two different characters.
- target is a lowercase English letter.
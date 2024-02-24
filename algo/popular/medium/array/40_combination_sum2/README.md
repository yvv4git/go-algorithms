# 40. Combination Sum II


## Level - medium


## Task
Given a collection of candidate numbers (candidates) and a target number (target), find all unique combinations in candidates where the candidate numbers sum to target.

Each number in candidates may only be used once in the combination.

Note: The solution set must not contain duplicate combinations.


## Объяснение
Задача заключается в нахождении всех уникальных комбинаций чисел из заданного массива, которые дают заданную сумму. 
Каждая число из массива может быть использовано только один раз.

Например, если задан массив [10,1,2,7,6,1,5] и заданная сумма 8, то одна из возможных комбинаций - [1, 7].

Чтобы решить эту задачу, можно использовать рекурсивный подход с отслеживанием суммы и индекса. 
На каждом шаге рекурсии выбирается следующее число из массива и вызывается рекурсивный вызов для оставшейся части массива. 
Если сумма текущей комбинации становится равной заданной сумме, то комбинация добавляется в результат.


## Example 1:
````
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
````


## Example 2:
````
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]
````


## Constraints:
- 1 <= candidates.length <= 100
- 1 <= candidates[i] <= 50
- 1 <= target <= 30
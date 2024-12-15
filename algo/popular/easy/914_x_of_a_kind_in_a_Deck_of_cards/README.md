# 914. X of a Kind in a Deck of Cards


## Level - easy


## Task
You are given an integer array deck where deck[i] represents the number written on the ith card.

Partition the cards into one or more groups such that:
- Each group has exactly x cards where x > 1, and
- All the cards in one group have the same integer written on them.

Return true if such partition is possible, or false otherwise.


## Объяснение
Дана колода карт, где каждая карта имеет два атрибута: значение (например, 2, 3, 4, ..., король, туз) и масть (например, черви, бубны, трефы, пики). 
Задача состоит в том, чтобы определить, можно ли разделить колоду на несколько групп так, чтобы в каждой группе было одинаковое количество карт одного значения.

Пример-1:
- Если колода состоит из карт [1, 2, 3, 4, 4, 3, 2, 1], то можно разделить её на две группы по четыре карты, 
где в каждой группе будут пары одинаковых карт (например, [1, 1], [2, 2], [3, 3], [4, 4]).

- Если колода состоит из карт [1, 1, 2, 2, 2, 2], то можно разделить её на две группы по три карты, 
где в каждой группе будут две карты одного значения и одна карта другого значения.



## Example 1:
```
Input: deck = [1,2,3,4,4,3,2,1]
Output: true
Explanation: Possible partition [1,1],[2,2],[3,3],[4,4].
```

## Example 2:
```
Input: deck = [1,1,1,2,2,2,3,3]
Output: false
Explanation: No possible partition.
```


## Constraints:
- 1 <= deck.length <= 10^4
- 0 <= deck[i] < 10^4
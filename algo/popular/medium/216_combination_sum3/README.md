# 216. Combination Sum III


## Level - medium


## Task
Find all valid combinations of k numbers that sum up to n such that the following conditions are true:
- Only numbers 1 through 9 are used.
- Each number is used at most once.

Return a list of all possible valid combinations. 
The list must not contain the same combination twice, and the combinations may be returned in any order.


## Объяснение
Задача заключается в том, чтобы найти все возможные комбинации из чисел от 1 до 9, которые дают в сумме определенное число n. 
В каждой комбинации должно быть использовано ровно k чисел.

Например, если k = 3 и n = 7, то возможными комбинациями будут [1,2,4]

Решение этой задачи может быть реализовано с помощью рекурсии или обхода в глубину (DFS). 
В этом случае, мы будем использовать рекурсию для генерации всех возможных комбинаций чисел, 
и проверять каждую на соответствие условиям задачи.


## Example 1:
````
Input: k = 3, n = 7
Output: [[1,2,4]]
Explanation:
1 + 2 + 4 = 7
There are no other valid combinations.
````


## Example 2:
````
Input: k = 3, n = 9
Output: [[1,2,6],[1,3,5],[2,3,4]]
Explanation:
1 + 2 + 6 = 9
1 + 3 + 5 = 9
2 + 3 + 4 = 9
There are no other valid combinations.
````

## Example 3:
````
Input: k = 4, n = 1
Output: []
Explanation: There are no valid combinations.
Using 4 different numbers in the range [1,9], 
the smallest sum we can get is 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.
````

## Constraints:
- 2 <= k <= 9
- 1 <= n <= 60
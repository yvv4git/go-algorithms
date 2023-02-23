# Climbing Stairs

## Task 
You are climbing a staircase. It takes n steps to reach the top.  
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?


## Description RU
Вы поднимаетесь по лестнице. Требуется n шагов, чтобы достичь вершины.  
Каждый раз вы можете подняться либо на 1, либо на 2 ступеньки. Сколькими различными способами вы можете подняться на вершину?

Вообще, похоже, что это комбинаторная задача, найти количество перестановок. Как известно, количество перестановок - это factorial.  
Т.е. просто надо найти факториал n.


## Constraints:
- 1 <= n <= 45

## Example 1:
```
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
```


## Example 2:
```
Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
```

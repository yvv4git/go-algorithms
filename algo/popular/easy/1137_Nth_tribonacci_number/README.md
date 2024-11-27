# 1137. N-th Tribonacci Number


## Level - easy


## Task
The Tribonacci sequence Tn is defined as follows: 

T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.

Given n, return the value of Tn.


## Объяснение
Задача заключается в нахождении N-го числа в последовательности Трибоначчи.
Последовательность Трибоначчи начинается с трех заданных чисел:
- T(0) = 0
- T(1) = 1
- T(2) = 1

Каждое последующее число в последовательности является суммой трех предыдущих чисел:  
T(n) = T(n-1) + T(n-2) + T(n-3) для n >= 3


## Примеры:
- T(3) = T(2) + T(1) + T(0) = 1 + 1 + 0 = 2
- T(4) = T(3) + T(2) + T(1) = 2 + 1 + 1 = 4
- T(5) = T(4) + T(3) + T(2) = 4 + 2 + 1 = 7


## Example 1:
```
Input: n = 4
Output: 4
Explanation:
T_3 = 0 + 1 + 1 = 2
T_4 = 1 + 1 + 2 = 4
```


## Example 2:
```
Input: n = 25
Output: 1389537
```


## Constraints:
- 0 <= n <= 37
- The answer is guaranteed to fit within a 32-bit integer, ie. answer <= 2^31 - 1.
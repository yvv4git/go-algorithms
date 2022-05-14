# Reverse int

## Task
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

## Constraints:
```
-231 <= x <= 231 - 1
```

## Example 1:
```
Input: x = 123
Output: 321
```

## Example 2:
```
Input: x = -123
Output: -321
```

## Example 3:
```
Input: x = 120
Output: 21
```


## Алгоритм
Любую задачу можно разбить на sub-задачи.  
В данном случае мне нужно десятичное число равернуть.  
Есть одна хитрость. Разобьем процесс на итерации:  
Step-1: 123 / 10 = 12 и остаток от деления 3.  
Step-2: 12 / 10 = 1 и остаток от деления 2.  
Step-3: 1 / 10 = 0 и остаток от дления 1.  

В итоге, алгоритм очень простой.
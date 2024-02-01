# 241. Different Ways to Add Parentheses


## Level - medium


## Task
Given a string expression of numbers and operators, return all possible results from computing all the different possible ways 
to group numbers and operators. You may return the answer in any order.

The test cases are generated such that the output values fit in a 32-bit integer and the number of different results does not exceed 10^4.


## Объяснение
Задача предлагает нам разбить выражение, состоящее из чисел и операций, 
на два числа, разделенных операцией, и рекурсивно разбивать каждое из этих чисел до тех пор, пока не останется одно число. 
Затем мы должны выполнить операцию между этими двумя числами и вернуть результат.

Например, если мы имеем выражение "2-1-1", мы можем разделить его на "2" и "-1-1", а затем рекурсивно разделить "-1-1" на "-1" и "-1". 
После этого мы можем выполнить операцию "-" между "2" и результатом рекурсивного вызова для "-1-1", что даст нам "2".

Задача требует написать функцию, которая будет принимать строку, представляющую выражение, и возвращать все возможные результаты, 
которые можно получить, разбивая выражение и вычисляя результаты.

## Example 1:
````
Input: expression = "2-1-1"
Output: [0,2]
Explanation:
((2-1)-1) = 0
(2-(1-1)) = 2
````


## Example 2:
````
Input: expression = "2*3-4*5"
Output: [-34,-14,-10,-10,10]
Explanation:
(2*(3-(4*5))) = -34
((2*3)-(4*5)) = -14
((2*(3-4))*5) = -10
(2*((3-4)*5)) = -10
(((2*3)-4)*5) = 10
````

## Constraints:
- 1 <= expression.length <= 20
- expression consists of digits and the operator '+', '-', and '*'.
- All the integer values in the input expression are in the range [0, 99].

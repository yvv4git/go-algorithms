# Weird Algorithm


## Task
Consider an algorithm that takes as input a positive integer n. 
If n is even, the algorithm divides it by two, and if n is odd, the algorithm multiplies it by three and adds one. 
The algorithm repeats this, until n is one. For example, the sequence for n=3 is as follows:

3 -> 10 -> 5 -> 16 -> 8 -> 4 -> 2 -> 1

Your task is to simulate the execution of the algorithm for a given value of n.


## Объяснение
Задача "Weird Algorithm" предполагает, что у нас есть некоторое целое число n.
Наша задача - применить к этому числу следующие правила:
- Если n - четное, разделить его на 2.
- Если n - нечетное, умножить его на 3 и прибавить 1.

Наша задача - найти последовательность чисел, которая применяет эти правила к n, пока n не станет 1.

## Input
The only input line contains an integer n.


## Output
Print a line that contains all values of n during the algorithm.


## Constraints
- 1 <= n <= 10^6



## Example
````
Input:
3
Output:
3 10 5 16 8 4 2 1
````
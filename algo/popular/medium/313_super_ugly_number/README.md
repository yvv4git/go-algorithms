# 313. Super Ugly Number


## Level - medium


## Task
A super ugly number is a positive integer whose prime factors are in the array primes.

Given an integer n and an array of integers primes, return the nth super ugly number.

The nth super ugly number is guaranteed to fit in a 32-bit signed integer.


## Объяснение
Задача требует от нас найти n-ое "супер уродливое число". Супер уродливое число определяется как положительное целое число, 
которое имеет простые делители, ограниченные заданным набором простых чисел.

Супер уродливое число — это число, которое может быть представлено в виде p_1^a1 * p_2^a2 * p_k^ak.
Где p_i - это простые числа из заданного набора, и неотрицательные целые числа.

Например, если задан набор простых чисел [2, 3, 5], то первые несколько супер уродливых чисел будут: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12, ...

Задача состоит в том, чтобы найти n-ое супер уродливое число для заданного набора простых чисел.


## Example 1:
````
Input: n = 12, primes = [2,7,13,19]
Output: 32
Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12 super ugly numbers given primes = [2,7,13,19].
````


## Example 2:
````
Input: n = 1, primes = [2,3,5]
Output: 1
Explanation: 1 has no prime factors, therefore all of its prime factors are in the array primes = [2,3,5].
````


## Constraints:
- 1 <= n <= 10^5
- 1 <= primes.length <= 100
- 2 <= primes[i] <= 1000
- primes[i] is guaranteed to be a prime number.
- All the values of primes are unique and sorted in ascending order.
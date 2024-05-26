# 717. 1-bit and 2-bit Characters


## Level - easy


## Task
We have two special characters:
- The first character can be represented by one bit 0.
- The second character can be represented by two bits (10 or 11).

Given a binary array bits that ends with 0, return true if the last character must be a one-bit character.


## Объяснение
Это задача на проверку корректности последовательности битов, представленных в виде массива целых чисел. 
Каждое целое число в массиве может быть либо 0, либо 1, либо 2. 
Если число равно 0, то оно представляет один бит, если число равно 1, то оно представляет два бита. 
Если число равно 2, то оно представляет два бита, но оно используется для обозначения символа, 
который занимает два бита, но не может быть последним в последовательности.

Задача состоит в том, чтобы определить, является ли последний символ в последовательности корректным, 
то есть не может ли он быть представлен только двумя битами.

Например, если у нас есть последовательность [1, 1, 1, 0], то последний символ корректен, 
потому что он может быть представлен одним битом. Если же у нас есть последовательность [1, 1, 2, 0], 
то последний символ некорректен, потому что он может быть представлен двумя битами.

Решение задачи должно быть эффективным по времени и памяти, поскольку оно может применяться к большим массивам данных.


## Example 1:
````
Input: bits = [1,0,0]
Output: true
Explanation: The only way to decode it is two-bit character and one-bit character.
So the last character is one-bit character.
````


## Example 2:
````
Input: bits = [1,1,1,0]
Output: false
Explanation: The only way to decode it is two-bit character and two-bit character.
So the last character is not one-bit character.
````


## Constraints:
- 1 <= bits.length <= 1000
- bits[i] is either 0 or 1.
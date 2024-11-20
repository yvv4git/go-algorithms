# 942. DI String Match


## Level - easy


## Task
A permutation perm of n + 1 integers of all the integers in the range [0, n] can be represented as a string s of length n where:
- s[i] == 'I' if perm[i] < perm[i + 1], and
- s[i] == 'D' if perm[i] > perm[i + 1].

Given a string s, reconstruct the permutation perm and return it. If there are multiple valid permutations perm, return any of them.


## Объяснение
Задача требует от нас создать список целых чисел, который соответствует заданной строке, состоящей из символов 'D' и 'I'.
Дана строка S, состоящая только из символов 'D' (уменьшение) и 'I' (увеличение). 
Нам нужно вернуть список целых чисел A длины N+1, где N — длина строки S, такой, что:
- Если S[i] == 'I', то A[i] < A[i+1]
- Если S[i] == 'D', то A[i] > A[i+1]


## Example 1:
```
Input: s = "IDID"
Output: [0,4,1,3,2]
```


## Example 2:
```
Input: s = "III"
Output: [0,1,2,3]
```


## Example 3:
```
Input: s = "DDI"
Output: [3,2,0,1]
```


## Constraints:
- 1 <= s.length <= 10^5
- s[i] is either 'I' or 'D'.
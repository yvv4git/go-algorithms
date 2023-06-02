# 693. Binary Number with Alternating Bits


## Level - easy


## Task
Given a positive integer, check whether it has alternating bits: namely, 
if two adjacent bits will always have different values.


## Суть задачи
Учитывая положительное целое число, проверьте, имеет ли оно чередующиеся биты, 
а именно, всегда ли два соседних бита будут иметь разные значения.
Например, в числе 7 биты такие 101, т.е. чередуются - ставим true.
Напротив, в числе 11 биты 111, т.е. не чередуются - ставим false.



## Constraints:
- 1 <= n <= 231 - 1


## Example 1:
````
Input: n = 5
Output: true
Explanation: The binary representation of 5 is: 101
````


## Example 2:
````
Input: n = 7
Output: false
Explanation: The binary representation of 7 is: 111.
````


## Example 3:
````
Input: n = 11
Output: false
Explanation: The binary representation of 11 is: 1011.
````
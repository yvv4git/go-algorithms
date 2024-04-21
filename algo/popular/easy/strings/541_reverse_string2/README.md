# 541. Reverse String II


## Level - medium


## Task
Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.

If there are fewer than k characters left, reverse all of them. 
If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and leave the other as original.


## Объяснение
Данная задача состоит в том, чтобы перевернуть каждый отдельный блок из k символов в строке s. 
Если количество символов в строке s меньше k, то просто переверните все символы в строке. 
Если количество символов в строке s больше k, то переверните первые k символов и продолжайте переворачивать каждый следующий блок из k символов. 
Необходимо, чтобы каждый блок начинался с позиции, кратной k.

Например, если вы вводите строку s = "abcdefg" и k = 2, то вы должны получить строку "bacdfeg".

Обратите внимание, что только первый символ каждого блока из k символов изменяется. Остальные символы остаются неизменными.


## Example 1:
````
Input: s = "abcdefg", k = 2
Output: "bacdfeg"
````


## Example 2:
````
Input: s = "abcd", k = 2
Output: "bacd"
````


## Constraints:
- 1 <= s.length <= 10^4
- s consists of only lowercase English letters.
- 1 <= k <= 10^4
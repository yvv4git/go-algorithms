# 93. Restore IP Addresses


## Level - medium


## Task
A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.

For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.

Given a string s containing only digits, return all possible valid IP addresses that can be formed by inserting dots into s. 
You are not allowed to reorder or remove any digits in s. You may return the valid IP addresses in any order.


## Объяснение
Задача заключается в том, чтобы разбить строку, представляющую целое число, на четыре части таким образом, 
чтобы каждая часть представляла допустимый IP-адрес. 
Допустимый IP-адрес состоит из четырех чисел, каждое из которых находится в диапазоне от 0 до 255 и не начинается с 0, 
если оно не является единственным числом в этой части.

Пример:
- Вход: s = "25525511135"
- Выход: ["255.255.11.135", "255.255.111.35"]

Цель задачи - написать функцию, которая принимает на вход строку, представляющую целое число, 
и возвращает все возможные комбинации IP-адресов, которые можно получить из этой строки.


## Example 1:
````
Input: s = "25525511135"
Output: ["255.255.11.135","255.255.111.35"]
````


## Example 2:
````
Input: s = "0000"
Output: ["0.0.0.0"]
````


## Example 3:
````
Input: s = "101023"
Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
````


## Constraints:
- 1 <= s.length <= 20
- s consists of digits only.
# 1108. Defanging an IP Address


## Level - easy


## Task
Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".


## Объяснение
Задача заключается в том, чтобы преобразовать IP-адрес (например, "192.168.1.1") в "defanged" формат. 
"Defanged" IP-адрес — это такой формат, в котором точки (.) заменяются на "[.]".

Например:
- Исходный IP-адрес: "192.168.1.1"
- Преобразованный IP-адрес: "192[.]168[.]1[.]1"


## Example 1:
```
Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"
```


## Example 2:
```
Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0"
```


## Constraints:
- The given address is a valid IPv4 address.
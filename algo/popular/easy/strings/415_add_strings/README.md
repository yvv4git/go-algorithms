# 415. Add Strings


## Level - medium


## Task
Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.

You must solve the problem without using any built-in library for handling large integers (such as BigInteger). 
You must also not convert the inputs to integers directly.


## Объяснение
Задача предлагает нам сложить два числа, представленных в виде строк. 
Это означает, что мы должны сложить два числа без использования каких-либо встроенных библиотек, 
которые бы преобразовывали эти строки в числа для сложения. 
Вместо этого, нам нужно реализовать алгоритм сложения вручную.

Основные шаги, которые нам нужно выполнить для решения этой задачи, включают:
1. Инициализация указателя для каждой строки с конца.
2. Инициализация переменной для хранения переноса (carry).
3. Создание пустой строки для результата.
4. Пока мы не достигли начала каждой строки, выполняйте следующие шаги:
-- Если указатель еще не дошел до начала строки, извлеките цифру из строки и преобразуйте ее в целое число. В противном случае считайте ее как 0.
-- Вычислите сумму цифр из обеих строк и переноса.
-- Добавьте последнюю цифру суммы в начало результата.
-- Обновите перенос для следующей итерации.
5. Если после завершения цикла перенос не равен 0, добавьте его в начало результата.
6. Возвращаем результат.



## Example 1:
````
Input: num1 = "11", num2 = "123"
Output: "134"
````

## Example 2:
````
Input: num1 = "456", num2 = "77"
Output: "533"
````


## Example 3:
````
Input: num1 = "0", num2 = "0"
Output: "0"
````


## Constraints:
- 1 <= num1.length, num2.length <= 10^4
- num1 and num2 consist of only digits.
- num1 and num2 don't have any leading zeros except for the zero itself.
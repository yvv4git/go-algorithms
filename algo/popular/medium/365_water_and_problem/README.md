# 365. Water and Jug Problem


## Level - medium


## Task
You are given two jugs with capacities x liters and y liters. You have an infinite water supply. Return whether the total amount of water in both jugs may reach target using the following operations:
- Fill either jug completely with water.
- Completely empty either jug.
- Pour water from one jug into another until the receiving jug is full, or the transferring jug is empty.


## Объяснение
Задача, также известная как "Water and Jug Problem" (Задача с ведрами и водой), является классической задачей в теории чисел 
и теории алгоритмов.
У вас есть два ведра с известными объемами x и y литров. Вы можете выполнять следующие операции с этими ведрами:
1. Наполнить ведро до краев.
2. Опустошить ведро.
3. Перелить воду из одного ведра в другое до тех пор, пока одно из ведер не станет полным или пустым.

Даны два целых числа x и y, представляющие объемы двух ведер, и целое число z, представляющее желаемый объем воды. 
Нужно определить, можно ли получить z литров воды, используя только эти два ведра и указанные операции.

Решение этой задачи основывается на теореме Безу и нахождении наибольшего общего делителя (НОД) чисел x и y. 
Если z является кратным НОД(x, y) и z не превышает суммы x и y, то можно получить z литров воды. 
В противном случае это невозможно.

Таким образом, ключевым моментом в решении задачи является проверка следующих условий:
1. z должно быть кратно НОД(x, y).
2. z должно быть меньше или равно x + y.

Если оба условия выполняются, то ответ положительный, иначе — отрицательный.


## Example 1:
````
Input: x = 3, y = 5, target = 4

Output: true

Explanation:

Follow these steps to reach a total of 4 liters:

Fill the 5-liter jug (0, 5).
Pour from the 5-liter jug into the 3-liter jug, leaving 2 liters (3, 2).
Empty the 3-liter jug (0, 2).
Transfer the 2 liters from the 5-liter jug to the 3-liter jug (2, 0).
Fill the 5-liter jug again (2, 5).
Pour from the 5-liter jug into the 3-liter jug until the 3-liter jug is full. This leaves 4 liters in the 5-liter jug (3, 4).
Empty the 3-liter jug. Now, you have exactly 4 liters in the 5-liter jug (0, 4).
Reference: The Die Hard example.
````


## Example 2:
````
Input: x = 2, y = 6, target = 5

Output: false
````


## Example 3:
````
Input: x = 1, y = 2, target = 3

Output: true

Explanation: Fill both jugs. The total amount of water in both jugs is equal to 3 now.
````


## Constraints:

- 1 <= x, y, target <= 10^3
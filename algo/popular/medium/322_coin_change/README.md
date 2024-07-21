# 322. Coin Change


## Level - medium


## Task
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount. 
If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.


## Объясение
Задача заключается в том, чтобы найти минимальное количество монет, которое нужно использовать, чтобы набрать заданную сумму. 
Если это невозможно, то нужно вернуть -1.

Формальное описание задачи:
- У вас есть массив различных номиналов монет coins, где coins[i] представляет собой номинал i-й монеты.
- Вам нужно найти минимальное количество монет, которое нужно использовать, чтобы набрать сумму amount. 
- Если это невозможно, вернуть -1.

Примеры:
````
Вход: coins = [1, 2, 5], amount = 11
Выход: 3
Объяснение: 11 = 5 + 5 + 1
````

````
Вход: coins = [2], amount = 3
Выход: -1
Объяснение: Невозможно набрать сумму 3, используя только монеты номиналом 2.
````

````
Вход: coins = [1], amount = 0
Выход: 0
Объяснение: Для суммы 0 не нужно использовать ни одной монеты.
````

Для решения этой задачи можно использовать динамическое программирование. 
Основная идея заключается в том, чтобы создать массив dp, где dp[i] будет представлять минимальное количество монет, 
необходимое для набора суммы i. Мы заполняем этот массив, используя значения из предыдущих подзадач.


## Объяснение


## Example 1:
````
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
````


## Example 2:
````
Input: coins = [2], amount = 3
Output: -1
````


## Example 3:
````
Input: coins = [1], amount = 0
Output: 0
````


## Constraints:
- 1 <= coins.length <= 12
- 1 <= coins[i] <= 2^31 - 1
- 0 <= amount <= 10^4
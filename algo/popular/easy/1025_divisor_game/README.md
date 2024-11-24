# 1025. Divisor Game


## Level - easy


## Task
Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number n on the chalkboard. On each player's turn, that player makes a move consisting of:
- Choosing any x with 0 < x < n and n % x == 0.
- Replacing the number n on the chalkboard with n - x.

Also, if a player cannot make a move, they lose the game.

Return true if and only if Alice wins the game, assuming both players play optimally.


## Объяснение
Два игрока, Алиса и Боб, играют в игру. У них есть число N. Правила игры следующие:
1. Игроки ходят по очереди.
2. На своём ходу игрок должен выбрать любой делитель x числа N (кроме самого N), такой что 0 < x < N.
3. После выбора делителя x, число N уменьшается до N - x.
4. Проигрывает тот, кто не может сделать ход (т.е. когда N становится равным 1).

Задача: Определить, выиграет ли Алиса, если оба игрока играют оптимально.

Если N — чётное число, Алиса может всегда выиграть, выбирая x = 1 на каждом своём ходу, 
что приведёт к тому, что Боб будет получать нечётное число. В конце концов, Боб получит N = 1 и проиграет.

Если N — нечётное число, Алиса не сможет выиграть, так как любой её ход приведёт к тому, 
что Боб получит чётное число, и он сможет использовать ту же стратегию, что и в случае с чётным числом.

## Example 1:
```
Input: n = 2
Output: true
Explanation: Alice chooses 1, and Bob has no more moves.
```


## Example 2:
```
Input: n = 3
Output: false
Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.
```


## Constraints:
- 1 <= n <= 1000
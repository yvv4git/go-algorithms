# 518. Coin Change II

## Task

You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the number of combinations that make up that amount. If that amount of money cannot be made up by any combination of the coins, return 0.

You may assume that you have an infinite number of each kind of coin.

The final answer is guaranteed to fit into a signed 32-bit integer.

## Объяснение

Задача требует подсчитать количество комбинаций (порядок не важен), которыми можно составить заданную сумму amount, используя монеты из coins неограниченное количество раз. Это классическая задача **Unbounded Knapsack (рюкзак с неограниченным числом предметов)**.

### Постановка задачи

1. Дан массив номиналов монет coins и целевая сумма amount
1. Каждую монету можно использовать неограниченное количество раз
1. Нужно вернуть количество **комбинаций** (не перестановок) — порядок монет не имеет значения
1. Если сумму нельзя составить — вернуть 0

### Ключевые аспекты

1. **Unbounded**: каждую монету можно брать сколько угодно раз
1. **Комбинации vs Перестановки**: комбинации не учитывают порядок (1+2 и 2+1 — это одно и то же)
1. **Размерность**: amount ≤ 5000, количество монет ≤ 300 — требуется O(N×M) или O(M) по времени

### Основная идея DP (Dynamic Programming)

Используется массив `dp[0..amount]`, где `dp[a]` — количество способов составить сумму `a`.

**Базовый случай**: `dp[0] = 1` (одна комбинация для нулевой суммы — не взять ничего).

**Переход**: Для каждой монеты `coin` проходим по суммам от `coin` до `amount`:

```text
dp[a] += dp[a - coin]
```

**Важный нюанс — порядок циклов**:

- Если сначала итерировать по **монетам**, затем по **суммам** — считаются **комбинации** (каждый набор монет учитывается 1 раз)
- Если наоборот (сначала суммы, потом монеты) — считаются **перестановки** (разные порядки одних и тех же монет считаются разными способами)

### Пример работы для amount = 5, coins = \[1, 2, 5\]

Начальное состояние: `dp = [1, 0, 0, 0, 0, 0]`

После монеты 1: `dp = [1, 1, 1, 1, 1, 1]`
После монеты 2: `dp = [1, 1, 2, 2, 3, 3]`
После монеты 5: `dp = [1, 1, 2, 2, 3, 4]`

Итоговый ответ: `dp[5] = 4`

### Сложность

- **Time**: O(N × M), где N = количество монет, M = amount
- **Space**: O(M) — одномерный массив dp

## Example 1

Input: amount = 5, coins = [1,2,5]
Output: 4
Explanation: there are four ways to make up the amount:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

## Example 2

Input: amount = 3, coins = [2]
Output: 0
Explanation: the amount of 3 cannot be made up just with coins of 2.

## Example 3

Input: amount = 10, coins = [10]
Output: 1

## Constraints

1 \<= coins.length \<= 300
1 \<= coins[i] \<= 5000
All the values of coins are unique.
0 \<= amount \<= 5000

# 740. Delete and Earn


## Level - medium


## Task
You are given an integer array nums. You want to maximize the number of points you get by performing the following operation any number of times:
- Pick any nums[i] and delete it to earn nums[i] points. Afterwards, you must delete every element equal to nums[i] - 1 and every element equal to nums[i] + 1.

Return the maximum number of points you can earn by applying the above operation some number of times.


## Объяснение
Вам дан массив целых чисел nums. 
Вы хотите максимизировать количество очков, выполняя следующую операцию любое количество раз: выберите любое nums[i], 
удалите его, чтобы заработать nums[i] очков, и затем удалите все элементы, равные nums[i] - 1 и nums[i] + 1.

Эта задача может быть решена с использованием динамического программирования.
Сначала мы группируем числа по их значениям и вычисляем сумму для каждого уникального числа (поскольку удаление одного числа удаляет все его копии).
Затем мы сортируем уникальные числа и применяем подход, аналогичный задаче "House Robber", где мы не можем выбрать два числа, отличающиеся на 1.

Мы создаем массив сумм для каждого числа от минимального до максимального. 
Затем используем DP, где dp[i] - максимальная сумма, которую можно получить, рассматривая числа до i-го.

- Если мы не берем текущее число, dp[i] = dp[i-1]
- Если берем, dp[i] = dp[i-2] + sum[i] (если i >= 2)

Итоговый ответ - dp[len-1]


## Example 1:
```
Input: nums = [3,4,2]
Output: 6
Explanation: You can perform the following operations:
- Delete 4 to earn 4 points. Consequently, 3 is also deleted. nums = [2].
- Delete 2 to earn 2 points. nums = [].
You earn a total of 6 points.
```


## Example 2:
```
Input: nums = [2,2,3,3,3,4]
Output: 9
Explanation: You can perform the following operations:
- Delete a 3 to earn 3 points. All 2's and 4's are also deleted. nums = [3,3].
- Delete a 3 again to earn 3 points. nums = [3].
- Delete a 3 once more to earn 3 points. nums = [].
You earn a total of 9 points.
```


## Constraints:
- 1 <= nums.length <= 2 * 10^4
- 1 <= nums[i] <= 10^4
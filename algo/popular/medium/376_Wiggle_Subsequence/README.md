# 376. Wiggle Subsequence

## TASK

A wiggle sequence is a sequence where the differences between successive numbers strictly alternate between positive and negative. The first difference (if one exists) may be either positive or negative. A sequence with one element and a sequence with two non-equal elements are trivially wiggle sequences.

- For example, [1, 7, 4, 9, 2, 5] is a wiggle sequence because the differences (6, -3, 5, -7, 3) alternate between positive and negative.
- In contrast, [1, 4, 7, 2, 5] and [1, 7, 4, 5, 5] are not wiggle sequences. The first is not because its first two differences are positive, and the second is not because its last difference is zero.

A subsequence is obtained by deleting some elements (possibly zero) from the original sequence, leaving the remaining elements in their original order.

Given an integer array nums, return the length of the longest wiggle subsequence of nums.

## Объяснение

Задача "Wiggle Subsequence" — найти длину самой длинной подпоследовательности, в которой разности между соседними элементами строго чередуются по знаку (положительная, отрицательная, положительная, ...).

Последовательность из одного элемента или из двух неравных элементов уже является wiggle-последовательностью.

### Подходы к решению

1. **Жадный алгоритм (Greedy)** — оптимальное решение за O(n) времени и O(1) памяти.
Идея в подсчёте количества "пиков" и "впадин": проходим по массиву и фиксируем каждый момент, когда направление изменения меняется (рост сменяется падением или наоборот).
Первый элемент всегда считается частью последовательности.

2. **Динамическое программирование (DP)** — используем два массива `up[i]` и `down[i]`:
   - `up[i]` — длина longest wiggle subsequence, заканчивающейся на i-м элементе с возрастающим последним шагом.
   - `down[i]` — то же, но с убывающим последним шагом.
   - Переходы: если `nums[i] > nums[j]`, то `up[i] = max(up[i], down[j] + 1)`; если `nums[i] < nums[j]`, то `down[i] = max(down[i], up[j] + 1)`.
   - Сложность: O(n²) времени, O(n) памяти.

3. **Оптимизированное DP** — модификация, где вместо массивов используются две переменные `up` и `down`, обновляемые за один проход:
   - Если `nums[i] > nums[i-1]`: `up = down + 1`
   - Если `nums[i] < nums[i-1]`: `down = up + 1`
   - Если `nums[i] == nums[i-1]`: ничего не меняется.
   - Сложность: O(n) времени, O(1) памяти.

## Example 1

```text
Input: nums = [1,7,4,9,2,5]
Output: 6
Explanation: The entire sequence is a wiggle sequence with differences (6, -3, 5, -7, 3).
```

## Example 2

```text
Input: nums = [1,17,5,10,13,15,10,5,16,8]
Output: 7
Explanation: There are several subsequences that achieve this length.
One is [1, 17, 10, 13, 10, 16, 8] with differences (16, -7, 3, -3, 6, -8).
```

## Example 3

```text
Input: nums = [1,2,3,4,5,6,7,8,9]
Output: 2
```

## Constraints

- 1 \<= nums.length \<= 1000
- 0 \<= nums[i] \<= 1000

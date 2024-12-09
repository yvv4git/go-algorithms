# 1013. Partition Array Into Three Parts With Equal Sum


## Level - easy


## Task
Given an array of integers arr, return true if we can partition the array into three non-empty parts with equal sums.

Formally, we can partition the array if we can find indexes i + 1 < j with (arr[0] + arr[1] + ... + arr[i] == arr[i + 1] + arr[i + 2] + ... + arr[j - 1] == arr[j] + arr[j + 1] + ... + arr[arr.length - 1])


## Объяснение
Дан массив целых чисел arr. Необходимо определить, 
можно ли разделить этот массив на три непересекающиеся подмассива (части) так, 
чтобы сумма элементов в каждой из этих частей была одинаковой.

Нужно проверить, существуют ли три индекса i, j, и k (где 0 <= i < j < k <= n-1) такие, что:
- Сумма элементов от начала массива до индекса i (включительно) равна sum1.
- Сумма элементов от индекса i+1 до индекса j (включительно) равна sum2.
- Сумма элементов от индекса j+1 до конца массива равна sum3.
- sum1 == sum2 == sum3.

Пример:
Рассмотрим массив arr = [0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1].
- Общая сумма: total_sum = 9.
- target_sum = 9 / 3 = 3.
- Первая часть: [0, 2, 1] (сумма 3).
- Вторая часть: [-6, 6, -7, 9, 1] (сумма 3).
- Третья часть: [2, 0, 1] (сумма 3).

Таким образом, массив можно разделить на три части с равными суммами.


## Example 1:
```
Input: arr = [0,2,1,-6,6,-7,9,1,2,0,1]
Output: true
Explanation: 0 + 2 + 1 = -6 + 6 - 7 + 9 + 1 = 2 + 0 + 1
```


## Example 2:
```
Input: arr = [0,2,1,-6,6,7,9,-1,2,0,1]
Output: false
```


## Example 3:
```
Input: arr = [3,3,6,5,-2,2,5,1,-9,4]
Output: true
Explanation: 3 + 3 = 6 = 5 - 2 + 2 + 5 + 1 - 9 + 4
```


## Constraints:
- 3 <= arr.length <= 5 * 10^4
- -10^4 <= arr[i] <= 10^4
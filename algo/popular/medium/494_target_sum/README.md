# 494. Target Sum


## Level - medium


## Task
You are given an integer array nums and an integer target.

You want to build an expression out of nums by adding one of the symbols '+' and '-' before each integer in nums and then concatenate all the integers.

- For example, if nums = [2, 1], you can add a '+' before 2 and a '-' before 1 and concatenate them to build the expression "+2-1".

Return the number of different expressions that you can build, which evaluates to target.


## Объяснение
У вас есть массив целых чисел и целевое значение. 
Вам нужно найти количество способов, которыми можно присвоить каждому элементу массива знак "плюс" или "минус" так, 
чтобы сумма всех элементов с учетом этих знаков равнялась целевому значению.

Формально, если у вас есть массив nums и целевое значение S, вам нужно найти количество различных комбинаций знаков, 
которые удовлетворяют уравнению, где n — длина массива nums:   
$ \sum_{i=0}^{n-1} \pm \text{nums}[i] = S $.

Пример:
````
Вход: nums = [1, 1, 1, 1, 1], S = 3
Выход: 5
Объяснение: Есть 5 способов присвоить знаки элементам массива, чтобы их сумма равнялась 3:
+1 +1 +1 -1 +1
+1 +1 -1 +1 +1
+1 -1 +1 +1 +1
-1 +1 +1 +1 +1
+1 +1 +1 +1 -1
````

Эту задачу можно решить с помощью динамического программирования или рекурсии с мемоизацией. 
Основная идея заключается в том, чтобы рассматривать задачу как подзадачи и использовать таблицу для хранения уже вычисленных результатов, 
чтобы избежать повторных вычислений.

## Example 1:
````
Input: nums = [1,1,1,1,1], target = 3
Output: 5
Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
````

## Example 2:
````
Input: nums = [1], target = 1
Output: 1
````


## Constraints:
- 1 <= nums.length <= 20
- 0 <= nums[i] <= 1000
- 0 <= sum(nums[i]) <= 1000
- -1000 <= target <= 1000
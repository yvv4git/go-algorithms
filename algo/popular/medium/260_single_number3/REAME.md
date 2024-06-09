# 260. Single Number III


## Level - medium


## Task
Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice. 
Find the two elements that appear only once. You can return the answer in any order.

You must write an algorithm that runs in linear runtime complexity and uses only constant extra space.


## Объяснение
Задача изучает массив целых чисел, в котором каждое число встречается дважды, за исключением двух чисел, 
которые встречаются только один раз. Задача состоит в том, чтобы найти эти два числа.

Например, если входные данные будут [1, 2, 1, 3, 2, 5], то выходные данные должны быть [3, 5].

Одним из возможных решений для этой задачи является использование хэш-таблицы для подсчета количества вхождений каждого числа. 
Затем пройти по хэш-таблице и вернуть числа, которые встречаются только один раз.

Однако, есть более эффективный алгоритм, который использует битовые операции для решения этой задачи. 
Алгоритм основан на том, что если мы сложим все числа в массиве, то каждое число, которое встречается дважды, 
будет умножено на 2, а числа, встречающиеся один раз, не будут влиять на сумму. 
Разница между суммой всех чисел и суммой уникальных чисел будет равна XOR двух одиноковых чисел.


## Example 1:
````
Input: nums = [1,2,1,3,2,5]
Output: [3,5]
Explanation:  [5, 3] is also a valid answer.
````


## Example 2:
````
Input: nums = [-1,0]
Output: [-1,0]
````


## Example 3:
````
Input: nums = [0,1]
Output: [1,0]
````


## Constraints:
- 2 <= nums.length <= 3 * 10^4
- -2^31 <= nums[i] <= 2^31 - 1
- Each integer in nums will appear twice, only two integers will appear once.
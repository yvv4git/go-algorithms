# 1005. Maximize Sum Of Array After K Negations


# Level - easy


## Task
Given an integer array nums and an integer k, modify the array in the following way:
- choose an index i and replace nums[i] with -nums[i].

You should apply this process exactly k times. You may choose the same index i multiple times.

Return the largest possible sum of the array after modifying it in this way.


## Объяснение
Задача "1005. Maximize Sum Of Array After K Negations" заключается в том, 
чтобы найти максимальную возможную сумму элементов массива после выполнения не более K операций изменения знака (переворачивания числа в противоположное).

Каждая операция изменения знака может быть выполнена только один раз и может быть применена только к одному элементу массива за раз. 
Элементы массива могут быть как положительными, так и отрицательными.

Цель задачи - найти максимальную сумму, которую можно получить, выполнив не более K операций изменения знака.

Например, если у нас есть массив [4, 2, 3] и K = 1, то мы можем изменить знак наименьшего по модулю числа, 
чтобы получить максимальную сумму. В этом случае, мы можем изменить знак наименьшего по модулю числа 2, получив массив [-4, 2, 3], и сумма элементов будет 5.

Задача требует написать функцию, которая будет принимать массив чисел и число K, и возвращать максимальную возможную сумму.


## Example 1:
````
Input: nums = [4,2,3], k = 1
Output: 5
Explanation: Choose index 1 and nums becomes [4,-2,3].
````


## Example 2:
````
Input: nums = [3,-1,0,2], k = 3
Output: 6
Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].
````


## Example 3:
````
Input: nums = [2,-3,-1,5,-4], k = 2
Output: 13
Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].
````


## Constraints:
- 1 <= nums.length <= 10^4
- -100 <= nums[i] <= 100
- 1 <= k <= 10^4
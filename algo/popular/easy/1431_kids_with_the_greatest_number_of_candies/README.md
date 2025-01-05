# 1431. Kids With the Greatest Number of Candies


## Level - easy


## Task
There are n kids with candies. You are given an integer array candies, 
where each candies[i] represents the number of candies the ith kid has, 
and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, 
they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.


## Task
У вас есть массив candies, где каждый элемент candies[i] представляет количество конфет у i-го ребенка. 
Также дано целое число extraCandies, которое обозначает количество дополнительных конфет, 
которые можно выдать любому ребенку.

Задача: для каждого ребенка определить, будет ли у него наибольшее количество конфет, если ему выдать все дополнительные конфеты (extraCandies). 
Вернуть массив булевых значений (true или false), где true означает, что у ребенка будет наибольшее количество конфет, а false — что нет.

Пример:  
candies = [2, 3, 5, 1, 3], extraCandies = 3

Объяснение:  
- Если выдать 3 дополнительные конфеты каждому ребенку:
- - Ребенок 1: 2 + 3 = 5 конфет.
- - Ребенок 2: 3 + 3 = 6 конфет.
- - Ребенок 3: 5 + 3 = 8 конфет.
- - Ребенок 4: 1 + 3 = 4 конфеты.
- - Ребенок 5: 3 + 3 = 6 конфет.
- Наибольшее количество конфет после выдачи дополнительных — 8.

Результат: [false, false, true, false, false], так как только у третьего ребенка будет наибольшее количество конфет.


## Example 1:
```
Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true] 
Explanation: If you give all extraCandies to:
- Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
- Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
- Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
- Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
- Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
```

## Example 2:
```
Input: candies = [4,2,1,1,2], extraCandies = 1
Output: [true,false,false,false,false] 
Explanation: There is only 1 extra candy.
Kid 1 will always have the greatest number of candies, even if a different kid is given the extra candy.
```

## Example 3:
```
Input: candies = [12,1,12], extraCandies = 10
Output: [true,false,true]
```


## Constraints:
- n == candies.length
- 2 <= n <= 100
- 1 <= candies[i] <= 100
- 1 <= extraCandies <= 50
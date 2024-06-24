# 275. H-Index II


## Level - medium


## Task
Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper 
and citations is sorted in ascending order, return the researcher's h-index.

According to the definition of h-index on Wikipedia: 
The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.

You must write an algorithm that runs in logarithmic time.


## Объяснение
адача "275. H-Index II" на LeetCode требует от нас вычислить индекс Хирша (h-index) для ученого, 
основываясь на его списке публикаций, где каждая публикация имеет определенное количество цитирований. 
Важно, что список цитирований уже отсортирован в порядке убывания.

Индекс Хирша определяется как максимальное значение h, такое что у ученого есть не менее h публикаций, 
каждая из которых цитировалась не менее h раз.

Пример:  
Если у ученого есть публикации с цитированиями [10, 8, 5, 3, 3], то индекс Хирша равен 3, 
так как есть 3 публикации, которые цитировались не менее 3 раз.

Особенности задачи:
- Массив citations уже отсортирован в порядке убывания.
- Требуется найти индекс Хирша за время O(log n), что предполагает использование бинарного поиска.

Решение задачи с использованием бинарного поиска позволяет эффективно находить индекс Хирша, 
так как массив уже отсортирован.


## Example 1:
````
Input: citations = [0,1,3,5,6]
Output: 3
Explanation: [0,1,3,5,6] means the researcher has 5 papers in total and each of them had received 0, 1, 3, 5, 6 citations respectively.
Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.
````


## Example 2:
````
Input: citations = [1,2,100]
Output: 2
````


## Constraints:
- n == citations.length
- 1 <= n <= 105
- 0 <= citations[i] <= 1000
- citations is sorted in ascending order.
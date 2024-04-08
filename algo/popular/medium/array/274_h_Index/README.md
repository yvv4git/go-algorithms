# 274. H-Index


## Level - medium


## Task
Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, 
return the researcher's h-index.

According to the definition of h-index on Wikipedia: 
The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.


## Объяснение
Это задача, в которой необходимо найти H-индекс для данного списка цитирований. 
H-индекс определяется следующим образом: у индекса N есть N чисел, каждое из которых не меньше N. 
Например, для списка цитирований [3, 0, 6, 1, 5] H-индекс будет равен 3, потому что у трех чисел есть не меньше трех цитирований.

Основная идея решения заключается в том, чтобы отсортировать список цитирований в порядке убывания и пройтись по нему, 
сравнивая каждое цитирование с его индексом. Если цитирование меньше или равно индексу, то это и есть H-индекс.

Индекс цитирования (h-index) - это мера, используемая в научной литературе для оценки репутации ученых.
Он представляет собой максимальное число цитирований, на которое ученый может быть отнесен.
Например, если ученый имеет 10 публикаций, и каждая из них имеет не менее 10 цитирований, то его h-индекс равен 10. 
Если же ученый имеет 10 публикаций, но первая из них имеет 10 цитирований, а остальные - менее, 
то его h-индекс равен 1, потому что он имеет одну публикацию, которая цитируется 10 раз.


## Example 1:
````
Input: citations = [3,0,6,1,5]
Output: 3
Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.
````


## Example 2:
````
Input: citations = [1,3,1]
Output: 1
````


## Constraints:
- n == citations.length
- 1 <= n <= 5000
- 0 <= citations[i] <= 1000
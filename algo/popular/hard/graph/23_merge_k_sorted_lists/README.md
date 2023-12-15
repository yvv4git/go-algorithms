# 23. Merge k Sorted Lists


## Level - hard


## Task
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.


## Объяснение
В этой задаче вам нужно объединить k отсортированных списков в один отсортированный список. Каждый список представлен в виде связного списка, где каждый узел содержит целое число.

Вам нужно реализовать алгоритм, который будет принимать в качестве входных данных k отсортированных списков, объединять их в один отсортированный список и возвращать его.

Важно отметить, что вы можете решить эту задачу с помощью различных алгоритмов, 
таких как слияние двух отсортированных списков, сортировка слиянием, сортировка кучей и т.д. 
Выбор алгоритма будет зависеть от требований к производительности и сложности вашего решения.


## Example 1:
````
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
1->4->5,
1->3->4,
2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
````

## Example 2:
````
Input: lists = []
Output: []
````

## Example 3:
````
Input: lists = [[]]
Output: []
````

## Constraints:
- k == lists.length
- 0 <= k <= 104
- 0 <= lists[i].length <= 500
- -104 <= lists[i][j] <= 104
- lists[i] is sorted in ascending order.
- The sum of lists[i].length will not exceed 104.
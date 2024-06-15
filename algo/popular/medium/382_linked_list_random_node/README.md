# 382. Linked List Random Node


## Level - medium


## Task
Given a singly linked list, return a random node's value from the linked list. 
Each node must have the same probability of being chosen.

Implement the Solution class:
- Solution(ListNode head) Initializes the object with the head of the singly-linked list head.
- int getRandom() Chooses a node randomly from the list and returns its value. All the nodes of the list should be equally likely to be chosen.


## Объяснение
Задача заключается в разработке алгоритма для выбора случайного узла из связного списка с равной вероятностью. 
Связный список (Linked List) — это структура данных, в которой каждый элемент (узел) содержит данные и ссылку на следующий узел в списке.

Для решения этой задачи можно использовать алгоритм "резервуарная выборка" (Reservoir Sampling). 
Этот алгоритм позволяет выбирать k элементов из потока данных (в данном случае из связного списка) с равной вероятностью, 
даже если общее количество элементов заранее неизвестно. 
В данном случае k = 1, так как нам нужно выбрать только один случайный узел.

Этот алгоритм гарантирует, что каждый узел будет выбран с равной вероятностью, 
так как вероятность выбора любого узла на каждом шаге равна 1/i, где i — текущий индекс узла.


## Example 1:
![img.png](img.png)
````
Input
["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
[[[1, 2, 3]], [], [], [], [], []]
Output
[null, 1, 3, 2, 2, 3]

Explanation
Solution solution = new Solution([1, 2, 3]);
solution.getRandom(); // return 1
solution.getRandom(); // return 3
solution.getRandom(); // return 2
solution.getRandom(); // return 2
solution.getRandom(); // return 3
// getRandom() should return either 1, 2, or 3 randomly. Each element should have equal probability of returning.
````


## Constraints:
- The number of nodes in the linked list will be in the range [1, 10^4].
- -10^4 <= Node.val <= 10^4
- At most 104 calls will be made to getRandom.


## Follow up:
- What if the linked list is extremely large and its length is unknown to you?
- Could you solve this efficiently without using extra space?
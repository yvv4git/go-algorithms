# 435. Non-overlapping Intervals


## Level - medium


## Task
Given an array of intervals where intervals[i] = [start_i, end_i], 
return the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.


## Что надо сделать?
Надо найти количество интервалов, которые нужно удалить, чтобы наш список интервалов стал не пересекающимся.
Например: [[1,2],[2,3],[3,4],[1,3]], здесь можно удалить [1,3] и тогда интервалы не будут пересекаться.


## Example 1:
````
Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of the intervals are non-overlapping.
````


## Example 2:
````
Input: intervals = [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of the intervals non-overlapping.
````


## Constraints:
- 1 <= intervals.length <= 10^5
- intervals[i].length == 2
- -5 * 10^4 <= start_i < end_i <= 5 * 10^4
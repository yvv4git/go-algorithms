# 1450. Number of Students Doing Homework at a Given Time


## Level - easy


## Task
Given two integer arrays startTime and endTime and given an integer queryTime.

The ith student started doing their homework at the time startTime[i] and finished it at time endTime[i].

Return the number of students doing their homework at time queryTime. More formally, return the number of students where queryTime lays in the interval [startTime[i], endTime[i]] inclusive.


## Объяснение
Это задача на понимание работы с интервалами времени и подсчёта количества событий, которые происходят в определённый момент.
Даны два массива:
- startTime — массив, где каждый элемент startTime[i] обозначает время начала выполнения домашней работы i-м студентом.
- endTime — массив, где каждый элемент endTime[i] обозначает время окончания выполнения домашней работы i-м студентом.

Также дано целое число queryTime, которое обозначает конкретный момент времени. 
Задача состоит в том, чтобы определить, сколько студентов выполняли домашнюю работу в момент времени queryTime.

Пример:
- startTime = [1, 2, 3]
- endTime = [3, 2, 7]
- queryTime = 4

Таким образом, только один студент (третий) выполнял домашнюю работу в момент времени 4. Ответ: 1.

Возможное решение:
- Пройдите по каждому студенту.
- Для каждого студента проверьте, находится ли queryTime в интервале [startTime[i], endTime[i]].
- Если да, увеличьте счётчик.
- Верните итоговый счётчик.


## Example 1:
```
Input: startTime = [1,2,3], endTime = [3,2,7], queryTime = 4
Output: 1
Explanation: We have 3 students where:
The first student started doing homework at time 1 and finished at time 3 and wasn't doing anything at time 4.
The second student started doing homework at time 2 and finished at time 2 and also wasn't doing anything at time 4.
The third student started doing homework at time 3 and finished at time 7 and was the only student doing homework at time 4.
```


## Example 2:
```
Input: startTime = [4], endTime = [4], queryTime = 4
Output: 1
Explanation: The only student was doing their homework at the queryTime.
```


## Constraints:
- startTime.length == endTime.length
- 1 <= startTime.length <= 100
- 1 <= startTime[i] <= endTime[i] <= 1000
- 1 <= queryTime <= 1000
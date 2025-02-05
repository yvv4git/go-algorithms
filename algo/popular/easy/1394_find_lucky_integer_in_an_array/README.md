# 1394. Find Lucky Integer in an Array


## Level - easy


## Task
Given an array of integers arr, a lucky integer is an integer that has a frequency in the array equal to its value.

Return the largest lucky integer in the array. If there is no lucky integer return -1.


## Объяснение
Cуть заключается в поиске "счастливого числа" в массиве. 
Дан массив целых чисел arr. Необходимо найти "счастливое число" в этом массиве. "Счастливое число" определяется следующим образом:
- Число x считается счастливым, если его частота (количество вхождений) в массиве равна самому числу x.
- Если в массиве несколько счастливых чисел, нужно вернуть наибольшее из них.
- Если счастливых чисел нет, вернуть -1.


## Example 1:
```
Input: arr = [2,2,3,4]
Output: 2
Explanation: The only lucky number in the array is 2 because frequency[2] == 2.
```


## Example 2:
```
Input: arr = [1,2,2,3,3,3]
Output: 3
Explanation: 1, 2 and 3 are all lucky numbers, return the largest of them.
```


Example 3:
```
Input: arr = [2,2,2,3,3]
Output: -1
Explanation: There are no lucky numbers in the array.
```

## Constraints:
- 1 <= arr.length <= 500
- 1 <= arr[i] <= 500
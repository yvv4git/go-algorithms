# 1103. Distribute Candies to People


## Level - easy
We distribute some number of candies, to a row of n = num_people people in the following way:

We then give 1 candy to the first person, 2 candies to the second person, and so on until we give n candies to the last person.

Then, we go back to the start of the row, giving n + 1 candies to the first person, n + 2 candies to the second person, and so on until we give 2 * n candies to the last person.

This process repeats (with us giving one more candy each time, and moving to the start of the row after we reach the end) until we run out of candies.  
The last person will receive all of our remaining candies (not necessarily one more than the previous gift).

Return an array (of length num_people and sum candies) that represents the final distribution of candies.


## Объяснение
Задача описывает процесс распределения конфет между группой людей. 
У нас есть определенное количество конфет candies, которые нужно распределить между num_people людьми, стоящими в ряд. 
Конфеты распределяются следующим образом:
1. Первый человек получает 1 конфету, второй — 2 конфеты, и так далее, пока не дойдет до последнего человека, который получает num_people конфет.
2. После этого процесс повторяется, начиная с первого человека, но на этот раз каждый человек получает на одну конфету больше, чем в предыдущем цикле. 
Например, первый человек получает num_people + 1 конфет, второй — num_people + 2, и так далее.
3. Процесс продолжается до тех пор, пока конфеты не закончатся. 
Если в конце конфет недостаточно для того, чтобы дать следующему человеку его "полное" количество, 
он получает оставшиеся конфеты, и распределение заканчивается.

Необходимо вернуть массив длиной num_people, где каждый элемент представляет собой общее количество конфет, 
полученных каждым человеком после завершения распределения.


## Example 1:
```
Input: candies = 7, num_people = 4
Output: [1,2,3,1]
Explanation:
On the first turn, ans[0] += 1, and the array is [1,0,0,0].
On the second turn, ans[1] += 2, and the array is [1,2,0,0].
On the third turn, ans[2] += 3, and the array is [1,2,3,0].
On the fourth turn, ans[3] += 1 (because there is only one candy left), and the final array is [1,2,3,1].
```


## Example 2:
```
Input: candies = 10, num_people = 3
Output: [5,2,3]
Explanation: 
On the first turn, ans[0] += 1, and the array is [1,0,0].
On the second turn, ans[1] += 2, and the array is [1,2,0].
On the third turn, ans[2] += 3, and the array is [1,2,3].
On the fourth turn, ans[0] += 4, and the final array is [5,2,3].
```


## Constraints:
- 1 <= candies <= 10^9
- 1 <= num_people <= 1000
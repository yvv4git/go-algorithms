# 1177. Can Make Palindrome from Substring


## Level - medium


## Task
You are given a string s and array queries where queries[i] = [lefti, righti, ki]. 
We may rearrange the substring s[left_i...right_i] for each query and then choose up to ki of them to replace with any lowercase English letter.

If the substring is possible to be a palindrome string after the operations above, the result of the query is true. 
Otherwise, the result is false.

Return a boolean array answer where answer[i] is the result of the ith query queries[i].

Note that each letter is counted individually for replacement, so if, for example s[lefti...righti] = "aaa", and ki = 2, 
we can only replace two of the letters. Also, note that no query modifies the initial string s.


## Объяснение
Задача состоит в том, чтобы определить, можно ли составить из каждой подстроки, образованной строкой s, 
палиндром с помощью операций замены символов. 
Операция замены символов позволяет заменить любой символ на любой другой символ.

Данная задача предполагает написание функции canMakePaliQueriesV1, которая принимает на вход строку s и массив запросов queries. 
Каждый запрос представляет собой массив из трех элементов: [left, right, k]. 
Здесь left и right - это индексы начала и конца подстроки в строке s, а k - это количество операций замены символов.

Функция должна возвращать массив булевых значений той же длины, что и массив запросов. В i-м элементе этого массива должно быть true, 
если из подстроки s[queries[i][0]:queries[i][1]] можно составить палиндром с помощью не более чем queries[i][2] операций замены символов, 
и false в противном случае.

## Example :
````
Input: s = "abcda", queries = [[3,3,0],[1,2,0],[0,3,1],[0,3,2],[0,4,1]]
Output: [true,false,false,true,true]
Explanation:
queries[0]: substring = "d", is palidrome.
queries[1]: substring = "bc", is not palidrome.
queries[2]: substring = "abcd", is not palidrome after replacing only 1 character.
queries[3]: substring = "abcd", could be changed to "abba" which is palidrome. Also this can be changed to "baab" first rearrange it "bacd" then replace "cd" with "ab".
queries[4]: substring = "abcda", could be changed to "abcba" which is palidrome.
````

## Example 2:
````
Input: s = "lyb", queries = [[0,1,0],[2,2,1]]
Output: [false,true]
````


## Constraints:
- 1 <= s.length, queries.length <= 10^5
- 0 <= left_i <= right_i < s.length
- 0 <= k_i <= s.length 
- s consists of lowercase English letters.
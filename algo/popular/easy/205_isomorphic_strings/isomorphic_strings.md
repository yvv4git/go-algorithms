# 205. Isomorphic Strings


## Level - easy


## Task
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.


## Объяснение
Две струны, X а также Y, называются изоморфными, если все вхождения каждого символа в X можно заменить другим символом, 
чтобы получить Y наоборот. Например, рассмотрим строки ACAB а также XCXY.
Они изоморфны, поскольку мы можем отобразить 'A' —> 'X', 'B' —> 'Y' а также 'C' —> 'C'.

Example 1:
````
Input: s = "egg", t = "add"
Output: true
````


Example 2:
````
Input: s = "foo", t = "bar"
Output: false
````

Example 3:
````
Input: s = "paper", t = "title"
Output: true
````


## Constraints:
- 1 <= s.length <= 5 * 104
- t.length == s.length
- s and t consist of any valid ascii character.
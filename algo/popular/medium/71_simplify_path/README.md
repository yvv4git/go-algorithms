# 71. Simplify Path


## Level - medium


## Task
Given a string path, which is an absolute path (starting with a slash '/') to a file or directory in a Unix-style file system, 
convert it to the simplified canonical path.

In a Unix-style file system, a period '.' refers to the current directory, a double period '..' refers to the directory up a level, 
and any multiple consecutive slashes (i.e. '//') are treated as a single slash '/'. 
For this problem, any other format of periods such as '...' are treated as file/directory names.

The canonical path should have the following format:
- The path starts with a single slash '/'.
- Any two directories are separated by a single slash '/'.
- The path does not end with a trailing '/'.
- The path only contains the directories on the path from the root directory to the target file or directory (i.e., no period '.' or double period '..')

Return the simplified canonical path.


## Объяснение
Задача  заключается в том, чтобы упростить абсолютный путь в файловой системе Unix-подобной. 
Путь может содержать символы "." для текущей директории и ".." для родительской директории.

Примеры:
- "/home/" => "/home" (Нет необходимости в упрощении)
- "/../" => "/" (Возврат в корень)
- "/home//foo/" => "/home/foo" (Удаление лишних слешей)
- "/a/./b/../../c/" => "/c" (Переход назад по директориям)

Цель задачи - написать функцию, которая принимает на вход строку, представляющую путь в файловой системе, и возвращает упрощенный путь. 
Упрощенный путь должен быть как можно короче, без лишних символов и с учетом переходов назад по директориям.


## Example 1:
````
Input: path = "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
````


## Example 2:
````
Input: path = "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
````


## Example 3:
````
Input: path = "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
````


## Constraints:
- 1 <= path.length <= 3000
- path consists of English letters, digits, period '.', slash '/' or '_'.
- path is a valid absolute Unix path.
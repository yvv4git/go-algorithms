# 551. Student Attendance Record I


## Level = easy


## Task
You are given a string s representing an attendance record for a student where each character signifies whether the student was absent, 
late, or present on that day. The record only contains the following three characters:
- 'A': Absent.
- 'L': Late.
- 'P': Present.

The student is eligible for an attendance award if they meet both of the following criteria:
- The student was absent ('A') for strictly fewer than 2 days total.
- The student was never late ('L') for 3 or more consecutive days.

Return true if the student is eligible for an attendance award, or false otherwise.


## Объяснение
Это задача, которая часто встречается в соревнованиях по программированию 
и часто используется для оценки навыков работы с строками и условными операторами.

Задача состоит в следующем:
Учитывая строку, которая обозначает последовательность записей присутствия и отсутствия студента, вернуть true, 
если студент не может быть принят в университет, и false, если студент может быть принят.

Студент не может быть принят, если он имеет более двух последовательных дней отсутствия ('A') 
или более трех несвоевременных посещений ('L').

Например, если входная строка - "PPALLP", то выход должен быть true, 
потому что студент не имеет более двух последовательных дней отсутствия и не имеет более трех несвоевременных посещений.

Если входная строка - "PPALLL", то выход должен быть false, потому что студент имеет три несвоевременных посещения.


## Example 1:
````
Input: s = "PPALLP"
Output: true
Explanation: The student has fewer than 2 absences and was never late 3 or more consecutive days.
````


## Example 2:
````
Input: s = "PPALLL"
Output: false
Explanation: The student was late 3 consecutive days in the last 3 days, so is not eligible for the award.
````


## Constraints:
- 1 <= s.length <= 1000
- s[i] is either 'A', 'L', or 'P'.
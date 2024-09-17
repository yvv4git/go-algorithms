/*
# 627. Swap Salary

Table: Salary

+-------------+----------+
| Column Name | Type     |
+-------------+----------+
| id          | int      |
| name        | varchar  |
| sex         | ENUM     |
| salary      | int      |
+-------------+----------+
id is the primary key (column with unique values) for this table.
The sex column is ENUM (category) value of type ('m', 'f').
The table contains information about an employee.



Write a solution to swap all 'f' and 'm' values (i.e., change all 'f' values to 'm' and vice versa) with a single update statement 
and no intermediate temporary tables.

Note that you must write a single update statement, do not write any select statement for this problem.

The result format is in the following example
*/

/*
# Объяснение.
Задача 627 "Swap Salary" (Обменить зарплату) заключается в том, чтобы поменять местами зарплаты мужчин и женщин в таблице "salaries" (зарплаты).

Таблица "salaries" содержит информацию о зарплатах сотрудников, включая их идентификатор (id), пол (sex) и зарплату (salary).

Задача состоит в том, чтобы написать SQL-запрос, который поменяет местами зарплаты мужчин и женщин,
т.е. мужчинам присвоить зарплату женщин, а женщинам - зарплату мужчин.
*/


/*
Example 1:

Input: 
Salary table:
+----+------+-----+--------+
| id | name | sex | salary |
+----+------+-----+--------+
| 1  | A    | m   | 2500   |
| 2  | B    | f   | 1500   |
| 3  | C    | m   | 5500   |
| 4  | D    | f   | 500    |
+----+------+-----+--------+
Output: 
+----+------+-----+--------+
| id | name | sex | salary |
+----+------+-----+--------+
| 1  | A    | f   | 2500   |
| 2  | B    | m   | 1500   |
| 3  | C    | f   | 5500   |
| 4  | D    | m   | 500    |
+----+------+-----+--------+
Explanation: 
(1, A) and (3, C) were changed from 'm' to 'f'.
(2, B) and (4, D) were changed from 'f' to 'm'.
*/


-- Prepare
DROP TABLE IF EXISTS `salaries`;

CREATE TABLE salary (
  id INT PRIMARY KEY,
  sex CHAR(1),
  salary INT
);

INSERT INTO salary (id, sex, salary)
VALUES
  (1, 'M', 5000),
  (2, 'F', 6000),
  (3, 'M', 7000),
  (4, 'F', 8000);

select * from salary;

-- Solution
UPDATE salary
SET sex = CASE 
              WHEN sex = 'm' THEN 'f'
              WHEN sex = 'f' THEN 'm'
          END;

select * from salary;
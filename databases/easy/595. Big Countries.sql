/*
# 595. Big Countries

Table: World

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| name        | varchar |
| continent   | varchar |
| area        | int     |
| population  | int     |
| gdp         | bigint  |
+-------------+---------+
name is the primary key (column with unique values) for this table.
Each row of this table gives information about the name of a country, the continent to which it belongs, its area, the population, and its GDP value.

A country is big if:
- it has an area of at least three million (i.e., 3000000 km2), or
- it has a population of at least twenty-five million (i.e., 25000000).
Write a solution to find the name, population, and area of the big countries.

Return the result table in any order.

The result format is in the following example.


*/


/*
# Example 1:

Input: 
World table:
+-------------+-----------+---------+------------+--------------+
| name        | continent | area    | population | gdp          |
+-------------+-----------+---------+------------+--------------+
| Afghanistan | Asia      | 652230  | 25500100   | 20343000000  |
| Albania     | Europe    | 28748   | 2831741    | 12960000000  |
| Algeria     | Africa    | 2381741 | 37100000   | 188681000000 |
| Andorra     | Europe    | 468     | 78115      | 3712000000   |
| Angola      | Africa    | 1246700 | 20609294   | 100990000000 |
+-------------+-----------+---------+------------+--------------+
Output: 
+-------------+------------+---------+
| name        | population | area    |
+-------------+------------+---------+
| Afghanistan | 25500100   | 652230  |
| Algeria     | 37100000   | 2381741 |
+-------------+------------+---------+
*/

/*

*/

/*
# Объяснение

Вам предоставлена таблица World, которая содержит информацию о различных странах, 
включая их название, континент, площадь, население и ВВП.

Необходимо найти страны, которые считаются "большими". 
Страна считается большой, если она соответствует хотя бы одному из следующих критериев:
- Площадь страны составляет не менее 3 миллионов квадратных километров (3000000 км²).
- Население страны составляет не менее 25 миллионов человек (25000000).

Вам нужно вернуть таблицу, содержащую название страны, её население и площадь для всех стран, 
которые соответствуют критериям "большой страны". Порядок строк в результирующей таблице не важен.
*/

-- PREPARE
CREATE TABLE World (
    name VARCHAR(255) PRIMARY KEY,
    continent VARCHAR(255),
    area INT,
    population INT,
    gdp BIGINT
);

INSERT INTO World (name, continent, area, population, gdp) VALUES
('Afghanistan', 'Asia', 652230, 25500100, 20343000000),
('Albania', 'Europe', 28748, 2831741, 12960000000),
('Algeria', 'Africa', 2381741, 37100000, 188681000000),
('Andorra', 'Europe', 468, 78115, 3712000000),
('Angola', 'Africa', 1246700, 20609294, 100990000000);


-- SOLUTION 
SELECT name, population, area
FROM World
WHERE area >= 3000000 OR population >= 25000000;
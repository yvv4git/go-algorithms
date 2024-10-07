/*
# 511. Game Play Analysis I

Table: Activity

+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| player_id    | int     |
| device_id    | int     |
| event_date   | date    |
| games_played | int     |
+--------------+---------+
(player_id, event_date) is the primary key (combination of columns with unique values) of this table.
This table shows the activity of players of some games.
Each row is a record of a player who logged in and played a number of games (possibly 0) before logging out on someday using some device.
*/

/*
Write a solution to find the first login date for each player.

Return the result table in any order.

The result format is in the following example.
*/


/*
Example 1:

Input: 
Activity table:
+-----------+-----------+------------+--------------+
| player_id | device_id | event_date | games_played |
+-----------+-----------+------------+--------------+
| 1         | 2         | 2016-03-01 | 5            |
| 1         | 2         | 2016-05-02 | 6            |
| 2         | 3         | 2017-06-25 | 1            |
| 3         | 1         | 2016-03-02 | 0            |
| 3         | 4         | 2018-07-03 | 5            |
+-----------+-----------+------------+--------------+
Output: 
+-----------+-------------+
| player_id | first_login |
+-----------+-------------+
| 1         | 2016-03-01  |
| 2         | 2017-06-25  |
| 3         | 2016-03-02  |
+-----------+-------------+
*/

/*
# Объяснение

Вам предоставлена таблица Activity, которая содержит информацию о действиях игроков в игре. 
Каждая строка таблицы представляет собой запись о том, 
что игрок с определенным player_id залогинился на устройство с device_id в конкретную дату event_date 
и сыграл определенное количество игр (games_played).

Необходимо написать SQL-запрос, который найдет первую дату входа (первый event_date) для каждого игрока.
*/


-- Prepare
CREATE TABLE Activity (
    player_id INT,
    device_id INT,
    event_date DATE,
    games_played INT,
    PRIMARY KEY (player_id, event_date)
);

INSERT INTO Activity (player_id, device_id, event_date, games_played) VALUES
(1, 2, '2016-03-01', 5),
(1, 2, '2016-05-02', 6),
(2, 3, '2017-06-25', 1),
(3, 1, '2016-03-02', 0),
(3, 4, '2018-07-03', 5);


-- Solution
SELECT 
    player_id, 
    MIN(event_date) AS first_login
FROM 
    Activity
GROUP BY 
    player_id;

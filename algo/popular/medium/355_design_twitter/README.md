# 355. Design Twitter


## Level - medium


## Task
Design a simplified version of Twitter where users can post tweets, follow/unfollow another user, 
and is able to see the 10 most recent tweets in the user's news feed.

Implement the Twitter class:
1. Twitter() Initializes your twitter object.
2. void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId. Each call to this function will be made with a unique tweetId.
3. List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user themself. Tweets must be ordered from most recent to least recent.
4. void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
5. void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.


## Объснение
Задача требует от вас разработать систему, которая имитирует функциональность Twitter. 
Вам нужно создать класс Twitter с несколькими методами, которые позволяют пользователям:
1. Публиковать твиты.
2. Подписываться на других пользователей.
3. Отписываться от других пользователей.
4. Видеть ленту новостей, содержащую последние твиты от пользователей, на которых они подписаны.

Вот основные методы, которые должны быть реализованы:
1. postTweet(userId, tweetId): Публикует новый твит от userId с идентификатором tweetId.
2. getNewsFeed(userId): Возвращает список последних твитов в ленте новостей userId. Лента должна содержать твиты от пользователя и от тех, на кого он подписан. Твиты должны быть отсортированы по времени публикации в порядке убывания.
3. follow(followerId, followeeId): Пользователь followerId подписывается на пользователя followeeId.
4. unfollow(followerId, followeeId): Пользователь followerId отписывается от пользователя followeeId.


## Example 1:
````
Input
["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
Output
[null, null, [5], null, null, [6, 5], null, [5]]

Explanation
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
twitter.follow(1, 2);    // User 1 follows user 2.
twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
twitter.unfollow(1, 2);  // User 1 unfollows user 2.
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.
````


## Constraints:
- 1 <= userId, followerId, followeeId <= 500
- 0 <= tweetId <= 104
- All the tweets have unique IDs.
- At most 3 * 104 calls will be made to postTweet, getNewsFeed, follow, and unfollow.
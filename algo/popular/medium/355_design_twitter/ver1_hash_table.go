package main

import (
	"container/heap"
	"fmt"
)

// Tweet представляет собой твит с идентификатором и временной меткой
type Tweet struct {
	ID        int
	Timestamp int
}

// TweetHeap представляет собой минимальную кучу твитов
type TweetHeap []Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].Timestamp < h[j].Timestamp }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TweetHeap) Push(x interface{}) {
	*h = append(*h, x.(Tweet))
}

func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Twitter представляет собой структуру для моделирования функциональности Twitter
type Twitter struct {
	userTweets  map[int][]Tweet      // Хранение твитов пользователей
	userFollows map[int]map[int]bool // Хранение подписок пользователей
	timestamp   int                  // Временная метка для твитов
}

// Constructor создает новый объект Twitter
func Constructor() Twitter {
	return Twitter{
		userTweets:  make(map[int][]Tweet),
		userFollows: make(map[int]map[int]bool),
		timestamp:   0,
	}
}

// PostTweet публикует новый твит от userId с идентификатором tweetId
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, exists := this.userTweets[userId]; !exists {
		this.userTweets[userId] = []Tweet{}
	}
	this.userTweets[userId] = append(this.userTweets[userId], Tweet{ID: tweetId, Timestamp: this.timestamp})
	this.timestamp++
}

// GetNewsFeed возвращает список последних твитов в ленте новостей userId
// Временная сложность: O(n log k), где n - общее количество твитов, k - количество твитов в ленте (10)
// Пространственная сложность: O(k)
func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &TweetHeap{}
	heap.Init(h)

	// Добавляем твиты пользователя
	if tweets, exists := this.userTweets[userId]; exists {
		for _, tweet := range tweets {
			heap.Push(h, tweet)
			if h.Len() > 10 {
				heap.Pop(h)
			}
		}
	}

	// Добавляем твиты от пользователей, на которых подписан userId
	if follows, exists := this.userFollows[userId]; exists {
		for followeeId := range follows {
			if tweets, exists := this.userTweets[followeeId]; exists {
				for _, tweet := range tweets {
					heap.Push(h, tweet)
					if h.Len() > 10 {
						heap.Pop(h)
					}
				}
			}
		}
	}

	// Собираем результат
	result := make([]int, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		result[i] = heap.Pop(h).(Tweet).ID
	}
	return result
}

// Follow подписывает followerId на followeeId
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, exists := this.userFollows[followerId]; !exists {
		this.userFollows[followerId] = make(map[int]bool)
	}
	this.userFollows[followerId][followeeId] = true
}

// Unfollow отписывает followerId от followeeId
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if follows, exists := this.userFollows[followerId]; exists {
		delete(follows, followeeId)
	}
}

func main() {
	/*
		METHOD: Hash table
		Задача решена с использованием хэш-таблиц для хранения твитов и подписок,
		а также минимальной кучи (heap) для эффективного получения последних твитов.
		Куча позволяет поддерживать только 10 последних твитов, что уменьшает временную сложность операции GetNewsFeed.

		TIME COMPLEXITY:
		- PostTweet - O(1) - добавление твита в конец списка.
		- GetNewsFeed - O(n log k), где n - общее количество твитов, k - количество твитов в ленте (10).
		- Follow - O(1) - добавление подписки в хэш-таблицу.
		- Unfollow - O(1) - удаление подписки из хэш-таблицы.

		SPACE COMPLEXITY:
		- PostTweet -  O(1) - дополнительная память не требуется.
		- GetNewsFeed - O(k) - хранение k твитов в куче.
		- Follow - O(1) - дополнительная память не требуется.
		- Unfollow - O(1) - дополнительная память не требуется.
	*/
	twitter := Constructor()

	// Пользователь 1 публикует новый твит
	twitter.PostTweet(1, 5)

	// Пользователь 1 получает свою ленту новостей
	fmt.Println(twitter.GetNewsFeed(1)) // Вывод: [5]

	// Пользователь 1 подписывается на пользователя 2
	twitter.Follow(1, 2)

	// Пользователь 2 публикует новый твит
	twitter.PostTweet(2, 6)

	// Пользователь 1 получает свою ленту новостей
	fmt.Println(twitter.GetNewsFeed(1)) // Вывод: [6, 5]

	// Пользователь 1 отписывается от пользователя 2
	twitter.Unfollow(1, 2)

	// Пользователь 1 получает свою ленту новостей
	fmt.Println(twitter.GetNewsFeed(1)) // Вывод: [5]
}

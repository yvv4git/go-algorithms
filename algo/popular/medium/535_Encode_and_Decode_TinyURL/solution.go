package main

import (
	"fmt"
	"strconv"
)

/*
APPROACH: Использование хэш-карт для хранения соответствий между длинными и короткими URL. Короткие URL генерируются с помощью счетчика.
TIME COMPLEXITY: O(1) для операций encode и decode, так как доступ к map в среднем занимает постоянное время.
SPACE COMPLEXITY: O(n), где n - количество уникальных URL, хранящихся в системе.
*/
func main() {
	obj := Constructor()
	url := "https://leetcode.com/problems/design-tinyurl"
	tiny := obj.encode(url)
	ans := obj.decode(tiny)
	fmt.Printf("Original: %s\nTiny: %s\nDecoded: %s\n", url, tiny, ans)
}

type Codec struct {
	shortToLong map[string]string
	longToShort map[string]string
	counter     int
}

func Constructor() Codec {
	return Codec{
		shortToLong: make(map[string]string),
		longToShort: make(map[string]string),
		counter:     0,
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if short, ok := this.longToShort[longUrl]; ok {
		return short
	}

	this.counter++
	short := "http://tinyurl.com/" + strconv.Itoa(this.counter)
	this.shortToLong[short] = longUrl
	this.longToShort[longUrl] = short

	return short
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.shortToLong[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

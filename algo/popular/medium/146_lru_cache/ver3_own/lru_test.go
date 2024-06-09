package ver3_own

import (
	"fmt"
	"testing"
)

func TestLRUv3(t *testing.T) {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // вернет 1
	lru.Put(3, 3)
	fmt.Println(lru.Get(2)) // вернет -1 (не найдено)
	lru.Put(4, 4)
	fmt.Println(lru.Get(1)) // вернет -1 (не найдено)
	fmt.Println(lru.Get(3)) // вернет 3
	fmt.Println(lru.Get(4)) // вернет 4
}

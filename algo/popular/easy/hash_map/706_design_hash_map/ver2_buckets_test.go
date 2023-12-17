package _06_design_hash_map

import (
	"fmt"
	"testing"
)

func TestHashMapV2(t *testing.T) {
	// Создаем новый объект MyHashMap
	hashMap := Constructor2()

	// Вставляем пары ключ-значение
	hashMap.Put(1, 1)
	hashMap.Put(2, 2)

	// Получаем значение по ключу
	fmt.Println(hashMap.Get(1)) // Выведет: 1
	fmt.Println(hashMap.Get(3)) // Выведет: -1, так как ключ 3 не существует

	// Обновляем значение по ключу
	hashMap.Put(2, 1)
	fmt.Println(hashMap.Get(2)) // Выведет: 1

	// Удаляем пару ключ-значение
	hashMap.Remove(2)
	fmt.Println(hashMap.Get(2)) // Выведет: -1, так как ключ 2 был удален
}

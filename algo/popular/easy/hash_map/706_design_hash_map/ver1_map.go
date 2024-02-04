package _06_design_hash_map

import "sync"

type MyHashMap struct {
	m     map[int]int
	mutex sync.RWMutex
}

// Constructor - crate instance.
func Constructor() MyHashMap {
	return MyHashMap{m: make(map[int]int)}
}

// Put - used for add element.
// TIME COMPLEXITY: O(1)
func (m *MyHashMap) Put(key int, value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.m[key] = value
}

// Get - used for getting element.
// TIME COMPLEXITY: O(1)
func (m *MyHashMap) Get(key int) int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if val, ok := m.m[key]; ok {
		return val
	}
	return -1
}

// Remove - used for remove element.
// TIME COMPLEXITY: O(1)
func (m *MyHashMap) Remove(key int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.m, key)
}

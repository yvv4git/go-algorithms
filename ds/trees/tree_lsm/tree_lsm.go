package main

import "sort"

// LSM - структура, представляющая LSM-дерево
type LSM struct {
	memTable    map[string]string   // временная структура для хранения пар ключ-значение
	sstableList []map[string]string // список SSTable (Sorted String Table)
}

// NewLSM - конструктор для создания нового LSM-дерева
func NewLSM() *LSM {
	return &LSM{
		memTable:    make(map[string]string),      // временная сложность O(1)
		sstableList: make([]map[string]string, 0), // временная сложность O(1)
	}
}

// Put - метод для добавления пары ключ-значение в LSM-дерево
// Временная сложность O(1), пространственная сложность O(n)
func (lsm *LSM) Put(key string, value string) {
	lsm.memTable[key] = value   // временная сложность O(1)
	if len(lsm.memTable) > 10 { // временная сложность O(1)
		lsm.sstableList = append(lsm.sstableList, lsm.memTable) // временная сложность O(1), пространственная сложность O(n)
		lsm.memTable = make(map[string]string)                  // временная сложность O(1)
	}
}

// Get - метод для получения значения по ключу из LSM-дерева
// Временная сложность O(n), пространственная сложность O(1)
func (lsm *LSM) Get(key string) (string, bool) {
	value, ok := lsm.memTable[key] // временная сложность O(1)
	if ok {
		return value, true // временная сложность O(1)
	}

	for i := len(lsm.sstableList) - 1; i >= 0; i-- { // временная сложность O(n)
		value, ok := lsm.sstableList[i][key] // временная сложность O(1)
		if ok {
			return value, true // временная сложность O(1)
		}
	}

	return "", false // временная сложность O(1)
}

// Merge - метод для слияния временной структуры с SSTable
// Временная сложность O(n log n), пространственная сложность O(n)
func (lsm *LSM) Merge() {
	keys := make([]string, 0, len(lsm.memTable)) // временная сложность O(n)
	for k := range lsm.memTable {                // временная сложность O(n)
		keys = append(keys, k) // временная сложность O(1)
	}
	sort.Strings(keys) // временная сложность O(n log n)

	newSSTable := make(map[string]string) // временная сложность O(1)
	for _, k := range keys {              // временная сложность O(n)
		newSSTable[k] = lsm.memTable[k] // временная сложность O(1)
	}

	lsm.sstableList = append(lsm.sstableList, newSSTable) // временная сложность O(1), пространственная сложность O(n)
	lsm.memTable = make(map[string]string)                // временная сложность O(1)
}

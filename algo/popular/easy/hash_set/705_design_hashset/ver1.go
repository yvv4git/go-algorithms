package _05_design_hashset

type MyHashSet struct {
	keyMap map[int]bool
}

// Constructor time complexity: O(1)
func Constructor() MyHashSet {
	return MyHashSet{
		keyMap: make(map[int]bool),
	}
}

// Add time complexity: O(1)
func (h *MyHashSet) Add(key int) {
	if _, ok := h.keyMap[key]; !ok {
		h.keyMap[key] = true
	}
}

// Remove time complexity: O(1) - best case, O(n) - worst case.
func (h *MyHashSet) Remove(key int) {
	if _, ok := h.keyMap[key]; ok {
		delete(h.keyMap, key)
	}
}

// Contains time complexity: O(1)
func (h *MyHashSet) Contains(key int) bool {
	_, ok := h.keyMap[key]
	return ok
}

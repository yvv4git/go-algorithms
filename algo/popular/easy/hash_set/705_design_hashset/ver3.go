package _05_design_hashset

const limit = 1_000_001

type MyHashSetV3 struct {
	payload [limit]bool
}

func ConstructorV3() MyHashSetV3 {
	return MyHashSetV3{payload: [limit]bool{}}
}

func (h *MyHashSetV3) Add(key int) {
	h.payload[key] = true
}

func (h *MyHashSetV3) Remove(key int) {
	h.payload[key] = false
}

func (h *MyHashSetV3) Contains(key int) bool {
	return h.payload[key]
}

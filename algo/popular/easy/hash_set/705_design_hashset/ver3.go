package _05_design_hashset

type MyHashSetV3 struct {
	Values [1_000_001]bool
}

func ConstructorV3() MyHashSetV3 {
	return MyHashSetV3{Values: [1_000_001]bool{}}
}

func (h *MyHashSetV3) Add(key int) {
	h.Values[key] = true
}

func (h *MyHashSetV3) Remove(key int) {
	h.Values[key] = false
}

func (h *MyHashSetV3) Contains(key int) bool {
	return h.Values[key]
}

package _05_design_hashset

type MyHashSetV2 struct {
	Val  int
	Node *MyHashSetV2
}

func ConstructorV2() MyHashSetV2 {
	return MyHashSetV2{Val: -1, Node: nil}
}

func (h *MyHashSetV2) Add(key int) {
	if h.Val == -1 {
		h.Val = key
		return
	} else if h.Val == key {
		return
	}
	for h.Node != nil {
		if h.Node.Val == key {
			return
		}
		h = h.Node
	}
	if h.Val == key {
		return
	}
	h.Node = &MyHashSetV2{Val: key, Node: nil}
}

func (h *MyHashSetV2) Remove(key int) {
	if h.Val == key {
		if h.Node != nil {
			h.Val = h.Node.Val
			h.Node = h.Node.Node
		} else {
			h.Val = -1
		}
		return
	}

	for h.Node != nil {
		if h.Node.Val == key {
			h.Node = h.Node.Node
			return
		}
		h = h.Node
	}
}

func (h *MyHashSetV2) Contains(key int) bool {
	for h != nil {
		if h.Val == key {
			return true
		}
		h = h.Node
	}
	return false
}

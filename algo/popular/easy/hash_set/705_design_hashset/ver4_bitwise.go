package _05_design_hashset

/*
Решение не подходит, так как uint располагает только 32 битами, uint64 только 64.
В общем, нет такого типа данных, если только использовать слайс.
*/

type MyHashSetV4 struct {
	set uint
}

// ConstructorV4 - создает новый MyHashSetV4
func ConstructorV4() MyHashSetV4 {
	return MyHashSetV4{0}
}

// Add - добавляет элемент в MyHashSetV4
func (h *MyHashSetV4) Add(element int) {
	h.set = h.set | (1 << element)
}

// Remove - удаляет элемент из MyHashSetV4
func (h *MyHashSetV4) Remove(element int) {
	h.set = h.set & ^(1 << element)
}

// Contains - проверяет, содержится ли элемент в MyHashSetV4
func (h *MyHashSetV4) Contains(element int) bool {
	return (h.set & (1 << element)) != 0
}

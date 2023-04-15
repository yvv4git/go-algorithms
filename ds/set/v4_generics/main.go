package main

/*import (
"fmt"
)

type Set(type T comparable) map[T]bool

func New(type T comparable)() Set(T) {
	return make(Set(T))
}

func (s Set(T)) Add(values ...T) {
	for _, value := range values {
		s[value] = true
	}
}

func (s Set(T)) Delete(values ...T) {
for _, value := range values {
delete(s, value)
}
}

func (s Set(T)) Len() int {
return len(s)
}

func (s Set(T)) Has(value T) bool {
_, ok := s[value]
return ok
}

func (s Set(T)) Iterate(it func(T)) {
	for v := range s {
		it(v)
	}
}

func (s Set(T)) Values() []T {
	var values []T
	s.Iterate(func(value T) {
		values = append(values, value)
	})
	return values
}

func (s Set(T)) Clone() Set(T) {
	set := make(Set(T))
	set.Add(s.Values()...)
	return set
}

func (s Set(T)) Union(other Set(T)) Set(T) {
	set := s.Clone()
	set.Add(other.Values()...)
	return set
}

func (s Set(T)) Intersection(other Set(T)) Set(T) {
	set := make(Set(T))
	s.Iterate(func(value T) {
		if other.Has(value) {
			set.Add(value)
		}
	})
	return set
}


func main() {
	it := func(val int){
		fmt.Println(val)
	}
	set1 := New(int)()
	set1.Add(1,2,3,4,2,4)
	set2 := New(int)()
	set2.Add(3,4,5,6,7)
	fmt.Println("union")
	set1.Union(set2).Iterate(it)
	fmt.Println("intersection")
	set1.Intersection(set2).Iterate(it)
}*/

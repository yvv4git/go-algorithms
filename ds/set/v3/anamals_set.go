package main

// The animalSet type is a type alias of `map[string]struct{}`
type animalSet map[string]struct{}

// Adds an animal to the set
func (s animalSet) add(animal string) {
	s[animal] = struct{}{}
}

// Removes an animal from the set
func (s animalSet) remove(animal string) {
	delete(s, animal)
}

// Returns a boolean value describing if the animal exists in the set
func (s animalSet) has(animal string) bool {
	_, ok := s[animal]
	return ok
}

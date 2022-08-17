package main

import "fmt"

func main() {
	// Initializing zoo as a new set
	zoo := animalSet{}

	// Adding members to the set
	zoo.add("Walrus")
	zoo.add("Parrot")
	zoo.add("Lion")
	zoo.add("Giraffe")
	zoo.add("Bear")

	// Adding an existing member to the set again
	zoo.add("Lion")

	// Removing members from the set
	zoo.remove("Parrot")

	fmt.Println(zoo)
	// map[Bear:{} Giraffe:{} Lion:{} Walrus:{}]

	// Checking the presence of elements in a set
	fmt.Println(zoo.has("Penguin"))
	// false
	fmt.Println(zoo.has("Bear"))
	// true
}

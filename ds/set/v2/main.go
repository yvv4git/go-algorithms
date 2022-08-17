package main

import (
	"fmt"
)

func main() {
	// Zoo acts as a set of strings
	zoo := map[string]struct{}{}

	// We can add members to the set
	// by setting the value of each key to an
	// empty struct
	zoo["Walrus"] = struct{}{}
	zoo["Parrot"] = struct{}{}
	zoo["Lion"] = struct{}{}
	zoo["Giraffe"] = struct{}{}

	// Adding a new member to the set
	zoo["Bear"] = struct{}{}

	// Adding an existing to the set
	zoo["Lion"] = struct{}{}

	// Removing a member from the set
	delete(zoo, "Parrot")

	fmt.Println(zoo)
	// map[Bear:{} Giraffe:{} Lion:{} Parrot:{} Walrus:{}]

}

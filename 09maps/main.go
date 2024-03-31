package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exploring Maps!!")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["Py"] = "Python"
	languages["Go"] = "Golang"
	fmt.Println("List of all languages", languages)

	fmt.Println("Shortcut Py refers to: ", languages["Py"])

	delete(languages, "JS")
	fmt.Println("List of all languages", languages)

	//looping through a map
	for ind, val := range languages {
		fmt.Printf("%v is %v.\n", ind, val)
	}

	//Using Maps to Implement Sets

	//Although the Go language doesn’t have the set data structure in-built, we can use the map data structure
	//and implement the set using it. We know that set is just a collection of elements without the key-value pair
	//that we have in a map. The map can be represented with an empty struct value. Below is the implementation
	// of a set by using a map in the Go language.
	// Zoo acts as a set of strings
	zoo := map[string]struct{}{}

	// We can add members to the set
	// by setting the value of each key to an
	// empty struct
	zoo["elephant"] = struct{}{}
	zoo["tiger"] = struct{}{}
	zoo["owl"] = struct{}{}
	zoo["lion"] = struct{}{}

	// Adding a new member to the set
	zoo["Bear"] = struct{}{}

	// Adding an existing to the set
	zoo["lion"] = struct{}{}

	// Removing a member from the set
	delete(zoo, "owl")

	fmt.Println(zoo)

}

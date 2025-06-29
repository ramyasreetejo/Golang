package main

import "fmt"

/*
In C++, we use inheritance to reuse and override behavior. In Go, we don’t have inheritance.
Instead, we embed structs to compose behavior. It’s more flexible, reduces tight coupling, and encourages clean interfaces.

For example, Animal is embedded inside Dog and Cat, and override Speak() in each type.
We can also use interfaces like Speaker for polymorphism without needing a hierarchy.

Finally, for cross-cutting concerns like logging, Logger is embedded into unrelated types like ServiceA and ServiceB.
That way, we get shared behavior without inheritance.
*/

// ----------- 1. Base Behavior Struct -----------
type Animal struct{}

func (a Animal) Speak() {
	fmt.Println("I make a sound")
}

// ----------- 2. Composition Instead of Inheritance -----------

type Dog struct {
	Animal
}

func (d Dog) Speak() {
	fmt.Println("Bark!")
}

type Cat struct {
	Animal
}

func (c Cat) Speak() {
	fmt.Println("Meow!")
}

// ----------- 3. Interface for Polymorphism -----------

type Speaker interface {
	Speak()
}

func Announce(s Speaker) {
	s.Speak()
}

// ----------- 4. Shared Behavior Across Unrelated Types -----------

type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

type ServiceA struct {
	Logger
	Name string
}

type ServiceB struct {
	Logger
	ID int
}

// ----------- 5. Main Function to Run Everything -----------

func main() {
	fmt.Println("== Composition vs Inheritance in Go ==")

	// Inheritance-like behavior using composition
	d := Dog{}
	c := Cat{}

	fmt.Println("\n-- Direct Calls --")
	d.Speak()        // Bark!
	d.Animal.Speak() // I make a sound (embedded)
	c.Speak()        // Meow!
	c.Animal.Speak() // I make a sound

	fmt.Println("\n-- Interface Polymorphism --")
	var s Speaker
	s = d
	Announce(s) // Bark!
	s = c
	Announce(s) // Meow!

	fmt.Println("\n-- Reusing Logic via Composition (Logger) --")
	a := ServiceA{Name: "Analytics"}
	b := ServiceB{ID: 42}

	a.Log("Service A started")
	b.Log("Service B failed")
}

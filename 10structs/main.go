package main

import "fmt"

func main() {
	fmt.Println("Exploring Structs in golang!!")

	User1 := User{"Ramya", "t2001ramyasree@gmail.com", true, 22}
	fmt.Println(User1)
	fmt.Printf("The User1 details are: %v \n", User1)
	fmt.Printf("The User1 details in more detail are: %+v \n", User1)
	fmt.Printf("The User1 name is %v and age is %v.\n", User1.Name, User1.Age)

}

// User starts w caps and so do Name Email etc becuase they are public and can be exported out of this package also.
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

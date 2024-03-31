package main

import "fmt"

func main() {
	fmt.Println("Exploring methods in golang!!")

	User1 := User{"Ramya", "t2001ramyasree@gmail.com", true, 22}
	fmt.Println(User1)
	//fmt.Printf("The User1 details are: %v \n", User1)
	//fmt.Printf("The User1 details in more detail are: %+v \n", User1)
	//fmt.Printf("The User1 name is %v and age is %v.\n", User1.Name, User1.Age)
	User1.FindStatus()
	User1.ChangeEmail1("ramya@go.dev")
	fmt.Println(User1)
	User1.ChangeEmail2("ramya@go.dev")
	fmt.Println(User1)
}

// User starts w caps and so do Name Email etc becuase they are public and can be exported out of this package also.
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) FindStatus() {
	fmt.Println("The activity status of user is: ", u.Status)
}

func (u User) ChangeEmail1(email string) {
	u.Email = email
	fmt.Println("The email of user is: ", u.Email)
}

func (u *User) ChangeEmail2(email string) {
	u.Email = email
	fmt.Println("The email of user is: ", u.Email)
}

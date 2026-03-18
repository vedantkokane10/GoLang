package main

import "fmt"

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

// methods -> methods are function which are part of structs

// this is now a method associated to User struct
func (u User) getStatus(){
	fmt.Println(u.Status)
}

// we start with -> func (argumentName structName) functionName() {
//	statements
// }

// func (u User) getStatus() here copy of u (user) is passed not the original u (user) 

// passing pointer so that changes are saved (reference of actual u is passed) 
func (u *User) changeMail(){
	u.Email = "newMail@gmail.com"
}

func main(){
	user := User{"Vedant", "ved@gmail.com", true, 23}
	user.getStatus()
	user.changeMail()
	fmt.Println(user.Email)
}


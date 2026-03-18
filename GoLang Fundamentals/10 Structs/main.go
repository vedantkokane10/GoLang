package main

import "fmt"

// Go isn’t an object-oriented (OO) language
// no inheritance, no super() or parent()

type User struct{
	Name string
	Email string
	Status bool
	Age int
	Father *User

}

// We are using pointers coz in golang everything is pass by value, so to avoid copy of Father type we can use pointer who point to the father 

func main(){
	raj := User{Name: "raj", Email: "raj@gmail.com", Status: true, Age: 23, Father: nil};
	fmt.Println(raj)
	
	// can skip the field name and rely on the order of the field declarations
	vedant := User{"Vedant", "vedant@gmail.com", true, 23, &raj};
	fmt.Println(vedant)
	fmt.Println(vedant.Age)

	fmt.Printf("Vedant's age is %+v \n", vedant.Age)

	fmt.Printf("Vedant's age is %v and email is %v \n", vedant.Age, vedant.Email)

	yash := User{Name:"Yash"}
	yash.Email = "yash@gmail.com"


	incremntAge(&raj)
	fmt.Println(raj)

	raj.changeStatus()

	fmt.Println(raj)
}

func incremntAge(user *User){
	user.Age++
}

// we can associate a method with a structure:
func (user *User) changeStatus() {
	user.Status = !user.Status
}

// Constructors
// Structures don’t have constructors. Instead, you create a function that returns an instance of the desired type (like a factory)
func NewUser(name string, email string, status bool, age int) *User{
	return &User{
		Name:name,
		Email: email,
		Status: status,
		Age: age,
	}
}

// Also valid
/* Our factory doesn’t have to return a pointer; this is absolutely valid
func NewUser(name string, email string, status bool, age int) User {
	return &User{
		Name:name,
		Email: email,
		Status: status,
		Age: age,
	}
}
*/

// Despite the lack of constructors, Go does have a built-in new function which is used to allocate the memory required by a type.
// The result of new(X) is the same as &X{}
// ram := new(User) is same as ram := &User{}

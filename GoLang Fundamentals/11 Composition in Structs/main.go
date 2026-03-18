package main

import "fmt"

// Go supports composition, which is the act of including one structure into another. In some languages, this is called a trait or a mixin.

type Person struct{
	Name string
}

func (p *Person) introduce() {
	fmt.Printf("Hi, I am %v", p.Name);
}

type Saiyan struct{
	*Person // pointer so this Saiyan struct is storing address which is pointing to the Person struct
	Power int
}



func main(){
	goku := Saiyan{
		Person: &Person{"Goku"}, // goku variable (of type type Saiyan) stores address where Person (type Person) with name Goku is stored
		Power: 500,
	}

	goku.introduce()

	// we didn’t give it an explicit field name, we can implicitly access the fields and functions of the composed type. However, the Go compiler did give it a field name
	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)


	
	
}
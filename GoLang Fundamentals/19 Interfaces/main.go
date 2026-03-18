package main

import "fmt"

// An interface is a contract. OR
// Interfaces are types that define a contract but not an implementation.
// Interfaces help decouple our code from specific implementations

type Speaker interface{
	Speak() string
}

// Any type that has: Speak() string automatically becomes a Speaker.

// Important Go thing (very important) In Java/C++, you often explicitly say: class Dog implements Speaker
// But in Go, you do NOT write that. Go uses implicit implementation.
// That means:
// If the methods match, Go automatically treats it as implementing the interface.
// This is one of Go’s coolest features.


type Dog struct{

}


func (d Dog) Speak() string{
	return "Woof"
}
// Dog now implements the Speaker interface

type Cat struct{

}

func (c Cat) Speak() string{
	return "Meow"
}
// Cat now implements the Speaker interface


func PrintSound(s Speaker){
	fmt.Println(s.Speak())
}

func main(){
	d := Dog{}
	c := Cat{}

	PrintSound(d)
	PrintSound(c)
}

/*
✅ Why interfaces are useful?
Without interfaces:
	func PrintDog(d Dog)
	func PrintCat(c Cat)

Too rigid ❌
With interface:
	func PrintSound(s Speaker)
Flexible ✅

So interfaces help with:
	abstraction
	reusability
	loose coupling
	clean design
	writing generic behavior

// Key rule in Go : A type satisfies an interface if it implements all methods in that interface.

Example:
type Shape interface {
	Area() float64
}
Then any type with: Area() float64
satisfies Shape.

Interface is type-based on behavior, not data
	In OOP languages, you often think:
	- inheritance
	- parent-child hierarchy

In Go, think: “What can this thing do?”
Not: “What class does it belong to?”

So Go prefers behavior over inheritance.


🔥 This is the Go mindset
	Instead of:
		Animal -> Dog -> Labrador

	Go likes:
		Speaker
		Runner
		Swimmer

		Because:
		a dog can speak
		a dog can run
		a fish can swim
		a robot can speak too

	So behavior matters more than hierarchy.

- Any type is a implements empty interface interface{} This means: “Accept any type.”
	Because every type satisfies an interface with zero methods.
	Ex. 
	func PrintAnything(x interface{}) {
		fmt.Println(x)
	}
	x can be a int, string, struct, slice, anything

- In modern Go, use any, any is just an alias for interface{}
	func PrintAnything(x any) {
		fmt.Println(x)
	}

- Type assertion (getting actual type back)
	If you have an interface variable and want the real value:
	var x any = "Vedant"
	s, ok := x.(string)
	if ok {
		fmt.Println("String:", s)
	} else {
		fmt.Println("Not a string")
	}

- ✅ When should you use interfaces?
	Use interface when:
	You want function to work with multiple types
	You care about behavior, not exact struct
	You want testability (mocking)
	You want loose coupling between packages

- Pointer receiver vs value receiver (VERY important in Go)
	- type Speaker interface {
			Speak()
	  }

	  type Dog struct{}
	  
	  // Value receiver:
	  func (d Dog) Speak() {}

	  Then both work:
		Dog{}
		&Dog{}

	  // Pointer receiver:
	  func (d *Dog) Speak() {}
	  
	  Then only this works:
		&Dog{} ✅
	  But this won’t satisfy interface:
		Dog{} ❌

	Example - 
		type Speaker interface {
			Speak()
		}

		type Dog struct{}

		func (d *Dog) Speak() {
			fmt.Println("Woof")
		}

		func main() {
			var s Speaker

			// s = Dog{}   // ERROR
			s = &Dog{}     // OK

			s.Speak()
		}
*/
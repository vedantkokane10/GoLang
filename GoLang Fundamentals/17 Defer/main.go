package main

import "fmt"

// The defer keyword in Go is used to schedule a function call to be executed just before the surrounding function returns. 
// This ensures that the deferred function runs regardless of how the surrounding function exits, whether through normal completion, a return statement, or a runtime panic.
// Its used in 
// Resource Cleanup: It is commonly used to ensure that resources like open files, network connections, and database handles are closed, preventing resource leaks.
// Unlocking Mutexes: In concurrent programming, defer ensures that a mutex is unlocked, preventing deadlocks.
// Timing Execution: defer can be used with an anonymous function to measure how long a function takes to execute, logging the duration at the end.
// Panic Handling: A deferred function is the only place where the built-in recover function can be called to catch a panic and resume normal execution, preventing the program from crashing

func main(){
	// example1()

	// example2()

	// example3()

	example4()
}

func example1 (){
	defer fmt.Println("World")
	fmt.Println("Hello")

	// output will be first hello and the at last World
	// coz defer cuts the line no 6 and puts it at the end before closing curly brace
}


func example2 (){
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("third")

	// Last in first out
	// output will be :
	// third
	// two
	// one
}

func example3(){
	for i:=1;i<=5;i++{
		defer fmt.Println(i)
	}

	// an internal stack will be used (in background, the classic function call stack we learnt in recusion) and it will store elements one by one on top:
	// [1]
	// [2,1]
	// [3,2,1]
	// [4,3,2,1]
	// [5,4,3,2,1]

}

func example4(){
	defer fmt.Println("World")
	fmt.Println("Hello")

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("third")

	for i:=1;i<=5;i++{
		defer fmt.Println(i)
	}

	/* output will be
		Hello
		5
		4
		3
		2
		1
		third
		two
		one
		World
	*/
}
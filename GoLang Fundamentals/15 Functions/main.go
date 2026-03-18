package main

import "fmt"

func f1(){
	fmt.Println("inside f1");
}

func greeter(name string){
	fmt.Println("hello ", name)
}

func adder(a int, b int) int {
	return a+b
}

// lets i wanna add any n numbers
func proAdder(values ...int) (int, string) {
	// values is now a slice
	sum := 0

	for _,val := range values{
		sum += val
	}

	return sum, "sum performed successfully"
}

// nested functions are not supported

func main(){
	f1();
	greeter("vedant")
	fmt.Println(adder(1,3))
	sum, message := proAdder(1,3,5,6,7)
	fmt.Println(sum, message)
}
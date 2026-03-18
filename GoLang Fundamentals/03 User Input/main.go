package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var welcome string = "welcome user"
	fmt.Println(welcome)


	reader := bufio.NewReader(os.Stdin)
	// bufio is a package
	// NewReader() is a public function which takes argument rd -> reader which means from where to read
	// here we passed rd as Operating system's standard input (os.Stdin)
	// the NewReader() returns a pointer 



	fmt.Println("Enter name: ")
	// input, err := reader.ReadString('\n')  <= this is the standard way of writing comma ok and handling error 
	input, _ := reader.ReadString('\n') // if we don't want to use err we can use '_' instead, we can use  '_' anywhere as it helps if we don't want to use that variable anywhere
	// reader.ReadString('\n') -> it tells the reader to read string until you encounter '/n' which is enter button
	// ReadString() function returns 2 values first is input data and second is error


	fmt.Println("Welcome",input)
}

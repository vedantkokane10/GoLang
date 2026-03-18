package main

import "fmt"

func main(){
	cnt := 5

	if cnt > 5 {
		fmt.Println("Greater than 5")
	} else if cnt < 5{
		fmt.Println("less than 5")
	} else{
		fmt.Println("Equal to 5")
	}
	// opening curly braces should come on same like as (if, else statements)


	// we can declare a variable before checking the condition assiociated with it within the if-statement
	if c:=11; c == 11 {
		fmt.Println("Equal to 11")
	} else if c < 11 {
		fmt.Println("less than 11")
	} else{
		fmt.Println("Greater than 1")
	}


}
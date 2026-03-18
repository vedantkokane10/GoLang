package main

import (
	"fmt"
	"math/rand"
)

func main() {
	diceNumber := rand.Intn(6) + 1


	// in go we have automatic break -> once a case is satsified the switch breaks
	switch diceNumber{
		case 1:
			fmt.Println(1)
		case 2:
			fmt.Println(2)
		case 3:
			fmt.Println(3)
		case 4:
			fmt.Println(4)
		case 5:
			fmt.Println(5)
		case 6:
			fmt.Println(6)
		default:
			fmt.Println("Wrong dice number")
	}


	// if we want that even if a case is true we shall check for some other cases and also execute them, then we can use fallthrough

	switch diceNumber{
		case 1:
			fmt.Println(1)
			fallthrough
		case 2:
			fmt.Println(2)
			fallthrough
		case 3:
			fmt.Println(3)
			fallthrough
		case 4:
			fmt.Println(4)
			fallthrough
		case 5:
			fmt.Println(5)
			fallthrough
		case 6:
			fmt.Println(6)
			fallthrough
		default:
			fmt.Println("Wrong dice number") 
	}

	/*
		3 -> random number generated
		3 -> second switch (3 first)
		4 -> (fallthrough from 3)
		5 -> (fallthrough from 4)
		6 -> (fallthrough from 5)
		Wrong dice number  -> (fallthrough from 6)
	
	*/
}

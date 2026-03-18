package main

import "fmt"

func main(){

	// syntax
	for i:=0; i<10; i++{
		fmt.Println(i)
	}

	days := []string{"monday", "tuesday", "wednesday"}
	// iterating over an slice using range
	for i := range days{
		fmt.Println(days[i])
	}

	// other way 
	for index, day := range days{
		fmt.Println(index , " ", day)
	}

	u := 0
	// in golang we don't have while loops, but we can use for-loops like for  while-loops purpose
	for u < 10{
		fmt.Println(u)
		if u == 5{
			// continue
		}
		if u == 8{
			break
		}
		u++
	} 
}
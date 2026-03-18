package main

import "fmt"

// pointer is reference to actual memory address

func main(){
	var one int = 1;
	var ptr *int = &one;
	fmt.Println(ptr);

	var ptr2 *string;
	fmt.Println(ptr2); // by default the value of any pointer is nil

	// declaring pointer with valorous operator
	num := 5
	ptr3 := &num
	fmt.Println("num = ", num, " ptr3 = ", ptr3)
	
	// incrementing num via ptr3
	*ptr3 += 1;
	fmt.Println("*ptr3 = ", *ptr3, " num = ", num)

	// multipling by 2 num via ptr3
	*ptr3 *= 2;
	fmt.Println("*ptr3 = ", *ptr3, " num = ", num)

}

// Won't focus much on this, since I know C++ very well
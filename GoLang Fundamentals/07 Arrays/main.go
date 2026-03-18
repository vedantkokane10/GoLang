package main

import "fmt"

// Note: as per docs arrays are not used for majority of the tasks unlike other langauges
// instead GO has slices which are frequently used

func main(){
	var arr [5]int
	fmt.Println(arr); // by default all elements are 0


	// adding elements
	arr[0] = 1;
	arr[4] = 2;
	arr[2] = 3;
	fmt.Println(arr);


	fmt.Println(len(arr)); // to get lemgth

	// initialized array 
	var fruits = [6]string {"apple", "banana", "orange", "kiwi"}
	fmt.Println(fruits);

	// iterating via for loops
	for i:=0;i<len(fruits);i++{
		fmt.Println(fruits[i]);
	}
}
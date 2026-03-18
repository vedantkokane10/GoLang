package main

import "fmt"

func main(){
	// To create an empty map, we use the builtin `make`:
	// `make(map[key-type]val-type)`.
	mp := make(map[string]int)

	mp["one"] = 1;
	mp["two"] = 2;
	mp["three"] = 3;

	fmt.Println(mp) // map[one:1 three:3 two:2]
	fmt.Println(mp["one"]) // 1


	delete(mp, "one") // for deleting a key
	fmt.Println(mp)

	// iterating over an map
	for key,value := range mp{
		fmt.Println(key, value)
	}
	// The range keyword in Go is used within a for loop to iterate over elements in various data structures, including arrays, slices, maps, strings, and channels. It typically returns one or two values in each iteration.


	// to check whether a element present or not
	_, isPresent := mp["three"] // it returns true/false
	fmt.Println(isPresent) // true 

	fmt.Println(mp["notFoundKey"]) // if a key not present and if we call it using [] then it returns 0


	// we can also declare and initialize a new map in the same line with this syntax.
	mp2 := map[string]int{"foo":1, "bar":2}
	fmt.Println(mp2)

	// no of keys inside a map
	fmt.Println(len(mp2))

	// Maps grow dynamically. However, we can supply a second argument to make to set an initial size:
	lookup := make(map[string]int, 100)
	fmt.Println(len(lookup))
	// If we have some idea of how many keys your map will have, defining an initial size can help with performance.

	
}
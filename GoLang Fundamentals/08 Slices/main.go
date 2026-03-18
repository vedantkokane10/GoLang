package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// A slice is a lightweight structure that wraps and represents a portion of an array
// we use slices when we don't know the no of elements that the data structure will store, if we know how many elements then we can use array.

func main() {

	// multiple ways to declare

	var fruits = []string{"apple", "banana"}
	fmt.Println(fruits)

	// adding new elements
	fruits = append(fruits, "mango", "orange") // append(slice, element/elements) -> add mango, orange inside fruits and store the new slice inside the existing fruits slice 
	fmt.Println(fruits)

	fmt.Println(fruits[0]) // same like array

	// slicing inside a slice :)
	fruits = append(fruits[1:]) // only keep elements from 1 to end (start:end)
	fmt.Println(fruits)
	// last range is not inclusive for ex. fruits[1:3] => 1,2 only and not 3 
	fruits = append(fruits[0:2]) 
	fmt.Println(fruits)


	// dynamic slice
	highScores := make([]int, 5); // make(type. size)
	// The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it.
	// The specification of the result depends on the type
	fmt.Println(highScores)

	highScores[0] = 2;
	highScores[1] = 5;
	highScores[2] = 4;
	highScores[3] = 3;
	highScores[4] = 55;
	fmt.Println(highScores)


	// highScores[5] = 9; // this will give error
	highScores = append(highScores, 9, 23) // at this step the memory allocation happens again
	fmt.Println(highScores)

	// sorting a slice
	sort.Ints(highScores); // use of sort package
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// Package sort provides primitives for sorting slices and user-defined collections.


	// removing a element from slice
	var courses = []string {"react", "node", "go", "ruby", "python", "java"}
	fmt.Println(courses);
	var index int = 3; // index to remove
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses);

	// ... is  is the variadic parameter unpacker (or "splat" operator)
	// It takes the slice resulting from course[index+1:] and unpacks it into individual arguments to be passed to the append function. 
	// Breakdown of the Expression
	// course[:index]: Creates a slice of all elements before the item to be removed.
	// course[index+1:]: Creates a slice of all elements after the item to be removed.
	// ...: Tells append to take every element from the second slice (course[index+1:]) and add them one by one to the first slice. 


	scoresTest := make([]int, 0, 10)
	// make(sliceType, currentSize, maxCapacity)
	// scoresTest[7] = 55;  -> gives error index out of range [0] with length 0
	// scoresTest = append(scores, 5); // -> [] -> [5]
	fmt.Println(scoresTest)  // -> []

	scoresTest2 := make([]int, 5) // -> [0,0,0,0,0]
	scoresTest2 = append(scoresTest2, 555) // -> [0,0,0,0,0,555]


	// Manually removing an element using slicing
	arr3 := []int{1,2,3,4,5}
	removeElementAtIndex(arr3, 2);
	fmt.Println(arr3)

	scoresTest3 := make([]int, 10)
	for i:=0;i<10;i++{
		scoresTest3[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scoresTest3)

	worst := make([]int, 5)
	copy(worst, scoresTest3[:5])
	fmt.Println(scoresTest3)
	fmt.Println(worst)

	worst2 := make([]int, 5)
	copy(worst2[2:4], scoresTest3[:5])
	fmt.Println(worst2)
}


// it swaps the last value and the value we want to remove
func removeElementAtIndex(arr []int, index int) []int{
	lastIndex := len(arr)-1

	arr[index], arr[lastIndex] = arr[lastIndex], arr[index]

	return arr[:lastIndex]
}

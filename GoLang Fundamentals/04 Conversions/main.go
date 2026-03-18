package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conversions")

	reader := bufio.NewReader(os.Stdin)


	fmt.Println("Enter a number")
	inputNumber, _ := reader.ReadString('\n') // suppose i enter 33, it will read it as "33\n"

	fmt.Println("Enter a floating number")
	inputFloat, _ := reader.ReadString('\n')

	// converting string -> float
	numberFloat, err := strconv.ParseFloat(strings.TrimSpace(inputFloat), 64) 
	// strings.TrimSpace() function we remove spaces from input. Ex "33\n" -> "33"
	// strconv is package used to convert string values into desired datatype values
	// ParseFloat() takes 2 parameters 1) the string data and 2) the bit size (32/64) (float32 / float64)

	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println("number - ", numberFloat)
		fmt.Printf("type - %T \n", numberFloat)
	}
	


	// converting string -> int
	numberInt, err := strconv.ParseInt(strings.TrimSpace(inputNumber), 10, 64) 
	// ParseInt(inputString, base, bitSize) for this ParseInt(inputString, base (10 - decimal), bitsize (int64 / int32))

	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println("number - ", numberInt)
		fmt.Printf("type - %T \n", numberInt)
	}




	// int -> string
	intToString := strconv.FormatInt(numberInt, 10)
	fmt.Println("number - ", intToString)
	fmt.Printf("type - %T \n", intToString)
	

	// float -> string
	floatToString := strconv.FormatFloat(numberFloat, 'f', 2, 64)
	fmt.Println("number - ", floatToString)
	fmt.Printf("type - %T \n", floatToString)

}

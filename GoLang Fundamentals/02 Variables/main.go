package main

import "fmt"

var globalVariable int = 33;
const LoginToken string = "Test_Login"; // LoginToken is public (since starts with a capital letter) and global varibale which is constant i.e cannot be updated

func main() {
	var userName string = "Vedant";
	fmt.Println(userName);
	fmt.Printf("Variable userName is of type: %T \n", userName); // %T give the type (datatype) of variable

	var isVerified bool = true;
	fmt.Println(isVerified);
	fmt.Printf("Variable isVerified is of type: %T \n", isVerified);

	var smallIntValue int8 = 127; // if i use a value greater than 127 then go will give me errors
	fmt.Println(smallIntValue);
	fmt.Printf("Variable smallIntValue is of type: %T \n", smallIntValue);

	var smallFloatValue float32 = 255.93838338838383838;  // it will clip the value to 255.93839 
	fmt.Println(smallFloatValue);
	fmt.Printf("Variable smallFloatValue is of type: %T \n", smallFloatValue);

	var largeFloatValue float64 = 255.93838338838383838;  // it will clip the value to 255.93838338838384, its more than what float32 offered
	fmt.Println(largeFloatValue);
	fmt.Printf("Variable largeFloatValue is of type: %T \n", largeFloatValue);


	// Different ways to declare variables

	// Implicit style
	var numberOfUsers = 2929; // go will automatically assign appropriate datatype based on value, once a variable's datatype is assigned then it cannot be changed to any other datatype
	fmt.Println(numberOfUsers);
	fmt.Printf("Variable numberOfUsers is of type: %T \n", numberOfUsers);
	// numberOfUsers = "323" this will give error


	// No var style
	testVariable := 123; // := is valorous operator is used when we don't use var, := performs both declaration (var) and assignment (=) in one step.
	fmt.Println(testVariable);
	fmt.Printf("Variable testVariable is of type: %T \n", testVariable);
	// IMP -> this kind of style is only allowed inside a method (function) if we do this globally then we will get errors
	// this is known as short assignment
	// Mixed Declaration: we cannot use var and := together. 

	// default values
	var intDefault int; // default value here is 0 
	fmt.Println(intDefault);
	fmt.Printf("Variable intDefault is of type: %T \n", intDefault);

	fmt.Println(LoginToken);
	fmt.Printf("Variable LoginToken is of type: %T \n", LoginToken);

}

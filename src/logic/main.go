package main

import "fmt"

func main() {
	// && logical and - both need to be true ( NOTE: two wrongs don't make a right!)
	// || logical or  - one of them need to be true
	//  ! logical not - unary operator, used before a single operand, returns the inverse value of the operand

	var positive, negative bool = true, false

	fmt.Println("Assigned values:\tpositive=", positive, "\tnegative=", negative)

	result := (positive && negative)
	fmt.Println("AND logic:\tpositive && negative =\t", result)

	result = (positive || negative)
	fmt.Println("OR logic:\tpositive || negative =\t", result)

	result = !positive
	fmt.Println("NOT logic:\t!positive=\t", result)
}

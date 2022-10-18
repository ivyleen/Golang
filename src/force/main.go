package main

import "fmt"

func main() {
	// Operator precedence defines the order in which the Go compiler evaluates expressions.

	// Operator precedence in decreasing order with each row:
	// * / % << >> &
	// + - | ^
	// == != < <= >=
	// &&
	// ||

	sum := 2*3 + 4 - 5
	fmt.Println("Default order: %v", sum)

	sum = 2 * ((3 + 4) - 5)
	fmt.Println("Forced order: %v", sum)

	sum = 7 % 3 * 2
	fmt.Println("Default order: %v", sum)

	sum = 7 % (3 * 2)
	fmt.Println("Forced order: %v", sum)
}

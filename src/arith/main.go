package main

import "fmt"

func main() {
	// Operators: +, -, *, /, %(reminder), ++, --.
	// Best practices: use parentheses when having multiple arithmetics.

	a := 10
	b := 5

	fmt.Println("Addition:\t", (a + b))
	fmt.Println("Substraction:\t", (a - b))
	fmt.Println("Multiplication:\t", (a * b))
	fmt.Println("Division:\t", (a / b))
	fmt.Println("Reminder:\t", (a % b))

	a++
	fmt.Println("Increment:\t", a)
	b--
	fmt.Println("Decrement:\t", b)
}

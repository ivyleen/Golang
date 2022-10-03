package main

import "fmt"

func main() {
	// Naming rules:
	// NO keywords, arithmetic operators, punctuation characters or spaces
	// MUST start with letter, but can contain digits after as well as mixed case

	// Data types: string, int, float64, byte, bool

	// Fmt specifiers
	// %s - string
	// %d - integer
	// %f - float
	// %c - a character
	// %t - bool
	// %p - memory address
	// %v - any of the above
	// %T - data type

	num := 100
	pi := 3.1415926536

	fmt.Printf("num: %v type %T \n", num, num)
	fmt.Printf("pi: %v type %T \n\n", pi, pi)

	fmt.Printf("%%7d displays %7d \n", num)
	fmt.Printf("%%07d displays %07d \n\n", num)

	fmt.Printf("Pi is approximately %1.10f \n", pi)
	fmt.Printf("Right-aligned %20.3f rounded pi \n", pi)
	fmt.Printf("Left-aligned %-20.4f rounded pi", pi)
}

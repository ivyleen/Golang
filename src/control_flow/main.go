package main

import "fmt"

func main() {
	// Conditional branching - fundamental programming technique that
	// offers the program two directions in which to proceed, depending on the results of the evaluation.

	// The else keyword and opening { curly bracket must appear on the same line as the } closing curly bracket of the previous block of statements!

	// Single statement condition.
	if 'A' == 'A' {
		fmt.Println("\nCharacters match.")
	}

	// Multiple statements condition.
	if 5 == 5 && 'A' == 'A' {
		fmt.Println("\nBoth expressions are true.")
	}

	// Nested multiple statements condition.
	if 5 > 1 {
		if 7 > 2 {
			fmt.Println("\nBoth expressions are true.")
		}
	}

	// If-else condition.
	if 5 < 1 {
		fmt.Println("\n1st condition is true.")
	} else if 'A' != 'A' {
		fmt.Println("\n2st condition is true.")
	} else {
		fmt.Println("\nBoth conditions are false.")
	}
}

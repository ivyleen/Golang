package main

import "fmt"

func main() {
	// Conditional branching - fundamental programming technique that
	// offers the program two directions in which to proceed, depending on the results of the evaluation.

	/////////////////////// If - else statements ///////////////////////

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

	/////////////////////// Switch cases ///////////////////////
	num := 2

	switch num {
	case 1:
		fmt.Println("\nNumber is One.")
	case 2:
		fmt.Println("\nNumber is Two.")
	case 3:
		fmt.Println("\nNumber is Three.")
	default:
		fmt.Println("\nNumber is unrecognized.")
	}

	// A switch with number of march-values are to each execute the same statements,
	// only the final case statement need include the statements to be executed.

	switch num {
	case 0:
	case 1:
	case 2:
		fmt.Println("\nLess than three.")
	case 3:
		fmt.Println("\nExactly three.")
	default:
		fmt.Println("\nGreater than 3 or negative.")
	}

	// When a match is found, the switch normally ends, but you can add the fallthrough keyword
	// as a final statement in any case block to continue to the next case statement.

	ch := 'a'

	switch ch {
	case 'a':
		fmt.Printf("\nAlthough the letter is lower case 'a',")
		fallthrough
	case 'A':
		fmt.Println("\nThe letter is the first in the alphabet.")
	default:
		fmt.Println("\nLetter is unrecognized.")
	}
}
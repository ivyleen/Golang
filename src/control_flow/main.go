package main

import "fmt"

func main() {
	// Conditional branching - fundamental programming technique that
	// offers the program two directions in which to proceed, depending on the results of the evaluation.

	IfElseExample()

	num := 5

	fmt.Println("Passing by value num in main func before SwitchCaseExample: ", num)

	SwitchCaseExample(num, 'b')

	LoopExample()

	fmt.Println("Confirming that num in main func hasn't changed before LabelExample: ", num)

	LabelExample(num)
}

// Functions: the functions below are functions that don't return anything as a value.
// If a function returns a value than the function will look more or less like this:
// func <name of the function>() <return type> {}

func IfElseExample() {
	fmt.Println("If - else statements")

	// The else keyword and opening { curly bracket must appear on the same line as the } closing curly bracket of the previous block of statements!

	// Single statement condition.
	if 'A' == 'A' {
		fmt.Println("Characters match.")
	}

	// Multiple statements condition.
	if 5 == 5 && 'A' == 'A' {
		fmt.Println("Both expressions are true.")
	}

	// Nested multiple statements condition.
	if 5 > 1 {
		if 7 > 2 {
			fmt.Println("Both expressions are true.")
		}
	}

	// If-else condition.
	if 5 < 1 {
		fmt.Println("1st condition is true.")
	} else if 'A' != 'A' {
		fmt.Println("2st condition is true.")
	} else {
		fmt.Println("Both conditions are false.")
	}
}

func SwitchCaseExample( num int, ch byte) {
	fmt.Println("\nSwitch cases")

	switch num {
	case 1:
		fmt.Println("Number is One.")
	case 2:
		fmt.Println("Number is Two.")
	case 3:
		fmt.Println("Number is Three.")
	default:
		fmt.Println("Number is unrecognized.")
	}

	// A switch with number of march-values are to each execute the same statements,
	// only the final case statement need include the statements to be executed.

	switch num {
	case 0:
	case 1:
	case 2:
		fmt.Println("Less than three.")
	case 3:
		fmt.Println("Exactly three.")
	default:
		fmt.Println("Greater than 3 or negative.")
	}

	// With characters is a bit different looking switch case:

	switch ch {
	case 'b', 'B':
		fmt.Println("Second letter in the alphabet.")
	default:
		fmt.Println("Not the second letter in the alphabet.")
	}

	// When a match is found, the switch normally ends, but you can add the fallthrough keyword
	// as a final statement in any case block to continue to the next case statement.

	ch = 'a'

	switch ch {
	case 'a':
		fmt.Printf("Although the letter is lower case 'a',")
		fallthrough
	case 'A':
		fmt.Println("The letter is the first in the alphabet.")
	default:
		fmt.Println("Letter is unrecognized.")
	}

	num -= 2
	fmt.Println("Num changed inside of SwitchCaseExample: ", num)
}

func LoopExample() {
	fmt.Println("\nLoops")

	// Loop is a piece of code in a program that automatically repeats.
	// One complete execution of all statements within a loop is called "iteration", or a "pass".
	// The length of the loop is controlled by a condition in the loop. While the tested expression is true, the loop continues.
	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Loop iteration: ", counter)
	}

	// The iterator of the most outer loop is usually named i, than nested iterator is j, and than k.
	for i := 1; i <= 5; i++ {
		fmt.Println("Outer loop iteration: ", i)

		for j := 1; j <= 5; j++ {
			fmt.Println("\tInner loop iteration: ", j)
		}
	}

	fmt.Println("\nWhile loops")

	// Loop while true = while loop in other languages.
	counter := 1 // already initialized variable
	for counter <= 5 {
		fmt.Println("While loop iteration: ", counter)
		counter++ // incrementing the variable until the for test-expression becomes false
	}

	// Infinite loop with a test expression contained in the loop body = do-while in other languages
	i := 10
	for {
		fmt.Println("\tCountdown: ", i)
		i--

		if i < 1 {
			fmt.Println("\t\tLift off!")
			break
		}
	}

	fmt.Println("\nBreak and continue")

	// Break and continue in loops:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 3 && j == 2 {
				fmt.Println("\tContinue when i =", i, "j = ", j)
			}

			if i == 2 && j == 2 {
				fmt.Println("\tBreaking when i =", i, "j = ", j)
			}

			fmt.Println("Running i =", i, "j = ", j)
		}
	}
}

func LabelExample(max int) {
	fmt.Println("\nGo to labels")

	// goto keyword allows the program flow to jump to a label at any position in the program code much like a hyperlink on a web page.
	// This can cause a lot of trouble so it's generally considered bad programming practice.

	// One example of valid usage:

	for i := 1; i <= max; i++ {
		for j := 1; j <= max; j++ {
			if i == 2 && j == 2 {
				goto end
			}

			fmt.Println("Running i =", i, "j = ", j)
		}
	}
end:
}


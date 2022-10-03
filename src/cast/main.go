package main

import (
	"fmt"
	"strconv"
)

func main() {
	/* Casting between the character types: */
	var char byte = 'A' // byte is an alias for uint8
	fmt.Printf("\nchar: {value - %v} {type - %T} {casting to string - %v} \n", char, char, string(char))

	/* Casting between character type and an integer. */
	var str string = "42"

	// NOTE: Functions can have multiple return values.
	// Function Atoi converts any string to an integer but will also give indication(return error message) if the conversion fails and nill if succeeds.
	num, err := strconv.Atoi(str)
	fmt.Printf("int from string: {value - %v} {type - %T} {error - %v}\n", num, num, err)

	// strconv.ParseFloat(): string -> float64 (will return the value + error message if it fails and nill if succeeds)
	fl, errMsg := strconv.ParseFloat(str, 64) // the 64 is indicator for the float size
	fmt.Printf("float from string:{value - %v} {type - %T} {error - %v} \n", fl, fl, errMsg)

	/* Casting between the integers. */
	decimal := float64(num)
	fmt.Printf("float from int: {value - %.2f} {type - %T} \n", decimal, decimal)

	/* Casting from integer to string. */
	intStr := strconv.Itoa(num)
	fmt.Printf("string from int: {value - %s} {type - %T} \n", intStr, intStr)
}

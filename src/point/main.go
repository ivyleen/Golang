package main

import "fmt"

func main() {
	// Golang supports pointers.

	// An initialized variable consists of three parts: name, value and address.
	// Example: name: num, value: 20, address: 0x00009e548
	num := 50

	// Declaration of pointer variable that will point to an int:
	var ptr *int

	// Pointers' value is an address.
	// For example we can assign the address of num as a value to our pointer variable:
	ptr = &num

	fmt.Printf("num value %v of type %T. \n", num, num)
	fmt.Printf("num address %v of type %T. \n\n", ptr, ptr)

	// If we want to see the value of that address we need to dereference the point variable:
	fmt.Printf("num via pointer: %v type %T. \n", *ptr, *ptr)
	fmt.Printf("ptr address: %v type %T \n\n", &ptr, &ptr)

	// We can change the value stored in the integer variable - by assignment to the pointer variable:
	*ptr = 1000
	fmt.Printf("num new value: %v type %T.", num, num)
}

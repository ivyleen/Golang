package main

import "fmt"

func main() {
	// Comparison operators are also known as "relational operators".

	var zero, num, max int = 0, 0, 1
	var up, dn byte = 'A', 'a'

	fmt.Println("Assigned values:\tzero=", zero, "\tnum=", num, "\tmax=", max, "\tup=", up, "\tdn=", dn)

	result := (num == zero) // 0 == 0
	fmt.Println("Equality:\tnum == zero\t", result)

	result = (up == dn) // A == a
	fmt.Println("Equality:\tup == dn\t", result)

	result = (max != zero) // 1 != 0
	fmt.Println("Inequality:\tmax != zero\t", result)

	result = (zero > max) // 0 > 1
	fmt.Println("Greater:\tzero > max\t", result)

	result = (max <= zero) // 1 <= 0
	fmt.Println("Less or equal:\tmax <= zero\t", result)
}

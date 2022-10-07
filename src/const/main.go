package main

import "fmt"

func main() {
	// Constant declarations must be initialized with its fixed value in the declaration:
	const pi = 3.14159

	const oneMillion, oneThoousand = 1000000, "1000"

	// iota is a keyword that will create a numerical sequence, beggining from iota(0), that increments by one for each named constant in the list.
	const (
		Monday    = iota + 1 // 1
		Tuesday              // 2
		Wednesday            // 3
		Thursday             // 4
		Friday               // 5
		Saturday             // 6
		Sunday               // 7
	)

	fmt.Printf("Pi approximately: %v \n\n", pi)

	fmt.Printf("Monday is %v day of the week. \n", Monday)
	fmt.Printf("Tuesday is %v day of the week. \n", Tuesday)
	fmt.Printf("Saturday is %v day of the week. \n", Saturday)
}
